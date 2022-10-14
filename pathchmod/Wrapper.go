package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/corecsv"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type Wrapper struct {
	DirChmod, FileChmod os.FileMode
	IsRecursive         bool
	IsSkipOnInvalid     bool
	IsContinueOnError   bool
	IsKeepExistingChmod bool
}

func DefaultWrapper() Wrapper {
	return Wrapper{
		DirChmod:  filemode.DirDefault,
		FileChmod: filemode.FileDefault,
	}
}

func NewWrapperFile(changeFileMode os.FileMode) Wrapper {
	return Wrapper{
		DirChmod:  filemode.DirDefault,
		FileChmod: changeFileMode,
	}
}

func NewWrapperDir(changeDirMode os.FileMode) Wrapper {
	return Wrapper{
		DirChmod:  changeDirMode,
		FileChmod: filemode.FileDefault,
	}
}

func DefaultWrapperRecursive() Wrapper {
	return Wrapper{
		DirChmod:    filemode.DirDefault,
		FileChmod:   filemode.FileDefault,
		IsRecursive: true,
	}
}

func NewWrapper(
	isKeepExistingChmod bool,
	condition chmodins.Condition,
	dirChmod,
	fileChmod os.FileMode,
) Wrapper {
	return Wrapper{
		DirChmod:            dirChmod,
		FileChmod:           fileChmod,
		IsRecursive:         condition.IsRecursive,
		IsSkipOnInvalid:     condition.IsSkipOnInvalid,
		IsContinueOnError:   condition.IsContinueOnError,
		IsKeepExistingChmod: isKeepExistingChmod,
	}
}

func NewWrapperDirFile(
	dirChmod,
	fileChmod os.FileMode,
) Wrapper {
	return Wrapper{
		DirChmod:            dirChmod,
		FileChmod:           fileChmod,
		IsKeepExistingChmod: true,
	}
}

func (it *Wrapper) Condition() *chmodins.Condition {
	if it == nil {
		return &chmodins.Condition{}
	}

	return &chmodins.Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}
}

func (it *Wrapper) OnInvalidFileMode(
	fileChmod os.FileMode,
) *Wrapper {
	if it == nil {
		return NewWrapperFile(fileChmod).ToPtr()
	}

	if it.FileChmod == 0 {
		it.FileChmod = fileChmod
	}

	return it
}

func (it *Wrapper) OnInvalidDirMode(
	dirChmod os.FileMode,
) *Wrapper {
	if it == nil {
		return NewWrapperDir(dirChmod).ToPtr()
	}

	if it.DirChmod == 0 {
		it.DirChmod = dirChmod
	}

	return it
}

func (it *Wrapper) DirChmodDisplay() string {
	if it == nil {
		return ""
	}

	return chmodhelper.FileModeFriendlyString(it.DirChmod)
}

func (it *Wrapper) FileChmodDisplay() string {
	if it == nil {
		return ""
	}

	return chmodhelper.FileModeFriendlyString(it.FileChmod)
}

func (it *Wrapper) References() []ref.Value {
	if it == nil {
		return []ref.Value{}
	}

	return []ref.Value{
		{
			Variable: "ParentDirChmod",
			Value:    it.DirChmodDisplay(),
		},
		{
			Variable: "FileChmod",
			Value:    it.FileChmodDisplay(),
		},
	}
}

func (it *Wrapper) SimpleFileRw(
	filePath string,
) *chmodhelper.SimpleFileReaderWriter {
	if it == nil {
		return chmodhelper.
			New.
			SimpleFileReaderWriter.
			Default(filePath)
	}

	return chmodhelper.
		New.
		SimpleFileReaderWriter.
		Path(it.DirChmod, it.FileChmod, filePath)
}

func (it *Wrapper) SimpleFileRwUsingParent(
	absParentDir,
	absFilePath string,
) *chmodhelper.SimpleFileReaderWriter {
	if it == nil {
		return chmodhelper.
			New.
			SimpleFileReaderWriter.
			Create(
				filemode.DirDefault,
				filemode.FileDefault,
				absParentDir,
				absFilePath)
	}

	return chmodhelper.
		New.
		SimpleFileReaderWriter.
		Create(
			it.DirChmod,
			it.FileChmod,
			absParentDir,
			absFilePath)
}

