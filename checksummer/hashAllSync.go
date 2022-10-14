package checksummer

import (
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func hashAllSync(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	if isRecursive {
		return hashAllSyncRecursive(root, hashType)
	}

	return hashAllSyncNonRecursive(root, hashType)
}
