package downloadinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(download *pathinsfmt.Download) *errorwrapper.Wrapper {
	if download == nil {
		return nil
	}

	if download.IsSkipOnExist && download.PathStat().IsExist {
		return nil
	}

	createErr := download.
		CreateDirInstruction().
		CreateDefault()

	if createErr.HasError() {
		return createErr
	}

	bashCommandArg := aria2cBashCommandArg(download)

	scriptRunningErr := errcmd.
		New.BashScript.ArgsErr(bashCommandArg)

	if scriptRunningErr.HasError() {
		return scriptRunningErr
	}

	return downloadChecksumVerify(download)
}
