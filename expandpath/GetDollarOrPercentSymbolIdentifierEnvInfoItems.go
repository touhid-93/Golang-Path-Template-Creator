package expandpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/regexnew"
)

func GetDollarOrPercentSymbolIdentifierEnvInfoItems(
	stringToCheck string,
) []EnvKeyInfo {
	envVariableRawKeys := regexnew.
		DollarIdentifierRegex.
		CompileMust().
		FindAllString(
			stringToCheck,
			constants.MinusOne)

	var envInfoItems []EnvKeyInfo

	if len(envVariableRawKeys) > 0 {
		envInfoItems = GetEnvInfoItemsKeyNames(
			envVariableRawKeys)
	}

	envVariableRawKeysNext := regexnew.
		PercentIdentifierRegex.
		CompileMust().
		FindAllString(
			stringToCheck, constants.MinusOne)

	if len(envVariableRawKeysNext) > 0 {
		envInfoItemsCopy := GetEnvInfoItemsKeyNames(
			envVariableRawKeysNext)

		for _, envInfo := range envInfoItemsCopy {
			envInfoItems = append(
				envInfoItems,
				envInfo)
		}
	}

	return envInfoItems
}
