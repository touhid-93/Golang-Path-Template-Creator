package fsinternal

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func ParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return splitinternal.GetBaseDir(location)
}
