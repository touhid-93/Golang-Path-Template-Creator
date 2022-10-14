package pathfilter

type recursiveFilterGetterParam struct {
	separator, rootPath     string
	eachFilterPath          string
	rootPathPlusSeparator   string
	extensionsLength        int
	additionalFiltersLength int
	additionalFilters       []string
	extensions              []string
}
