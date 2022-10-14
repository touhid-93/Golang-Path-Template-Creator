package checksummer

import "gitlab.com/evatix-go/pathhelper/hashas"

func hashAll(
	isAsync bool,
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	if isAsync {
		return hashAllAsync(
			isRecursive,
			root,
			hashType)
	}

	return hashAllSync(
		isRecursive,
		root,
		hashType)
}
