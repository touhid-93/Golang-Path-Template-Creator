package pathhelper

import "gitlab.com/evatix-go/core/constants"

func GetSlugsOfDefaultHyphenAll(
	paths ...string,
) []string {
	return GetSlugsOf(
		true,
		constants.HyphenRune,
		constants.HyphenRune,
		paths)
}
