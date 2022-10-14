package pathinsfmt

type SymbolicLink struct {
	Src                string `json:"Src"`
	Dst                string `json:"Dst"`
	IsClearBefore      bool   `json:"IsClearBefore,omitempty"`
	IsMkDirAll         bool   `json:"IsMkDirAll,omitempty"`
	IsNormalizePath    bool   `json:"IsNormalizePath,omitempty"`
	IsSkipOnSrcMissing bool   `json:"IsSkipOnSrcMissing,omitempty"`
	IsSkipOnExist      bool   `json:"IsSkipOnExist,omitempty"`
}
