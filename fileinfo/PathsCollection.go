package fileinfo

import (
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type PathsCollection struct {
	rootPath          string
	pathWrappers      *[]*SimplePathWrapper
	allRecursivePaths *errstr.ResultsWithErrorCollection
	allRecursiveFiles *errstr.ResultsWithErrorCollection
	allRecursiveDirs  *errstr.ResultsWithErrorCollection
	directories       *[]string
	files             *[]string
	separator         string
	ErrorWrapper      *errorwrapper.Wrapper
	parentWrappers    *Wrappers
}

func NewPaths(rootPath, separator string, capacity int) *PathsCollection {
	paths := make([]*SimplePathWrapper, 0, capacity)

	return &PathsCollection{
		rootPath:     rootPath,
		pathWrappers: &paths,
		separator:    separator,
	}
}

func NewPathsUsingWrappers(
	rootPath, separator string,
	wrappers *Wrappers,
) *PathsCollection {
	if wrappers == nil {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   nil,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	if wrappers.IsEmpty() {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   wrappers.ErrorWrapper,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	paths := make(
		[]*SimplePathWrapper,
		wrappers.Length())

	for i, wrapper := range wrappers.Items {
		paths[i] = &SimplePathWrapper{
			Path:        wrapper.RawPath,
			IsDirectory: wrapper.IsDirectory,
		}
	}

	return &PathsCollection{
		rootPath:       rootPath,
		pathWrappers:   &paths,
		ErrorWrapper:   wrappers.ErrorWrapper,
		parentWrappers: wrappers,
		separator:      separator,
	}
}

func NewPathsUsing(
	directoryPath, separator string,
	isNormalize bool,
) *PathsCollection {
	wrappers := NewWrappersPtr(
		directoryPath,
		separator,
		isNormalize)

	return NewPathsUsingWrappers(
		directoryPath,
		separator,
		wrappers)
}

func NewPathsUsingPaths(
	rootPath, separator string,
	recursivePaths *[]string,
) *PathsCollection {
	if recursivePaths == nil {
		return NewPaths(
			rootPath,
			separator,
			0)
	}

	wrappers :=
		NewPaths(
			rootPath,
			separator,
			len(*recursivePaths))

	wrappers.allRecursivePaths =
		&errstr.ResultsWithErrorCollection{
			Values:        *recursivePaths,
			ErrorWrappers: errwrappers.Empty(),
		}

	return wrappers
}

func (it *PathsCollection) Directories() *[]string {
	if it.directories != nil {
		return it.directories
	}

	if it.IsEmpty() {
		it.directories =
			core.EmptyStringsPtr()

		return it.directories
	}

	directories := make([]string, 0, it.Length())

	for _, pathWrapper := range *it.pathWrappers {
		if !pathWrapper.IsDirectory {
			continue
		}

		directories = append(directories, pathWrapper.Path)
	}

	it.directories = &directories

	return it.directories
}

func (it *PathsCollection) Files() *[]string {
	if it.files != nil {
		return it.files
	}

	if it.IsEmpty() {
		it.files = core.EmptyStringsPtr()

		return it.files
	}

	files := make([]string, 0, it.Length())

	for _, pathWrapper := range *it.pathWrappers {
		if pathWrapper.IsDirectory {
			continue
		}

		files = append(files, pathWrapper.Path)
	}

	it.files = &files

	return it.files
}

func (it *PathsCollection) IsEmpty() bool {
	return it.pathWrappers == nil ||
		it.ErrorWrapper.HasError() ||
		len(*it.pathWrappers) == 0
}

func (it *PathsCollection) Length() int {
	if it.pathWrappers == nil || *it.pathWrappers == nil {
		return 0
	}

	return len(*it.pathWrappers)
}

func (it *PathsCollection) IsParentWrappersEmpty() bool {
	return it.parentWrappers == nil ||
		it.parentWrappers.IsEmpty()
}

func (it *PathsCollection) HasParentWrappers() bool {
	return it.parentWrappers != nil
}

func (it *PathsCollection) ParentWrappers() *Wrappers {
	return it.parentWrappers
}

func (it *PathsCollection) Add(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*it.pathWrappers = append(
		*it.pathWrappers,
		wrapper)

	return it
}

func (it *PathsCollection) AddPtr(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*it.pathWrappers = append(
		*it.pathWrappers,
		wrapper)

	return it
}

func (it *PathsCollection) AddWrapper(
	pathWrapper *SimplePathWrapper,
) *PathsCollection {
	*it.pathWrappers = append(
		*it.pathWrappers,
		pathWrapper)

	return it
}

func (it *PathsCollection) Strings() *[]string {
	list := make(
		[]string,
		it.Length())

	for i, wrapper := range *it.pathWrappers {
		list[i] = wrapper.String()
	}

	return &list
}

func (it PathsCollection) String() string {
	list := make(
		[]string,
		constants.ArbitraryCapacity4)
	compiledPaths := strings.Join(
		*it.Strings(),
		constants.NewLineUnix)

	list[coreindexes.I0] = "Root Location :" + it.rootPath
	list[coreindexes.I1] = "Separator :" + it.separator
	if it.ErrorWrapper.HasError() {
		list[coreindexes.I2] = "Error :" + it.ErrorWrapper.
			String()
	}

	list[coreindexes.I3] = compiledPaths

	return strings.Join(
		list,
		constants.NewLineUnix)
}

func (it *PathsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *PathsCollection) UnmarshalJSON(data []byte) error {
	var dataModel PathsCollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.rootPath = dataModel.RootPath
		it.pathWrappers = dataModel.PathWrappers
		it.separator = dataModel.Separator
		it.ErrorWrapper = dataModel.ErrorWrapper
		it.parentWrappers = dataModel.ParentWrappers
	}

	return err
}

func (it *PathsCollection) JsonModel() *PathsCollectionDataModel {
	return &PathsCollectionDataModel{
		RootPath:       it.rootPath,
		PathWrappers:   it.pathWrappers,
		Separator:      it.separator,
		ErrorWrapper:   it.ErrorWrapper,
		ParentWrappers: it.parentWrappers,
	}
}

func (it *PathsCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it PathsCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it PathsCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *PathsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PathsCollection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return nil, defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(jsonResult.Bytes, &it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (it *PathsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PathsCollection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *PathsCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *PathsCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *PathsCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *PathsCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}
