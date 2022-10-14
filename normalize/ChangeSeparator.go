package normalize

import "strings"

func ChangeSeparator(
	path,
	currentSeparator,
	changeSeparator string,
) string {
	return strings.ReplaceAll(
		path,
		currentSeparator,
		changeSeparator)
}
