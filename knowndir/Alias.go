package knowndir

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type Alias string

//goland:noinspection ALL
const (
	ApacheLinuxPath     Alias = "/etc/apache/"
	AppData             Alias = "AppData"
	AppDataUnix         Alias = "/usr/share"
	Bin                 Alias = "bin"
	BinUnix             Alias = "/usr/bin"
	Documents           Alias = "Documents"
	Downloads           Alias = "Downloads"
	Drivers             Alias = "drivers"
	DriversUnix         Alias = "/lib/modules/$(uname -r)/kernel/drivers/"
	Etc                 Alias = "etc"
	Fonts               Alias = "Fonts"
	FontsUnix           Alias = "usr/share/fonts"
	GitGlobalWin        Alias = ".gitconfig"
	GitGlobalUnix       Alias = "/etc/gitconfig"
	GitGlobalUnixXdg    Alias = "XDG_CONFIG_HOME/git/config"
	HostFile            Alias = "hosts"
	Local               Alias = "Local"
	LocalTempWin        Alias = "local\\temp"
	LocalTempUnix       Alias = "tmp"
	Music               Alias = "Music"
	NginxLinuxPath      Alias = "/etc/nginx/"
	Pictures            Alias = "Pictures"
	ProgramFiles32Or64  Alias = "Program Files"
	ProgramFilesX32In64 Alias = "Program Files (x86)"
	ProgramData         Alias = "Program Data"
	Roaming             Alias = "Roaming"
	Services            Alias = "services"
	SSHGlobal           Alias = ".ssh"
	System32            Alias = "System32"
	System64            Alias = "SysWOW64"
	SystemUnix          Alias = "/etc/systemd/system"
	Temp                Alias = "Temp"
	TempDir             Alias = "TMPDIR"
	UnixRoot            Alias = "/"
	User                Alias = "User"
	UserBin             Alias = "UserBin"
	Users               Alias = "Users"
	Videos              Alias = "Videos"
	WindowsDirectory    Alias = "windir"
	WindowsCDrive       Alias = "C:\\"
	// for paths of nginx and  apache
	Conf             Alias = "conf.d"
	ConfAvailable    Alias = "conf-available"
	ConfEnabled      Alias = "conf-enabled"
	ModsAvailable    Alias = "mods-available"
	ModsEnabled      Alias = "mods-enabled"
	ModulesAvailable Alias = "modules-available"
	ModulesEnabled   Alias = "modules-enabled"
	SitesAvailable   Alias = "sites-available"
	SitesEnabled     Alias = "sites-enabled"
	MimeTypes        Alias = "mime.types"
	EtcEnvironment   Alias = "/etc/environment"
)

func (alias Alias) Value() string {
	return string(alias)
}

func (alias Alias) String() string {
	return string(alias)
}

// CombineWithKnownDirs directory.Value() + constants.PathSeparator + knownDirectories.join(constants.PathSeparator)
// Warning: It doesn't perform complex tasks like long path normalize, long path (windows) fix, double separator to single and so on.
func (alias Alias) CombineWithKnownDirs(knownDirectories ...Alias) string {
	paths := make([]string, 0, len(knownDirectories)+2)
	paths = append(paths, alias.Value())

	for _, knownDirectory := range knownDirectories {
		paths = append(paths, knownDirectory.Value())
	}

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

func (alias Alias) CombineWith(paths ...string) string {
	paths = append(paths, alias.Value())

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

// GetPrefixCombinedWith Alias.Value() + constants.PathSeparator + paths with separator
// Warning: It doesn't perform complex tasks like long path normalize, long path (windows) fix, double separator to single and so on.
func (alias Alias) GetPrefixCombinedWith(paths ...string) string {
	paths = append([]string{alias.Value()}, paths...)

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

func (alias Alias) ValuePtr() *string {
	value := alias.Value()

	return &value
}
