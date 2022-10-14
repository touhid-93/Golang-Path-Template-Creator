package unipath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/dirinfo"
	"gitlab.com/evatix-go/pathhelper/fileinfo"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathgetter"
	"gitlab.com/evatix-go/pathhelper/pathwrapper"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

const (
	defaultCapacity = constants.N3
)

type Wrapper struct {
	isFinalized    bool
	finalPath      string
	separator      string
	finalizedError *errorwrapper.Wrapper
	collection     *corestr.Collection
}

func New(sep string) *Wrapper {
	return &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.New.Collection.Cap(defaultCapacity),
	}
}

func NewUsingPath(curPath, sep string) *Wrapper {
	wrapper := &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.New.Collection.Cap(constants.ArbitraryCapacity1),
	}

	return wrapper.Add(curPath)
}

func NewCap(cap int, sep string) *Wrapper {
	return &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.New.Collection.Cap(cap),
	}
}

func NewCapStartingPath(cap int, sep, startingPath string) *Wrapper {
	wrapper := &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.New.Collection.Cap(cap),
	}

	return wrapper.Add(startingPath)
}

func (it *Wrapper) IsFinalized() bool {
	return it.isFinalized
}

func (it *Wrapper) Lock() {
	it.collection.Lock()
}

func (it *Wrapper) Unlock() *Wrapper {
	it.collection.Unlock()

	return it
}

func (it *Wrapper) Separator() string {
	return it.separator
}

// AddLock anyPath can contain separators or without separator both are fine
func (it *Wrapper) AddLock(
	anyPath string,
) *Wrapper {
	it.Lock()
	defer it.Unlock()

	return it.Add(anyPath)
}

// anyPath can contain separators or without separator both are fine
// One cannot add path after finalize.
func (it *Wrapper) Add(
	anyPath string,
) *Wrapper {
	it.handleFinalizeError()
	it.collection.Add(anyPath)

	return it
}

func (it *Wrapper) handleFinalizeError() {
	if it.IsFinalized() {
		it.finalizedError.HandleErrorWithMsg(
			"Finalized unipath cannot add or modify data.")
	}
}

func (it *Wrapper) AddStringsPtr(
	stringItems *[]string,
) *Wrapper {
	it.handleFinalizeError()
	it.collection.AddStringsPtr(stringItems)

	return it
}

func (it *Wrapper) AddPointerStringsPtr(
	stringItems *[]*string,
) *Wrapper {
	it.handleFinalizeError()
	it.collection.AddPointerStringsPtr(stringItems)

	return it
}

func (it *Wrapper) Length() int {
	return it.collection.Length()
}

func (it *Wrapper) HasItems() bool {
	return it.collection.HasItems()
}

// Has same one needs to be inserted
func (it *Wrapper) Has(
	pathSplit string,
) bool {
	return it.collection.Has(pathSplit)
}

func (it *Wrapper) IsWindowsSeparator() bool {
	return it.separator == constants.WindowsPathSeparator
}

func (it *Wrapper) IsUnixSeparator() bool {
	return it.separator == constants.ForwardSlash
}

func (it *Wrapper) IsEmpty() bool {
	return it.collection.IsEmpty()
}

func (it *Wrapper) IsValid() bool {
	if it.collection.IsEmpty() {
		return false
	}

	_, errWrap := it.GetFileInfo()

	return errWrap.IsEmpty()
}

func (it *Wrapper) IsEqual(wrapper *Wrapper) bool {
	if wrapper == nil && it == nil {
		return true
	}

	if wrapper == nil || it == nil {
		return false
	}

	if wrapper.isFinalized != it.isFinalized {
		return false
	}

	if wrapper.GetFinalizePath() != it.GetFinalizePath() {
		return false
	}

	if wrapper.separator != it.separator {
		return false
	}

	if wrapper.finalizedError != nil && it.finalizedError != nil {
		if !wrapper.finalizedError.IsEquals(it.finalizedError) {
			return false
		}
	}

	if wrapper.finalizedError == nil || it.finalizedError == nil {
		return false
	}

	return it.collection.
		IsEqualsPtr(
			wrapper.collection)
}

