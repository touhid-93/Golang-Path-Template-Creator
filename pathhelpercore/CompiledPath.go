package pathhelpercore

type CompiledPath struct {
	SourcePath
	Compiled   string `json:"Compiled"`
	IsResolved bool   `json:"IsResolved"`
}
