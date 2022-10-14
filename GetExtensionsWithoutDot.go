package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetExtensionsWithoutDot(
	currentPath string,
) string {
	_, ext := splitinternal.GetBothExtension(
		currentPath)

	return ext
}
