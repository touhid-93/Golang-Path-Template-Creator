package pathfuncs

import (
	"os"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func FilterFullPaths(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	filter Filter,
	rootPath string,
	fullPaths ...string,
) *corestr.SimpleSlice {
	length := len(fullPaths)

	if filter == nil || length == 0 {
		return corestr.Empty.SimpleSlice()
	}

	foundItems := stringslice.MakeDefault(length)
	for _, fullPath := range fullPaths {
		fileInfo, err := os.Stat(fullPath)
		hasFileInfo := fileInfo != nil

		arg := &FilterArg{
			RootPath:    rootPath,
			FileName:    splitinternal.GetName(fullPath),
			FullPath:    fullPath,
			IsFile:      hasFileInfo && !fileInfo.IsDir(),
			IsDirectory: hasFileInfo && fileInfo.IsDir(),
			FileInfo:    fileInfo,
			InputError:  err,
		}

		filterResult := filter(arg)
		errCollection.AddWrapperPtr(
			filterResult.ErrorWrapper)

		if filterResult.IsTake {
			foundItems = append(foundItems, fullPath)
		}

		if filterResult.IsBreak {
			break
		}

		hasErr := filterResult.ErrorWrapper.HasError()
		if !isContinueOnError && hasErr {
			break
		}
	}

	return corestr.New.SimpleSlice.Strings(foundItems)
}