func (it *Wrapper) Finalize() *errstr.Result {
	if it.isFinalized == true {
		// done
		return &errstr.Result{
			Value:        it.finalPath,
			ErrorWrapper: it.finalizedError,
		}
	}

	// set finalize error
	it.isFinalized = true
	it.finalizedError = errorwrapper.
		NewPtr(
			errtype.FinalizedResourceCannotAccess)

	finalPath := it.
		collection.
		Join(it.separator)

	it.finalPath = normalize.PathUsingSeparatorIf(
		true,
		true,
		true,
		it.separator,
		finalPath,
	)

	return errstr.New.Result.ValueOnly(
		it.finalPath)
}

func (it *Wrapper) GetFinalizePath() *errstr.Result {
	if !it.isFinalized {
		// not finalized
		return &errstr.Result{
			Value:        it.finalPath,
			ErrorWrapper: errorwrapper.NewPtr(errtype.CompileFailed),
		}
	}

	return &errstr.Result{
		Value:        it.finalPath,
		ErrorWrapper: nil,
	}
}

func (it *Wrapper) GetFilesOnPath(
	isNormalize bool,
) *errstr.Results {
	currentPath := it.String()

	return pathgetter.Files(
		isNormalize,
		it.separator,
		currentPath,
	)
}

func (it *Wrapper) GetBaseDirFiles(
	isNormalize bool,
) *errstr.Results {
	currentPath := it.GetBaseDir()

	return pathgetter.Files(
		isNormalize,
		it.separator,
		currentPath,
	)
}

func (it *Wrapper) GetDirectoriesOfBaseDir(
	isNormalize bool,
) *errstr.Results {
	currentPath := it.GetBaseDir()

	return pathgetter.Dirs(
		it.separator,
		currentPath,
		isNormalize)
}

// GetRecursiveFilesOnBaseDir Get Recursive files from the basedir
func (it *Wrapper) GetRecursiveFilesOnBaseDir() *errstr.Results {
	currentPath := it.GetBaseDir()

	return recursivepaths.Files(currentPath)
}

func (it *Wrapper) GetBaseDir() string {
	currentPath := it.String()

	return splitinternal.GetBaseDir(
		currentPath)
}

func (it *Wrapper) GetBaseDirName() string {
	currentPath := it.String()

	return splitinternal.GetBaseDirNameOrEmpty(
		currentPath)
}

func (it *Wrapper) GetBaseDirNames() *[]string {
	currentPath := it.String()

	return splitinternal.GetBaseDirNames(
		currentPath)
}

func (it *Wrapper) Splits() *[]string {
	currentPath := it.String()

	return splitinternal.GetAllSplitsWithSep(
		currentPath,
		it.separator)
}

func (it *Wrapper) SplitsUsing(separator string) *[]string {
	currentPath := it.String()

	return splitinternal.GetAllSplitsWithSep(
		currentPath,
		separator)
}

func (it *Wrapper) GetBaseDirFileInfo() (os.FileInfo, *errorwrapper.Wrapper) {
	currentPath := it.GetBaseDir()
	curFileInfo, err := os.Stat(currentPath)

	if err != nil {
		return curFileInfo,
			errnew.Type.Error(errtype.FileInfo, err)
	}

	return curFileInfo, nil
}

func (it *Wrapper) IsBaseDirExists() bool {
	currentFileInfo, errWrap := it.GetBaseDirFileInfo()

	if errWrap.HasError() {
		return false
	}

	return currentFileInfo.IsDir()
}

func (it *Wrapper) IsFileExists() bool {
	currentFileInfo, errWrap := it.GetFileInfo()

	if errWrap.HasError() {
		return false
	}

	return !currentFileInfo.IsDir()
}

