package pathhelpercore

import "gitlab.com/evatix-go/core/coredata/corestr"

type CompiledPaths struct {
	*SourcePaths
	Compiled   *corestr.Collection `json:"Compiled,omitempty"`
	IsResolved bool                `json:"IsResolved"`
	HasIssues  bool                `json:"HasIssues"`
}
