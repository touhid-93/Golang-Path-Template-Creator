package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func RuneRepeatFixToSingle(fixingSlugRune rune, inputString string) string {
	fixingSlug := string(fixingSlugRune)
	fixingSlugRepeat3 := strings.Repeat(
		fixingSlug,
		constants.Capacity3)
	fixingSlugRepeat2 := strings.Repeat(
		fixingSlug,
		constants.Capacity2)

	inputString = strings.ReplaceAll(
		inputString,
		fixingSlugRepeat3,
		fixingSlug)

	return strings.ReplaceAll(
		inputString,
		fixingSlugRepeat2,
		fixingSlug)
}
