package checksummer

import (
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/hashas"
)

func hashAllSyncRecursive(
	root string,
	hashType hashas.Variant,
) (map[string][]byte, error) {
	filesHashes := make(
		map[string][]byte,
		constants.ArbitraryCapacity64)

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			errResult := FileRaw(path, hashType)

			// TODO: deal with hash error?
			if errResult.HasError() {
				return nil
			}

			filesHashes[path] = errResult.Values

			return nil
		})

	return filesHashes, err
}
