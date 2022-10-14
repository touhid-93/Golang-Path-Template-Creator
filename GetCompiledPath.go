package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func GetCompiledPath(
	pathTemplate string,
	compilingMap *map[string]string,
) string {
	if ispath.Empty(pathTemplate) {
		return pathTemplate
	}

	if compilingMap == nil || len(*compilingMap) == 0 {
		return pathTemplate
	}

	for key, value := range *compilingMap {
		pathTemplate = strings.ReplaceAll(
			pathTemplate,
			key,
			value)
	}

	return pathTemplate
}
