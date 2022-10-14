package pathinsfmt

type PathVar struct {
	VarName string `json:"VarName"`
	Value   string `json:"Value"`
	IsRegex bool   `json:"IsRegex"`
}
