package pathinsfmt

import (
	"log"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/enum/scripttype"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type ExecutableProcessor struct {
	Name                                        string   `json:"Name"`
	IsEnabled                                   bool     `json:"IsEnabled,omitempty"`
	IsNormalizePath                             bool     `json:"IsNormalizePath,omitempty"`
	FailedMessage                               string   `json:"FailedMessage,omitempty"`
	SuccessMessage                              string   `json:"SuccessMessage,omitempty"`
	OutputToFile                                string   `json:"OutputToFile,omitempty"`
	BinaryPath                                  string   `json:"BinaryPath,omitempty"`
	Args                                        []string `json:"Args,omitempty"`
	IsSecure, IsDisplayToConsole, IsWriteToFile bool
	ScriptType                                  scripttype.Variant `json:"ScriptType"`
	lazyCmdOnce                                 *errcmd.CmdOnce
}

func (e *ExecutableProcessor) GetExecuteOutputByExecuting(
	errWrapperCollection *errwrappers.Collection,
) *ProcessorExecOutput {
	cmdOnce := e.CreateCmdOnce()
	output := ProcessorExecOutput{
		CmdOnce:             cmdOnce,
		ConsoleResult:       cmdOnce.CompiledResult(),
		ExecutableProcessor: e,
	}

	if e.IsDisplayToConsole {
		log.Print(output.ConsoleResult.OutputString())
	}

	if e.IsWriteToFile {
		normalizePath := normalize.PathUsingSeparatorUsingSingleIf(
			e.IsNormalizePath,
			osconsts.PathSeparator,
			e.OutputToFile)

		writeErr := fsinternal.WriteFileDefault(
			normalizePath,
			output.ConsoleResult.OutputBytes())

		errWrapperCollection.AddWrapperPtr(writeErr)
	}

	errWrapperCollection.AddWrapperPtr(
		cmdOnce.CompiledErrorWrapper())

	return &output
}

func (e *ExecutableProcessor) LazyCmdOnce() *errcmd.CmdOnce {
	if e.lazyCmdOnce != nil {
		return e.lazyCmdOnce
	}

	e.lazyCmdOnce = e.CreateCmdOnce()

	return e.lazyCmdOnce
}

func (e *ExecutableProcessor) CreateCmdOnce() *errcmd.CmdOnce {
	argsCompiled := errcmd.ArgsJoinSlice(e.Args)
	normalizedBinaryPath := normalize.PathUsingSeparatorUsingSingleIf(
		e.IsNormalizePath,
		osconsts.PathSeparator,
		e.BinaryPath)

	script := errcmd.ArgsJoin(
		normalizedBinaryPath,
		argsCompiled)
	hasOutput := e.IsDisplayToConsole || e.IsWriteToFile

	cmdOnce := errcmd.New.Script.Args(
		hasOutput,
		e.IsSecure,
		e.ScriptType,
		script)

	return cmdOnce
}

func (e *ExecutableProcessor) Dispose() {
	if e == nil {
		return
	}

	e.Args = nil
	e.lazyCmdOnce.Dispose()
}
