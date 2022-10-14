package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

func LinuxTouchFileUsingLock(fullPath string) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return errcmd.
		New.BashScript.ArgsErr(
		cmdprefix.Touch,
		fullPath,
	)
}
