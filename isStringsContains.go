package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
)

func isStringsContains(array []string, findingItem string) bool {
	if ispathinternal.EmptyArray(array) {
		return false
	}

	for _, arrayItem := range array {
		if strings.Compare(arrayItem, findingItem) == 0 {
			return true
		}
	}

	return false
}
