package expandpath

import "gitlab.com/evatix-go/core/coreutils/stringutil"

// GetCompiledPath
//
//  exactly replaces as is.
func GetCompiledPath(
	pathTemplate string,
	compilingMap map[string]string,
) string {
	return stringutil.ReplaceTemplate.DirectKeyUsingMap(
		pathTemplate,
		compilingMap)
}
