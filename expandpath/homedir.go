package expandpath

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

const (
	homeCaps     = "HOME"
	userProfiles = "USERPROFILE"
	homeDrive    = "HOMEDRIVE"
	homePath     = "HOMEPATH"
)

// TODO refactor later
// DisableCache will disable caching of the homeCaps directory. Caching is enabled
// by default.
var DisableCache bool
var homedirCache string
var cacheLock sync.RWMutex

// dirInternal returns the homeCaps directory for the executing user.
//
// This uses an OS-specific method for discovering the homeCaps directory.
// An error is returned if a homeCaps directory cannot be detected.
func dirInternal() (string, error) {
	if !DisableCache {
		cacheLock.RLock()
		cached := homedirCache
		cacheLock.RUnlock()
		if cached != "" {
			return cached, nil
		}
	}

	cacheLock.Lock()
	defer cacheLock.Unlock()

	var result string
	var err error
	if osconsts.IsWindows {
		result, err = dirWindows()
	} else {
		// Unix-like system, so just assume Unix
		result, err = dirUnix()
	}

	if err != nil {
		return "", err
	}

	homedirCache = result

	return result, nil
}

// Home expands the path to include the homeCaps directory if the path
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
// returned as-is.
// Reference : https://bit.ly/3rZbbKH
func Home(path string) (string, *errorwrapper.Wrapper) {
	if len(path) == 0 {
		return path, errorwrapper.StaticEmptyPtr
	}

	if path[0] != constants.TildeChar {
		return path, errorwrapper.StaticEmptyPtr
	}

	if len(path) > constants.One &&
		path[1] != constants.ForwardSlash[0] &&
		path[1] != constants.BackSlash[0] {
		return constants.EmptyString,
			errnew.Message.New(
				errtype.PathExpand,
				"cannot expand user-specific homeCaps dirInternal")
	}

	dir, err := dirInternal()
	if err != nil {
		return "", errnew.Error.Type(
			errtype.PathExpand,
			err)
	}

	return filepath.Join(dir, path[1:]),
		errorwrapper.StaticEmptyPtr
}

func dirUnix() (string, error) {
	homeEnv := homeCaps
	if runtime.GOOS == "plan9" {
		// On plan9, env pathsconst are lowercase.
		homeEnv = "home"
	}

	// First prefer the HOME environmental variable
	if home := os.Getenv(homeEnv); home != "" {
		return home, nil
	}

	var stdout bytes.Buffer

	// If that fails, try OS specific commands
	if osconsts.IsDarwinOrMacOs {
		cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
		cmd.Stdout = &stdout
		if err := cmd.Run(); err == nil {
			result := strings.TrimSpace(stdout.String())
			if result != "" {
				return result, nil
			}
		}
	} else {
		cmd := exec.Command("getent", "passwd", strconv.Itoa(os.Getuid()))
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			// If the error is ErrNotFound, we ignore it. Otherwise, return it.
			if err != exec.ErrNotFound {
				return "", err
			}
		} else {
			if passwd := strings.TrimSpace(stdout.String()); passwd != "" {
				// username:password:uid:gid:gecos:homeCaps:shell
				passwdParts := strings.SplitN(passwd, ":", 7)
				if len(passwdParts) > 5 {
					return passwdParts[5], nil
				}
			}
		}
	}

	// If all else fails, try the shell
	stdout.Reset()
	cmd := exec.Command("sh", "-c", "cd && pwd")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading homeCaps directory")
	}

	return result, nil
}

func dirWindows() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv(homeCaps); home != "" {
		return home, nil
	}

	// Prefer standard environment variable USERPROFILE
	if home := os.Getenv(userProfiles); home != "" {
		return home, nil
	}

	drive := os.Getenv(homeDrive)
	path := os.Getenv(homePath)
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
}
