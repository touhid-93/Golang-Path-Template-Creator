package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/urischemes"
)

func whichPrefix(stringToCheck string) urischemes.Type {
	if strings.HasPrefix(stringToCheck, constants.UriSchemePrefixStandard) {
		return urischemes.UriSchemePrefixStandard
	}

	if strings.HasPrefix(stringToCheck, constants.UriSchemePrefixTwoSlashes) {
		return urischemes.UriSchemePrefixTwoSlashes
	}

	return urischemes.UriUnknown
}
