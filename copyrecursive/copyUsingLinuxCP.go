package copyrecursive

import (
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func copyUsingLinuxCP(opts Options, src, dst string) *errorwrapper.Wrapper {
	if opts.IsMove {
		return moveUsingLinuxMV(opts, src, dst)
	}

	// Options mapping for cp command
	// 	IsSkipOnExist:  -n
	// 	IsRecursive:    -r,
	// 	IsMove:       use   mv?,
	// 	IsClearDestination:     --remove-destination,
	// 	IsUseShellOrCmd:    false,
	// 	IsNormalize:        false,
	// 	IsExpandVar:        false,
	args := make([]string, 0, constants.ArbitraryCapacity8)
	if opts.IsRecursive {
		args = append(args, cpArgRecursive)
	}

	if opts.IsSkipOnExist {
		args = append(args, cpArgSkipOnExist)
	}

	if opts.IsClearDestination {
		args = append(args, cpArgRemoveDst)
	}

	// TODO: linux `cp` hack
	//  this allows us to copy contents of the directory
	//  otherwise it'll copy with the root folder name as well
	src = filepath.Clean(src)
	src = src + string(os.PathSeparator) + constants.Dot

	args = append(args, src)
	args = append(args, dst)

	return errcmd.New.NoOutput.CreateErrorVerbose(
		cpCommand,
		args...,
	).CompiledErrorWrapper()
}
