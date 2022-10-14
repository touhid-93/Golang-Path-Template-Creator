package checksummer

import (
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

func hashAllSyncNonRecursive(
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	dirEntries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	filesHashes := make(
		map[string][]byte,
		constants.ArbitraryCapacity64)

	var hashErrors []string

	for _, dirEntry := range dirEntries {
		// If the error is not nil or if the file is not regular, skip it
		if fileMode, err := dirEntry.Info(); err != nil || !fileMode.Mode().IsRegular() {
			continue
		}

		path := filepath.Join(root, dirEntry.Name())

		errResult := FileRaw(path, hashType)

		if errResult.HasError() {
			hashErrors = append(
				hashErrors,
				errResult.ErrorWrapper.ErrorString())

			continue
		}

		filesHashes[path] = errResult.Values
	}

	return filesHashes, errcore.SliceToError(hashErrors)
}
