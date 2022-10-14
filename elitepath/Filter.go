package elitepath

import (
	"regexp"

	"gitlab.com/evatix-go/core/regexnew"
)

type Filter struct {
	NameFilter      *ValueFilter `json:"NameFilter,omitempty"`
	PathFilter      *ValueFilter `json:"PathFilter,omitempty"`
	ExistFilter     *ExistFilter `json:"ExistFilter,omitempty"`
	NameRegexFilter string       `json:"NameRegexFilter,omitempty"`
	PathRegexFilter string       `json:"PathRegexFilter,omitempty"`
}

func (it Filter) nameRegex() (*regexp.Regexp, error) {
	if it.NameRegexFilter == "" {
		return nil, errEmptyRegex
	}

	return regexnew.CreateLock(it.NameRegexFilter)
}

func (it Filter) pathRegex() (*regexp.Regexp, error) {
	if it.PathRegexFilter == "" {
		return nil, errEmptyRegex
	}

	return regexnew.CreateLock(it.PathRegexFilter)
}

func (it *Filter) IsMatch(path *Path) bool {
	if path == nil {
		return false
	}

	if it == nil {
		return true
	}

	if !it.isMatchNamePath(path.FileNameWithExt(), path.CompiledPath()) {
		return false
	}

	if !it.ExistFilter.IsMatch(path) {
		return false
	}

	return true
}

func (it *Filter) IsMatchLazy(lazyPath *LazyPath) bool {
	if lazyPath == nil {
		return false
	}

	if it == nil {
		return true
	}

	fullPath := lazyPath.Path.CompiledPath()
	name := lazyPath.LocationInfo().FileNameWithExtension
	if !it.isMatchNamePath(name, fullPath) {
		return false
	}

	if !it.ExistFilter.IsMatchLazy(lazyPath) {
		return false
	}

	return true
}

func (it *Filter) isMatchNamePath(name, fullPath string) bool {
	if !it.PathFilter.IsMatch(fullPath) {
		return false
	}

	if !it.NameFilter.IsMatch(name) {
		return false
	}

	if it.PathRegexFilter == "" && it.NameRegexFilter == "" {
		return true
	}

	// any regex exist
	nameRegex, err := it.nameRegex()

	if err == nil && nameRegex != nil && !nameRegex.MatchString(name) {
		return false
	}

	pathRegex, err2 := it.pathRegex()

	if err2 == nil && pathRegex != nil && !pathRegex.MatchString(fullPath) {
		return false
	}

	return true
}
