package pathinsfmt

type CliRunner struct {
	FilesSelector []FilesSelector       `json:"FilesSelector,omitempty"`
	Processors    []ExecutableProcessor `json:"ExecutableProcessor,omitempty"`
}
