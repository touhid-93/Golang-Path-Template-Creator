package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func RemoveDoubleUriSeparator(path, separator string) string {
	if strings.Contains(path, constants.DoubleDash) {
		for {
			path = strings.ReplaceAll(path, constants.DoubleDash, separator)
			if !strings.Contains(path, constants.DoubleDash) {
				break
			}
		}
	}

	if strings.Contains(path, constants.DoubleUnderscore) {
		for {
			path = strings.ReplaceAll(path, constants.DoubleUnderscore, separator)
			if !strings.Contains(path, constants.DoubleUnderscore) {
				break
			}
		}
	}

	return path
}
