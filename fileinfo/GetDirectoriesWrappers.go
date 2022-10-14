package fileinfo

func GetDirectoriesWrappers(rootPath string, wrapperIn []*Wrapper) *Wrappers {
	if len(wrapperIn) == 0 {
		return EmptyWrappers()
	}

	length := len(wrapperIn)
	dirs := make([]*Wrapper, 0, length)

	for _, wrapper := range wrapperIn {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		RootPath: rootPath,
		Items:    dirs,
		files:    EmptyWrappers(),
	}

	dirWrappers.directories = dirWrappers

	return dirWrappers
}
