package downloadinsexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func InstallAria() *errorwrapper.Wrapper {
	return errcmd.
		New.BashScript.ArgsErr(
		installAriaBash)
}
