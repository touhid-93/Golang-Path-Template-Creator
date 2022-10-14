package pathhelpercore

type SourceDestination struct {
	Source        *CompiledPath `json:"Source"`
	Destination   *CompiledPath `json:"Destination"`
	IsForce       bool
	IsSkipOnExist bool
	IsOverwrite   bool
}
