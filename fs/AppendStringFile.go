package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendStringFile(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return AppendFile(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
