package fs

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
)

func ParentDir(location string) string {
	if location == "" {
		return constants.EmptyString
	}

	return filepath.Dir(location)
}
