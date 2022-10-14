package pathhelper

import "gitlab.com/evatix-go/core/constants"

var (
	uriRemovePrefixes = []string{
		constants.UriSchemePrefixStandard,
		constants.UriSchemePrefixTwoSlashes,
	}
)
