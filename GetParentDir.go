package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
)

func GetParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return GetBaseDir(location)
}
