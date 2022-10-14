package cmdprefix

import "fmt"

// recursiveFormat (recursive format, non recursive format)
func recursiveFormat(
	isRecursive bool,
	formatRecursive string,
	nonRecursiveFormat string,
	args ...interface{},
) string {
	if isRecursive {
		return fmt.Sprintf(formatRecursive, args...)
	}

	return fmt.Sprintf(nonRecursiveFormat, args...)
}