func (it *Wrapper) GetBaseDirInfoResult() *dirinfo.Result {
	baseDir := it.GetBaseDir()

	return dirinfo.New(baseDir)
}

func (it *Wrapper) GetFileInfo() (os.FileInfo, *errorwrapper.Wrapper) {
	filePath := it.String()
	curFileInfo, err := os.Stat(filePath)

	if err != nil {
		return curFileInfo, errnew.Type.Error(errtype.FileInfo, err)
	}

	return curFileInfo, nil
}

func (it *Wrapper) GetFileInfoWrapper() *fileinfo.Wrapper {
	filePath := it.String()

	return fileinfo.New(filePath, it.separator)
}

func (it *Wrapper) GetFileInfoWrappers() *fileinfo.Wrappers {
	filePath := it.String()

	return fileinfo.NewWrappersPtr(
		filePath,
		it.separator,
		false)
}

func (it *Wrapper) Parent() *Wrapper {
	filePath := it.GetBaseDir()

	return NewCap(
		constants.ArbitraryCapacity1,
		it.separator).
		Add(filePath)
}

func (it *Wrapper) Collection() *corestr.Collection {
	return it.collection
}

func (it *Wrapper) ListPtr() *[]string {
	return it.collection.ListPtr()
}

func (it *Wrapper) Strings() []string {
	return *it.collection.ListPtr()
}

func (it *Wrapper) StringsPtr() *[]string {
	return it.collection.ListPtr()
}

func (it *Wrapper) getFinalizedError() *errorwrapper.Wrapper {
	if it.IsFinalized() {
		return it.finalizedError
	}

	return nil
}

func (it *Wrapper) ToWrapperUpto(
	uptoLastIndexMinus int,
	sep string,
	isNormalize bool,
) *Wrapper {
	currPath := it.ToStringUptoLastMinus(
		uptoLastIndexMinus,
		sep,
		isNormalize)

	return New(sep).Add(currPath)
}

func (it *Wrapper) ToStringUptoLastMinus(
	uptoLastIndexMinus int,
	sep string,
	isNormalize bool,
) string {
	if uptoLastIndexMinus < 0 {
		errcore.
			CannotBeNegativeType.
			HandleUsingPanic(
				"uptoLastIndexMinus cannot be negative.",
				uptoLastIndexMinus)
	}

	generatedPath := it.
		collection.
		Take(it.Length() - uptoLastIndexMinus).
		Join(sep)

	generatedPathNext := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		sep,
		generatedPath)

	return generatedPathNext
}

func (it *Wrapper) ToString(
	sep string,
	isNormalize bool,
) string {
	generatedPath := it.
		collection.
		Join(sep)

	generatedPathNext := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		sep,
		generatedPath)

	return generatedPathNext
}

func (it Wrapper) String() string {
	if it.IsFinalized() {
		return it.finalPath
	}

	toStr := it.ToString(
		it.separator,
		true)

	return toStr
}

func (it *Wrapper) GetWindowsPath() string {
	if it.IsFinalized() && it.IsWindowsSeparator() {
		return it.finalPath
	}

	toStr := it.ToString(
		constants.WindowsPathSeparator,
		true)

	return toStr
}

func (it *Wrapper) GetUnixPath() string {
	if it.IsFinalized() && it.IsUnixSeparator() {
		return it.finalPath
	}

	toStr := it.ToString(
		constants.ForwardSlash,
		true)

	return toStr
}

func (it *Wrapper) GetAsPathWrapper() *pathwrapper.Wrapper {
	toStr := it.String()
	pathWrapper := pathwrapper.Wrapper(toStr)

	return &pathWrapper
}

func (it *Wrapper) GetAsPathExt() *pathext.Wrapper {
	toStr := it.String()

	return pathext.NewPtr(toStr)
}
