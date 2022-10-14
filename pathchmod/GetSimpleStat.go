package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func GetSimpleStat(
	location string,
) *SimpleStat {
	info, isExist, err := chmodhelper.GetPathExistStatExpand(
		location)

	if err != nil {
		pathErr := errnew.
			Path.
			Error(
				errtype.InvalidPath,
				err,
				location)

		return &SimpleStat{
			Location:        location,
			FileInfo:        info,
			HasFileInfo:     info != nil,
			InvalidFileInfo: info == nil,
			IsNotExist:      true,
			IsExist:         false,
			IsDir:           false,
			IsFile:          false,
			ErrorWrapper:    pathErr,
		}
	}

	var name string
	if info != nil {
		name = info.Name()
	}

	return &SimpleStat{
		Location:        location,
		FileInfo:        info,
		HasFileInfo:     info != nil,
		InvalidFileInfo: info == nil,
		Name:            name,
		IsNotExist:      !isExist,
		IsExist:         isExist,
		IsDir:           isExist && info != nil && info.IsDir(),
		IsFile:          isExist && info != nil && !info.IsDir(),
		ErrorWrapper:    nil,
	}
}
