package envpath

type (
	// linuxEnvPathCrudFunc crudEnvPaths - path for add, update or remove
	// linuxCurrentEnvPathRaw - linux current environment path string, read from /etc/environment file.
	//  - Format PATH="path1:path2:path3"
	linuxEnvPathCrudFunc func(
		crudEnvPaths []string,
		linuxCurrentEnvPathRaw string,
	) string
)
