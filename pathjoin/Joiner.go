package pathjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Joiner struct {
	items []string
}

func NewJoiner(capacity int) *Joiner {
	list := make(
		[]string,
		constants.Zero,
		capacity)

	return &Joiner{items: list}
}

func NewJoiner5() *Joiner {
	list := make(
		[]string,
		constants.Zero,
		constants.ArbitraryCapacity5)

	return &Joiner{items: list}
}

func EmptyJoiner() *Joiner {
	return &Joiner{items: []string{}}
}

func (it *Joiner) Length() int {
	return len(it.items)
}

func (it *Joiner) IsEmpty() bool {
	return len(it.items) == 0
}

func (it *Joiner) HasItems() bool {
	return len(it.items) > 0
}

func (it *Joiner) Add(addingPath string) *Joiner {
	it.items = append(
		it.items,
		addingPath)

	return it
}

func (it *Joiner) Adds(addingPaths ...string) *Joiner {
	if addingPaths == nil {
		return it
	}

	for _, curPath := range addingPaths {
		it.items = append(
			it.items,
			curPath)
	}

	return it
}

// ToString isNormalizePlusLongPathFix if true then adds UNC path fix for Windows
func (it *Joiner) ToString(
	isExpandEnvVars,
	isNormalizePlusLongPathFix bool,
	sep string,
) string {
	finalPath := strings.Join(it.items, sep)

	expand := expandpath.ExpandVariablesIf(
		isExpandEnvVars,
		finalPath)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		sep,
		expand)
}

// OsSeparatorJoin
//
// Usages osconsts.PathSeparator as separator
func (it *Joiner) OsSeparatorJoin(
	isExpandEnvVars,
	isNormalizePlusLongPathFix bool,
) string {
	return it.ToString(
		isExpandEnvVars,
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator)
}

func (it *Joiner) OsSeparatorJoinNormalized(
	isExpandEnvVars bool,
) string {
	return it.ToString(
		isExpandEnvVars,
		true,
		osconsts.PathSeparator,
	)
}

func (it *Joiner) OsSeparatorJoinExpand(
	isNormalizePlusLongPathFix bool,
) string {
	return it.ToString(
		true,
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
	)
}

// String normalize + expand and long path fix true
//
// Usages osconsts.PathSeparator as separator
func (it Joiner) String() string {
	return it.OsSeparatorJoin(
		true,
		true)
}
