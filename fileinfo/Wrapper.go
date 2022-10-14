package fileinfo

import (
	"encoding/json"
	"os"
	"time"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Wrapper struct {
	FileInfo     os.FileInfo
	ErrorWrapper *errorwrapper.Wrapper
	RawPath      string
	IsDirectory  bool
	IsFile       bool
	IsEmptyPath  bool
	pathExists   issetter.Value
	Separator    string
	baseDir      *string
	parent       *Wrapper
}

func (it *Wrapper) HasError() bool {
	return it.ErrorWrapper.HasError()
}

func (it *Wrapper) IsPathExists() bool {
	if it.pathExists.IsUninitialized() {
		isPathExists := !it.HasError() && (it.IsDirectory || it.IsFile)
		it.pathExists = issetter.GetBool(isPathExists)
	}

	return it.pathExists.IsTrue()
}

func (it *Wrapper) BaseDir() string {
	if it.baseDir != nil {
		return *it.baseDir
	}

	baseDir := splitinternal.GetBaseDir(it.RawPath)
	it.baseDir = &baseDir

	return baseDir
}

func (it *Wrapper) Parent() *Wrapper {
	if it.parent != nil {
		return it.parent
	}

	it.parent = New(
		it.BaseDir(),
		it.Separator)

	return it.parent
}

func (it *Wrapper) GetBothExtensions() (dotExt, ext string) {
	return splitinternal.GetBothExtension(it.RawPath)
}

func (it *Wrapper) FileName() (filename string) {
	return it.FileInfo.Name()
}

func (it *Wrapper) FileNameWithoutExt() (filename string) {
	return splitinternal.GetFileNameWithoutExt(it.RawPath)
}

func (it *Wrapper) Size() int64 {
	return it.FileInfo.Size()
}

func (it *Wrapper) ModifyTime() time.Time {
	return it.FileInfo.ModTime()
}

func (it *Wrapper) AllSplits() *[]string {
	return splitinternal.GetAllSplitsWithSep(
		it.RawPath,
		it.Separator)
}

func (it Wrapper) String() string {
	return it.RawPath
}

func (it *Wrapper) ToString(sep string) string {
	return normalize.PathUsingSeparatorUsingSingleIf(
		true,
		sep,
		it.RawPath,
	)
}

func (it *Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *Wrapper) UnmarshalJSON(data []byte) error {
	var dataModel WrapperDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.RawPath = dataModel.RawPath
		it.IsDirectory = dataModel.IsDirectory
		it.IsFile = dataModel.IsFile
		it.IsEmptyPath = dataModel.IsEmptyPath
		it.Separator = dataModel.Separator

		fileInfo, err2 := os.Stat(dataModel.RawPath)
		it.FileInfo = fileInfo
		it.ErrorWrapper = errnew.
			Path.
			Error(errtype.PathInfoFailed, err2, dataModel.RawPath)
	}

	return err
}

func (it *Wrapper) JsonModel() *WrapperDataModel {
	return &WrapperDataModel{
		RawPath:     it.RawPath,
		IsDirectory: it.IsDirectory,
		IsFile:      it.IsFile,
		IsEmptyPath: it.IsEmptyPath,
		Separator:   it.Separator,
	}
}

func (it *Wrapper) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Wrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it Wrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (it *Wrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrapper {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Wrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Wrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Wrapper) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Wrapper) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}
