package fileinfopath

import (
	"os"
	"time"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/iserror"
	"gitlab.com/evatix-go/core/namevalue"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompareinternal"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

type Instance struct {
	FileInfo         os.FileInfo
	FullPath         string
	Error            error
	compiledToString corestr.SimpleStringOnce
}

func New(location string) *Instance {
	fileInfo, err := os.Stat(location)

	return &Instance{
		FileInfo: fileInfo,
		FullPath: location,
		Error:    err,
	}
}

func NewUsingStat(
	fullPath string,
	pathExistStat *chmodhelper.PathExistStat,
) *Instance {
	return &Instance{
		FileInfo: pathExistStat.FileInfo,
		FullPath: fullPath,
		Error:    pathExistStat.Error,
	}
}

func (it *Instance) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *Instance) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *Instance) IsInvalidFileInfo() bool {
	return it == nil || it.FileInfo == nil
}

func (it *Instance) HasFileInfo() bool {
	return it != nil && it.FileInfo != nil
}

func (it *Instance) IsFile() bool {
	return it.IsExist() && !it.FileInfo.IsDir()
}

func (it *Instance) IsDir() bool {
	return it.IsExist() && it.FileInfo.IsDir()
}

func (it *Instance) IsExist() bool {
	return it != nil && it.FileInfo != nil && (it.Error == nil || !os.IsNotExist(it.Error))
}

// IsInvalidPath
//
// it == nil || it.FileInfo == nil || it.Error != nil
func (it *Instance) IsInvalidPath() bool {
	return it == nil || it.FileInfo == nil || it.Error != nil
}

// FileName
// Returns file name with extension
func (it *Instance) FileName() string {
	if it.HasFileInfo() {
		return it.FileInfo.Name()
	}

	return constants.EmptyString
}

func (it *Instance) FileNameWithoutExt() string {
	if it.HasFileInfo() {
		return splitinternal.GetFileNameWithoutExt(it.FileInfo.Name())
	}

	return constants.EmptyString
}

func (it *Instance) BothExtension() (dotExt, ext string) {
	if it.HasFileInfo() {
		return splitinternal.GetBothExtension(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *Instance) DotExtension() (fileName, dotExt string) {
	if it.HasFileInfo() {
		return splitinternal.GetFileNameDotExt(it.FileInfo.Name())
	}

	return constants.EmptyString, constants.EmptyString
}

func (it *Instance) FileNameExt() string {
	return it.FileName()
}

func (it *Instance) Mode() os.FileMode {
	if it.HasFileInfo() {
		return it.FileInfo.Mode()
	}

	return constants.Zero
}

func (it *Instance) LastModifiedAt() *time.Time {
	if it.HasFileInfo() {
		mod := it.FileInfo.ModTime()

		return &mod
	}

	return nil
}

func (it *Instance) Size() *int64 {
	if it.HasFileInfo() {
		sz := it.FileInfo.Size()

		return &sz
	}

	return nil
}

func (it *Instance) IsEqualDefault(right *Instance) bool {
	return it.IsEqual(
		true,
		false,
		false,
		right)
}

func (it *Instance) IsEqual(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	right *Instance,
) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.FullPath == right.FullPath {
		return true
	}

	if iserror.NotEqual(it.Error, right.Error) {
		return false
	}

	return ispathinternal.FileInfoDetailedEqual(
		isQuickVerifyOnPathEqual,
		isPathMustMatchIfDir,
		isVerifyContent,
		it.FullPath,
		right.FullPath,
		it.FileInfo,
		right.FileInfo)
}

func (it *Instance) CompareFileInfoLastModifiedDate(right os.FileInfo) corecomparator.Compare {
	return pathcompareinternal.FileInfoLastModified(it.FileInfo, right)
}

func (it *Instance) CompareSize(anotherInstance *Instance) corecomparator.Compare {
	return pathcompareinternal.SizePtr(it.Size(), anotherInstance.Size())
}

func (it *Instance) CompareLastModified(anotherInstance *Instance) corecomparator.Compare {
	return pathcompareinternal.LastModifiedPtr(it.LastModifiedAt(), anotherInstance.LastModifiedAt())
}

func (it *Instance) NotFileError() *errorwrapper.Wrapper {
	if it.IsFile() {
		return nil
	}

	return errnew.
		Path.
		Messages(
			errtype.File,
			it.FullPath,
			"Cannot read invalid path or a directory. (required file)")
}

func (it *Instance) NotDirError() *errorwrapper.Wrapper {
	if it.IsDir() {
		return nil
	}

	return errnew.
		Path.
		Messages(
			errtype.Directory,
			it.FullPath,
			"Cannot read invalid path or a file. (required directory)")
}

func (it *Instance) String() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.compiledToString.IsInitialized() {
		return it.compiledToString.String()
	}

	nameValues := namevalue.NewNewNameValuesCollectionUsing(
		false,
		namevalue.Instance{
			Name:  "FullPath",
			Value: it.FullPath,
		},
	)

	nameValues.AddsIf(
		it.FileInfo != nil,
		namevalue.Instance{
			Name:  "FileInfo",
			Value: FileInfoString(it.FileInfo),
		})

	nameValues.AddsIf(
		it.Error != nil,
		namevalue.Instance{
			Name:  "Error",
			Value: it.Error,
		})

	toString := nameValues.JoinJsonStrings(
		consts.FileInfoEachLineJoiner,
	)

	return it.compiledToString.GetPlusSetOnUninitialized(toString)
}

func (it *Instance) Json() corejson.Result {
	return corejson.New(it)
}

func (it *Instance) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Instance) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it Instance) JsonModelAny() interface{} {
	return it
}

func (it *Instance) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *Instance) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Instance) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Instance) ErrorWrapper(errType errtype.Variation) *errorwrapper.Wrapper {
	if it.HasError() {
		return errnew.
			Path.
			Error(errType, it.Error, it.FullPath)
	}

	return nil
}
