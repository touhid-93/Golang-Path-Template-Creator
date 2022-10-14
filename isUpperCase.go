package pathhelper

import "gitlab.com/evatix-go/core/constants"

func isUpperCase(eachChar rune) bool {
	return eachChar >= constants.UpperCaseA && eachChar <= constants.UpperCaseZ
}
