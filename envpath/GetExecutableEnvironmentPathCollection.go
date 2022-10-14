package envpath

import (
	normalize "gitlab.com/evatix-go/pathhelper/expandpath"
)

func GetExecutableEnvironmentPathCollection() *ExecutableEnvironmentPathCollection {
	rawPaths := ReadEnvPaths()
	pathsCollection := NewExecutableEnvironmentPathCollectionPtr(len(rawPaths))

	for _, rawPath := range rawPaths {
		expandedPath := normalize.ExpandVariables(rawPath)
		executableEnvironmentPath := ExecutableEnvironmentPath{
			Variable: rawPath,
			Expanded: expandedPath,
		}

		pathsCollection.AddPtr(&executableEnvironmentPath)
	}

	return pathsCollection
}
