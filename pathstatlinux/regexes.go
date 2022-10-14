package pathstatlinux

import "regexp"

var (
	bracketsMatcherWithContents = regexp.MustCompile(`\(.+/.+\)`)
)
