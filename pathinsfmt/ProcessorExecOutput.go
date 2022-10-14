package pathinsfmt

import "gitlab.com/evatix-go/errorwrapper/errcmd"

type ProcessorExecOutput struct {
	CmdOnce             *errcmd.CmdOnce
	ConsoleResult       *errcmd.Result
	ExecutableProcessor *ExecutableProcessor
}
