package pathinsfmt

import "gitlab.com/evatix-go/errorwrapper"

type ProcessorExecOutputs struct {
	CompiledError *errorwrapper.Wrapper
	Outputs       []ProcessorExecOutput
}
