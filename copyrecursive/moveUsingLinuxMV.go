package copyrecursive

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func moveUsingLinuxMV(opts Options, src, dst string) *errorwrapper.Wrapper {
	// Options mapping for mv command
	// 	IsSkipOnExist:  -n
	// 	IsClearDestination:     --remove-destination,
	args := make([]string, 0, constants.ArbitraryCapacity8)
	if opts.IsSkipOnExist {
		args = append(args, mvArgSkipOnExist)
	}

	args = append(args, src)
	args = append(args, dst)

	return errcmd.New.Create(
		false,
		false,
		mvCommand, args...,
	).CompiledErrorWrapper()
}
