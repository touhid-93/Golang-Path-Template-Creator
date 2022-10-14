package envpath

type ExecutableEnvironmentPathCollection struct {
	pathsMap *map[string]*ExecutableEnvironmentPath
	paths    *[]*ExecutableEnvironmentPath
}

func NewExecutableEnvironmentPathCollection(capacity int) ExecutableEnvironmentPathCollection {
	pathsMap := make(map[string]*ExecutableEnvironmentPath, capacity)
	paths := make([]*ExecutableEnvironmentPath, 0, capacity)

	return ExecutableEnvironmentPathCollection{
		pathsMap: &pathsMap,
		paths:    &paths,
	}
}

func NewExecutableEnvironmentPathCollectionPtr(capacity int) *ExecutableEnvironmentPathCollection {
	pathsMap := make(map[string]*ExecutableEnvironmentPath, capacity)
	paths := make([]*ExecutableEnvironmentPath, 0, capacity)

	return &ExecutableEnvironmentPathCollection{
		pathsMap: &pathsMap,
		paths:    &paths,
	}
}

func (receiver *ExecutableEnvironmentPathCollection) AddPtr(
	exeEnvPath *ExecutableEnvironmentPath,
) {
	if exeEnvPath != nil {
		(*receiver.pathsMap)[exeEnvPath.Variable] = exeEnvPath
	}
}

func (receiver *ExecutableEnvironmentPathCollection) IsExists(
	exeEnvPath *ExecutableEnvironmentPath,
) bool {
	_, has := (*receiver.pathsMap)[exeEnvPath.Variable]

	return has
}

func (receiver *ExecutableEnvironmentPathCollection) List() *[]*ExecutableEnvironmentPath {
	return receiver.paths
}

func (receiver *ExecutableEnvironmentPathCollection) OnlyNamesCollection() *[]*ExecutableEnvironmentPath {
	return receiver.paths
}
