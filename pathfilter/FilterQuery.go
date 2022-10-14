package pathfilter

import (
	"strings"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
)

type Query struct {
	AdditionalFilters *[]string
	extensions        *[]string
	filterLength      *int
	extensionsLength  *int
	extensionsFixed   bool
}

func NewArray(array ...string) *[]string {
	return &array
}

func NewQuery(additionalFilters, extensions *[]string) *Query {
	return &Query{
		AdditionalFilters: additionalFilters,
		extensions:        extensions,
	}
}

func (q *Query) FilterLength() int {
	if q.filterLength != nil {
		return *q.filterLength
	}

	if q.AdditionalFilters == nil {
		zero := 0
		q.filterLength = &zero

		return zero
	}

	length := len(*q.AdditionalFilters)
	q.filterLength = &length

	return length
}

func (q *Query) CreateJoinedPathsWithFilters(separator, rootPath string) *[]string {
	length := q.FilterLength()
	if length == 0 {
		return &[]string{rootPath}
	}

	list := make([]string, length)

	for i, filter := range *q.AdditionalFilters {
		newPath := rootPath + separator + filter
		list[i] = newPath
	}

	return &list
}

func (q *Query) Extensions() *[]string {
	if q.extensions != nil && q.extensionsFixed {
		return q.extensions
	}

	if q.extensions == nil && !q.extensionsFixed {
		q.extensions = core.EmptyStringsPtr()
		q.extensionsFixed = true

		return q.extensions
	}

	for i, ext := range *q.extensions {
		if strings.Contains(ext, constants.Dot) {
			continue
		}

		(*q.extensions)[i] = ExtensionFilterStart + ext
	}

	q.extensionsFixed = true

	return q.extensions
}

func (q *Query) ExtensionsLength() int {
	if q.extensionsLength != nil {
		return *q.extensionsLength
	}

	if q.extensions == nil {
		q.extensionsLength = constants.ZeroPtr

		return *q.extensionsLength
	}

	length := len(*q.extensions)
	q.extensionsLength = &length

	return length
}
