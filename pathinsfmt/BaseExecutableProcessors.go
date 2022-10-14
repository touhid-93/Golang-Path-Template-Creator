package pathinsfmt

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type BaseExecutableProcessors struct {
	lazyCmdOnceCollection *errcmd.CmdOnceCollection
	ExecutableProcessors  []ExecutableProcessor `json:"ExecutableProcessors,omitempty"`
}

func (b *BaseExecutableProcessors) ProcessorsLength() int {
	if b == nil {
		return constants.Zero
	}

	return len(b.ExecutableProcessors)
}

func (b *BaseExecutableProcessors) HasProcessors() bool {
	return b.ProcessorsLength() > 0
}

func (b *BaseExecutableProcessors) IsEmptyProcessor() bool {
	return b.ProcessorsLength() == 0
}

func (b *BaseExecutableProcessors) LazyCmdOnceCollection() *errcmd.CmdOnceCollection {
	if b.lazyCmdOnceCollection != nil {
		return b.lazyCmdOnceCollection
	}

	b.lazyCmdOnceCollection = b.CreateCmdOnceCollection()

	return b.lazyCmdOnceCollection
}

func (b *BaseExecutableProcessors) CreateCmdOnceCollection() *errcmd.CmdOnceCollection {
	cmdOnceCollection := errcmd.NewCmdOnceCollection(
		b.ProcessorsLength() + constants.One)

	if b.IsEmptyProcessor() {
		return cmdOnceCollection
	}

	for _, executableProcessor := range b.ExecutableProcessors {
		if !executableProcessor.IsEnabled {
			continue
		}

		cmdOnceCollection.Add(
			executableProcessor.CreateCmdOnce())
	}

	return cmdOnceCollection
}

func (b *BaseExecutableProcessors) ExecuteAllOutputs() *ProcessorExecOutputs {
	outputs := make(
		[]ProcessorExecOutput,
		constants.Zero,
		b.ProcessorsLength())

	if b.IsEmptyProcessor() {
		return &ProcessorExecOutputs{
			Outputs:       outputs,
			CompiledError: nil,
		}
	}

	errCollection := errwrappers.Empty()

	for _, executableProcessor := range b.ExecutableProcessors {
		if !executableProcessor.IsEnabled {
			continue
		}

		outputs = append(
			outputs,
			*executableProcessor.
				GetExecuteOutputByExecuting(errCollection))
	}

	return &ProcessorExecOutputs{
		Outputs:       outputs,
		CompiledError: errCollection.GetAsErrorWrapperPtr(),
	}
}
