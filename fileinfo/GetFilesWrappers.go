package fileinfo

func GetFilesWrappers(rootPath string, wrappersIn []*Wrapper) *Wrappers {
	if len(wrappersIn) == 0 {
		return EmptyWrappers()
	}

	length := len(wrappersIn)
	files := make([]*Wrapper, 0, length)

	for _, wrapper := range wrappersIn {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrappers := &Wrappers{
		RootPath:    rootPath,
		Items:       files,
		directories: EmptyWrappers(),
	}

	filesWrappers.files = filesWrappers

	return filesWrappers
}
