package expandnormalize

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func FixIf(
	isNormalizeLogPathFix,
	isExpandEnvVars bool,
	location string,
) string {
	if location == "" {
		return location
	}

	expanded := expandpath.ExpandVariablesIf(
		isExpandEnvVars,
		location)

	return normalize.PathUsingSingleIf(
		isNormalizeLogPathFix,
		expanded)
}
