package fileinfo

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

type Wrappers struct {
	RootPath            string
	Separator           string
	Items               []*Wrapper
	directories         *Wrappers
	files               *Wrappers
	ErrorWrapper        *errorwrapper.Wrapper
	pathsCollection     *PathsCollection
	recursiveDirs       *PathsCollection
	fileNamesCollection *FileNamesCollection
}

func (it *Wrappers) HasAny() bool {
	return !it.IsEmpty()
}

func (it *Wrappers) RootFiles() *Wrappers {
	if it.files != nil {
		return it.files
	}

	if it.IsEmpty() {
		it.files = EmptyWrappers()
	}

	files := make(
		[]*Wrapper,
		0,
		it.Length())

	for _, wrapper := range it.Items {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrapper := &Wrappers{
		Items: files,
	}

	filesWrapper.files = filesWrapper
	it.files = filesWrapper

	return it.files
}

func (it *Wrappers) RootDirs() *Wrappers {
	if it.directories != nil {
		return it.directories
	}

	if it.IsEmpty() {
		it.directories = EmptyWrappers()
	}

	dirs := make(
		[]*Wrapper,
		0,
		it.Length())

	for _, wrapper := range it.Items {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		Items: dirs,
	}

	dirWrappers.directories = dirWrappers
	it.directories = dirWrappers

	return it.directories
}

func (it *Wrappers) PathsCollection() *PathsCollection {
	if it.pathsCollection != nil {
		return it.pathsCollection
	}

	it.pathsCollection = NewPathsUsingWrappers(
		it.RootPath,
		it.Separator,
		it)

	return it.pathsCollection
}

func (it *Wrappers) FileNamesCollection() *FileNamesCollection {
	if it.fileNamesCollection != nil {
		return it.fileNamesCollection
	}

	it.fileNamesCollection = NewFileNamesUsingWrappers(it)

	return it.fileNamesCollection
}

func (it *Wrappers) IsEmpty() bool {
	return it == nil ||
		it.ErrorWrapper.HasError() ||
		len(it.Items) == 0
}

func (it *Wrappers) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *Wrappers) IsNameContains(
	name string,
	isCaseSensitive bool,
) bool {
	return it.
		FileNamesCollection().
		IsContains(
			name,
			isCaseSensitive)
}

func (it Wrappers) Json() corejson.Result {
	return corejson.New(it)
}

func (it Wrappers) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Wrappers) JsonModel() *Wrappers {
	return it
}

func (it *Wrappers) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Wrappers) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Wrappers) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Wrappers) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Wrappers) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Wrappers) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrappers, error) {
	err := jsonResult.Unmarshal(
		&it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Wrappers) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrappers {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}