func (it *Wrapper) MarshalJSON() (jsonBytes []byte, parsedErr error) {
	if it == nil {
		return nil, errnew.
			Null.
			Simple(it).
			CompiledErrorWithStackTraces()
	}

	model := wrapperModel{
		DirChmod:            rwxCreator(it.DirChmod),
		FileChmod:           rwxCreator(it.FileChmod),
		IsRecursive:         it.IsRecursive,
		IsSkipOnInvalid:     it.IsSkipOnInvalid,
		IsContinueOnError:   it.IsContinueOnError,
		IsKeepExistingChmod: it.IsKeepExistingChmod,
	}

	jsonResult := model.Json()

	return jsonResult.Raw()
}

func (it *Wrapper) UnmarshalJSON(rawJsonBytes []byte) error {
	if it == nil {
		return errnew.
			Null.
			Simple(it).
			CompiledErrorWithStackTraces()
	}

	var model wrapperModel
	err := corejson.
		Deserialize.
		UsingBytes(rawJsonBytes, &model)

	if err == nil {
		it.DirChmod = model.DirChmod.ToFileMode()
		it.FileChmod = model.FileChmod.ToFileMode()
		it.IsRecursive = model.IsRecursive
		it.IsSkipOnInvalid = model.IsSkipOnInvalid
		it.IsContinueOnError = model.IsContinueOnError
		it.IsKeepExistingChmod = model.IsKeepExistingChmod
	}

	return err
}

func (it *Wrapper) FriendlyChmodOfDirFiles(
	location string,
) string {
	return FriendlyChmod.OfDirFiles(location)
}

func (it *Wrapper) LogFriendlyChmodOfDirFiles(
	location string,
) {
	FriendlyChmod.LogOfDirFiles(location)
}

func (it *Wrapper) FriendlyChmodOfPath(
	location string,
) string {
	return FriendlyChmod.OfPath(location)
}

func (it *Wrapper) LogFriendlyChmodOfPath(
	location string,
) {
	FriendlyChmod.LogOfPath(location)
}

func (it *Wrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it *Wrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Wrapper) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it Wrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it *Wrapper) ApplyDirs(
	dirPaths ...string,
) (
	rwxInstruction *chmodins.RwxInstruction,
	errWrap *errorwrapper.Wrapper,
) {
	if it == nil {
		return nil, errnew.
			Null.
			WithMessage(
				"dir-paths chmod apply failed : "+
					corecsv.StringsToStringDefault(dirPaths...),
				it)
	}

	// TODO wrap error with details
	return ApplyChmodOnFiles(it.IsRecursive,
		it.IsSkipOnInvalid,
		it.IsContinueOnError,
		it.DirChmod,
		dirPaths...)
}

func (it *Wrapper) ApplyFiles(
	filePaths ...string,
) (
	rwxInstruction *chmodins.RwxInstruction,
	errWrap *errorwrapper.Wrapper,
) {
	if it == nil {
		return nil, errnew.
			Null.
			WithMessage(
				"file-paths chmod apply failed : "+
					corecsv.StringsToStringDefault(filePaths...),
				it)
	}

	// TODO wrap error with details
	return ApplyChmodOnFiles(
		it.IsRecursive,
		it.IsSkipOnInvalid,
		it.IsContinueOnError,
		it.DirChmod,
		filePaths...)
}

func (it *Wrapper) ApplyFileOptions(
	isApply,
	isApplyOnMismatch bool,
	changeFileMode os.FileMode,
	location string,
) (errWrap *errorwrapper.Wrapper) {
	if !isApply {
		return nil
	}

	if it == nil {
		return errnew.Null.WithMessage(
			"location chmod apply failed : "+
				location,
			it)
	}

	if isApplyOnMismatch {
		return ApplyOnMismatch(
			it.IsSkipOnInvalid,
			changeFileMode,
			location)
	}

	_, errWrap = ApplyChmod(
		it.IsRecursive,
		it.IsSkipOnInvalid,
		changeFileMode,
		location)

	return errWrap
}

func (it *Wrapper) ToNonPtr() Wrapper {
	if it == nil {
		return Wrapper{}
	}

	return *it
}

func (it Wrapper) ToPtr() *Wrapper {
	return &it
}
