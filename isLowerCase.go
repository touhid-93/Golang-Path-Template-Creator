package pathhelper

import "gitlab.com/evatix-go/core/constants"

func isLowerCase(eachChar rune) bool {
	return eachChar >= constants.LowerCaseA && eachChar <= constants.LowerCaseZ
}
