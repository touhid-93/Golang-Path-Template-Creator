package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func JsonReadUnmarshalOnExist(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	if !IsPathExistsUsingLock(filePath) {
		return nil
	}

	return JsonReadUnmarshal(filePath, unmarshallObjectRef)
}
