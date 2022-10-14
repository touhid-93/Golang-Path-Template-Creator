package fsinternal

import "gitlab.com/evatix-go/errorwrapper"

func WriteStringToFile(
	filePath string, content string,
) *errorwrapper.Wrapper {
	return WriteFileStringDefault(
		filePath,
		content,
	)
}
