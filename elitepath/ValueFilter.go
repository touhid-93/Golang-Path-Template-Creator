package elitepath

import "gitlab.com/evatix-go/core/enums/stringcompareas"

type ValueFilter struct {
	Value           string
	IsCaseSensitive bool
	Compare         stringcompareas.Variant
}

func (it *ValueFilter) IsIgnoreCase() bool {
	return !it.IsCaseSensitive
}

func (it *ValueFilter) IsMatch(content string) bool {
	if it == nil {
		return true
	}

	return it.Compare.IsCompareSuccess(
		it.IsIgnoreCase(),
		content,
		it.Value)
}
