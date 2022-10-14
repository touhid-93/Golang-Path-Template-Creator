package pathstatlinux

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func Get(location string) *Info {
	if !fsinternal.IsPathExists(location) {
		return InvalidInfo(location)
	}

	if osconsts.IsWindows {
		return InvalidInfoUsingErr(
			location,
			errnew.
				Path.
				Messages(
					errtype.NotSupportInWindows,
					location,
					"pathstatlinux package is not supported in windows."))
	}

	pathStat := errcmd.New.BashScript.ArgsDefault("stat", location)
	errorWrapper := pathStat.CompiledErrorWrapper()

	if errorWrapper.HasError() {
		return InvalidInfoUsingErr(
			location,
			errorWrapper)
	}

	lines := pathStat.CompiledTrimmedOutputLines()

	return ProcessLinesToInfo(
		lines,
		location,
		errorWrapper)
}
