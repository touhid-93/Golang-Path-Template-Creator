package splitinternal

import "gitlab.com/evatix-go/core/constants"

func GetBaseDirNames(currentPath string) (baseDirNames *[]string) {
	list := make(
		[]string,
		constants.Zero,
		constants.ArbitraryCapacity15)

	baseDirPath, baseDirName :=
		GetBaseDirPlusName(currentPath)

	for baseDirName != constants.EmptyString {
		list = append(list, baseDirName)

		baseDirPath, baseDirName =
			GetBaseDirPlusName(baseDirPath)

		if baseDirName == constants.EmptyString {
			list = append(list, baseDirPath)
		}
	}

	return &list
}
