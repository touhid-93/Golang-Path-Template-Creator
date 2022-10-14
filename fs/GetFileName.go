package fs

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetFileName(location string) string {
	return splitinternal.GetName(location)
}
