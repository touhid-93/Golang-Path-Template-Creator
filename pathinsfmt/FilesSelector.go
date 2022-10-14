package pathinsfmt

type FilesSelector struct {
	Path        string      `json:"Location"`
	Filters     []string    `json:"Filters,omitempty"`
	SkipFilters []string    `json:"SkipFilters,omitempty"`
	Extensions  []string    `json:"Extensions,omitempty"`
	Processors  []string    `json:"ExecutableProcessor,omitempty"`
	Attributes  *Attributes `json:"Attributes,omitempty"`
}
