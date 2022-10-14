package expandpath

import (
	"os"
)

// ExpandEnvironmentVariable
//
//  function takes an array of environment variables (string) as input
//  and outputs a map of expanded path of those variables if the paths exist.
func ExpandEnvironmentVariable(envInfoItems []EnvKeyInfo) map[string]string {
	var expandedPathMap = make(
		map[string]string,
		len(envInfoItems))

	for _, envInfo := range envInfoItems {
		name := envInfo.SimplifiedName
		_, isExist := os.LookupEnv(name)

		if isExist {
			expandedPathMap[envInfo.GivenAs] = os.Getenv(name)
		}
	}

	return expandedPathMap
}
