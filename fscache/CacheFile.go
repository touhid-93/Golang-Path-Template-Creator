package fscache

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coretaskinfo"
	"gitlab.com/evatix-go/core/isany"
	"gitlab.com/evatix-go/enum/strtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/refs"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type CacheFile struct {
	ChmodWrapper                   pathchmod.Wrapper           // user input
	AbsFilePath                    string                      // user input
	IsAcquireLock                  bool                        // user input
	IsCollectReadError             bool                        // user input
	IsCollectWriteError            bool                        // user input
	IsWriteEmptyOnNull             bool                        // user input
	IsRemoveFileBeforeWrite        bool                        // user input
	OnInvalidGenerator             *errfunc.OnInvalidGenerator // generator func
	OnReadActionInvoker            pathhelpercore.InvokerFunc
	OnWriteActionInvoker           pathhelpercore.InvokerFunc
	OnCacheExpireActionInvoker     pathhelpercore.InvokerFunc
	OnInvalidGenerateActionInvoker pathhelpercore.InvokerFunc
	isCompiled                     bool // user input
	cacheData                      interface{}
	cacheJsonResult                *errjson.Result
	compileErr                     *errorwrapper.Wrapper // contains all references
	lazyReferences                 *refs.Collection
}

func (it *CacheFile) FsWriter() *fs.Writer {
	if it == nil {
		return &fs.Writer{}
	}

	return &fs.Writer{
		ChmodWrapper:            it.ChmodWrapper.ToNonPtr(),
		Location:                it.AbsFilePath,
		IsApplyChmodMust:        true,
		IsApplyChmodOnMismatch:  true,
		IsWriteEmptyOnNull:      it.IsWriteEmptyOnNull,
		IsKeepExistingFileChmod: true,
	}
}

func (it *CacheFile) SetChmodWrapper(
	chmodWrapper *pathchmod.Wrapper,
) *CacheFile {
	if it == nil {
		return &CacheFile{
			ChmodWrapper: chmodWrapper.ToNonPtr(),
		}
	}

	it.ChmodWrapper = chmodWrapper.ToNonPtr()

	return it
}

func (it *CacheFile) ReadCacheOrDirectJsonResult() *errjson.Result {
	if it.IsCompiled() {
		return it.cacheJsonResult
	}

	jsonResult, errWrap := it.ReadFromFileAsJsonResult()
	// since cache data is nil thus it is not set for isCompiled.
	// todo
	finalErrWrap := it.wrapErrorWithReferences(
		errWrap)

	it.compileErr = finalErrWrap
	errJson := errjson.New.Result.Create(
		jsonResult,
		finalErrWrap)
	it.cacheJsonResult = errJson

	return errJson
}

func (it *CacheFile) IsOnInvalidGeneratorDefined() bool {
	return it != nil && it.OnInvalidGenerator != nil
}

func (it *CacheFile) IsOnInvalidGeneratorEmpty() bool {
	return it == nil || it.OnInvalidGenerator == nil
}

func (it *CacheFile) GetSetInvalidGeneratorOnEmpty(
	generator *errfunc.OnInvalidGenerator,
) *CacheFile {
	if it.IsOnInvalidGeneratorDefined() {
		return it
	}

	it.OnInvalidGenerator = generator

	return it
}

// GetOnce
//
//  on compile success reflect set to ToPtr
//  on compile error returns previous compiling error
func (it *CacheFile) GetOnce(
	toPtr interface{},
) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isCompiled && it.compileErr.IsEmpty() {
		return errnew.
			Reflect.
			SetFromTo(it.cacheData, toPtr)
	}

	if it.isCompiled && it.compileErr.HasError() {
		return it.compileErr
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	// not generated yet
	// check in file first
	var readFromFileErrWrap *errorwrapper.Wrapper
	isFileExist := it.isFileExistInternal()
	if it.IsFileExist() {
		// read from it
		readFromFileErrWrap = it.readFromFileInternal(
			toPtr)
	}

	if isFileExist && readFromFileErrWrap.IsEmpty() {
		// all success
		return nil
	}

	if it.isCollectReadErrorWrap(isFileExist, readFromFileErrWrap) {
		// error needs to return
		return readFromFileErrWrap
	}

	// may not exist in file or read error
	// generate
	it.IsAcquireLock = it.isLockAcquire()
	it.OnInvalidGenerator.IsLockRequired = false // two locks is not required.

	generateErrWrap := it.OnInvalidGenerator.GenerateTo(toPtr)
	finalErrWrap := it.wrapErrorWithReferences(generateErrWrap)
	it.isCompiled = true
	it.compileErr = finalErrWrap
	it.cacheData = toPtr
	jsonResult := corejson.NewPtr(toPtr)
	it.cacheJsonResult = errjson.New.Result.Create(
		jsonResult,
		finalErrWrap)

	if it.IsOnInvalidateGenerateActionInvokerDefined() {
		it.StateActionInvoke(
			CacheStates.InvalidateGenerateState,
			errbyte.New.Results.Create(
				generateErrWrap,
				it.cacheJsonResult.SafeValues()))
	}

	if generateErrWrap.HasError() {
		return generateErrWrap
	}

	// clear, Save
	savingErrorWrap := it.saveInternal(toPtr)
	if it.IsCollectWriteError && savingErrorWrap.HasError() {
		return savingErrorWrap
	}

	return nil
}

func (it *CacheFile) isCollectReadErrorWrap(
	isFileExist bool,
	readFromFileErrWrap *errorwrapper.Wrapper,
) bool {
	return it.IsCollectReadError &&
		isFileExist &&
		readFromFileErrWrap.HasError()
}

func (it *CacheFile) ReadFromFile(
	toPtr interface{},
) *errorwrapper.Wrapper {
	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if it.isNotFileExistInternal() {
		return it.fileNotExistErrorWrap()
	}

	// file exist
	return it.readFromFileInternal(toPtr)
}

func (it *CacheFile) ReadFromFileAsJsonResult() (
	fileJsonResult *corejson.Result,
	readingErrWrap *errorwrapper.Wrapper,
) {
	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if it.isNotFileExistInternal() {
		return nil, it.fileNotExistErrorWrap()
	}

	// file exist
	readJsonResult := fs.ReadErrorJsonResult(it.AbsFilePath)

	if it.IsReadActionInvokerDefined() {
		it.StateActionInvoke(
			CacheStates.ReadState,
			errbyte.New.Results.Create(
				readingErrWrap.ErrWrap(),
				readJsonResult.SafeValues()))
	}

	return readJsonResult.Result, readJsonResult.ErrorWrapper
}

func (it *CacheFile) StateActionInvoke(
	state strtype.Variant,
	fileRawBytesResults *errbyte.Results,
) *errorwrapper.Wrapper {
	switch state {
	case CacheStates.ReadState:
		if it.IsReadActionInvokerDefined() {
			return it.OnReadActionInvoker(
				it.FileInfo(CacheStates.ReadState, fileRawBytesResults))
		}
	case CacheStates.WriteState:
		if it.IsWriteActionInvokerDefined() {
			return it.OnWriteActionInvoker(
				it.FileInfo(CacheStates.WriteState, fileRawBytesResults))
		}
	case CacheStates.CacheExpireState:
		if it.IsExpireActionInvokerDefined() {
			return it.OnCacheExpireActionInvoker(
				it.FileInfo(CacheStates.CacheExpireState, fileRawBytesResults))
		}
	case CacheStates.InvalidateGenerateState:
		if it.IsOnInvalidateGenerateActionInvokerDefined() {
			return it.OnInvalidGenerateActionInvoker(
				it.FileInfo(CacheStates.InvalidateGenerateState, fileRawBytesResults))
		}
	}

	return errnew.NotSupportedOption("CacheState", state)
}

func (it *CacheFile) IsReadActionInvokerDefined() bool {
	return it != nil && it.OnReadActionInvoker != nil
}

func (it *CacheFile) IsWriteActionInvokerDefined() bool {
	return it != nil && it.OnWriteActionInvoker != nil
}

func (it *CacheFile) IsExpireActionInvokerDefined() bool {
	return it != nil && it.OnCacheExpireActionInvoker != nil
}

func (it *CacheFile) IsOnInvalidateGenerateActionInvokerDefined() bool {
	return it != nil && it.OnInvalidGenerateActionInvoker != nil
}

func (it *CacheFile) IsReadActionInvokerEmpty() bool {
	return !it.IsReadActionInvokerDefined()
}

func (it *CacheFile) IsWriteActionInvokerEmpty() bool {
	return !it.IsWriteActionInvokerDefined()
}

func (it *CacheFile) IsExpireActionInvokerEmpty() bool {
	return !it.IsExpireActionInvokerDefined()
}

func (it *CacheFile) IsOnInvalidateGenerateActionInvokerEmpty() bool {
	return !it.IsOnInvalidateGenerateActionInvokerDefined()
}

func (it *CacheFile) FileInfo(
	stateType strtype.Variant,
	fileRawBytesResults *errbyte.Results,
) *pathhelpercore.FileInfo {
	return &pathhelpercore.FileInfo{
		StateType:           stateType,
		RootInfo:            it.RootNameInfo(),
		FilePath:            it.AbsFilePath,
		FileRawBytesResults: fileRawBytesResults,
	}
}

func (it *CacheFile) fileNotExistErrorWrap() *errorwrapper.Wrapper {
	return errnew.Ref.TypeInfoMsgRefs(
		errtype.FileNotExist,
		it.RootNameInfo(),
		"file not exist to read or unmarshal",
		it.References().Items()...)
}

func (it *CacheFile) References() *refs.Collection {
	if it == nil {
		return nil
	}

	if it.lazyReferences != nil {
		return it.lazyReferences
	}

	references := refs.NewUsingInfo(
		it.RootNameInfo())

	references.Add("AbsFilePath", it.AbsFilePath)
	references.Adds(it.ChmodWrapper.References()...)
	it.lazyReferences = references

	return it.lazyReferences
}

func (it *CacheFile) RootNameInfo() *coretaskinfo.Info {
	if it == nil {
		return nil
	}

	return it.OnInvalidGenerator.NameInfo
}

func (it *CacheFile) readFromFileInternal(
	toPtr interface{},
) *errorwrapper.Wrapper {
	// file exist
	readErrWrap := fs.JsonReadUnmarshal(
		it.AbsFilePath,
		toPtr)

	return it.wrapErrorWithReferences(
		readErrWrap)
}

func (it *CacheFile) wrapErrorWithReferences(
	readErrWrap *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if readErrWrap.IsEmpty() {
		return nil
	}

	return readErrWrap.
		ConcatNew().
		MsgRefsOnly(it.References().Items()...)
}

func (it *CacheFile) IsFileExist() bool {
	if it.IsAcquireLock {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.isFileExistInternal()
}

func (it *CacheFile) IsCacheIntegrityAlright(
	toPtr interface{},
) bool {
	if it.IsAcquireLock {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.isCacheIntegrityAlrightInternal(toPtr)
}

func (it *CacheFile) isFileExistInternal() bool {
	return chmodhelper.IsPathExists(it.AbsFilePath)
}

func (it *CacheFile) isNotFileExistInternal() bool {
	return !chmodhelper.IsPathExists(it.AbsFilePath)
}

func (it *CacheFile) IsNotFileExist() bool {
	return !chmodhelper.IsPathExists(it.AbsFilePath)
}

func (it *CacheFile) IsCompiled() bool {
	return it != nil && it.isCompiled
}

func (it *CacheFile) IsCompiledSafe() bool {
	return it.IsCompiled() && it.compileErr.IsEmpty()
}

func (it *CacheFile) InvalidateCacheFile() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.invalidateCacheInternal()
}

func (it *CacheFile) invalidateCacheInternal() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if !it.IsCompiled() {
		// only remove cache file
		return it.expireCacheFileInternal()
	}

	// invalidate props
	it.invalidateCachePropertiesForReloading()

	return it.expireCacheFileInternal()
}

func (it *CacheFile) invalidateCachePropertiesForReloading() {
	// remove data cache
	it.cacheData = nil
	it.compileErr = nil
	it.cacheJsonResult = nil
	it.isCompiled = false
}

func (it *CacheFile) RemoveFileCache() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.expireCacheFileInternal()
}

func (it *CacheFile) ExpireCacheFile() *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.expireCacheFileInternal()
}

func (it *CacheFile) isLockAcquire() bool {
	return it != nil &&
		it.IsAcquireLock ||
		it.OnInvalidGenerator.IsLockRequired
}

func (it *CacheFile) expireCacheFileInternal() *errorwrapper.Wrapper {
	removeErrWrap := deletepaths.SingleOnExist(it.AbsFilePath)

	return it.wrapErrorWithReferences(
		removeErrWrap)
}

func (it *CacheFile) IsApplicable() bool {
	return it != nil &&
		it.isCompiled &&
		it.compileErr.IsEmpty()
}

func (it *CacheFile) IsErrorOnNull() bool {
	return it != nil &&
		!it.IsWriteEmptyOnNull
}

// Save
//
//  Casting happens:
//  - self or self pointer returns directly
//  - []Bytes to Result
//  - string (json) to Result
//  - Jsoner to Result
//  - bytesSerializer to Result
//  - error to Result
//  - AnyItem
func (it *CacheFile) Save(
	fromAny interface{},
) *errorwrapper.Wrapper {
	if it == nil {
		return errnew.Null.Simple(it)
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.saveInternal(fromAny)
}

// saveInternal
//
//  no lock
//
//  Casting happens:
//  - self or self pointer returns directly
//  - []Bytes to Result
//  - string (json) to Result
//  - Jsoner to Result
//  - bytesSerializer to Result
//  - error to Result
//  - AnyItem
func (it *CacheFile) saveInternal(
	fromAnyItem interface{},
) *errorwrapper.Wrapper {
	isNull := isany.Null(fromAnyItem)
	isExitOnNull := it.IsErrorOnNull() && isNull

	if isExitOnNull {
		cannotWriteErrWrap := errnew.Null.WithMessage(
			"cannot Save cache data on nil given",
			fromAnyItem)

		it.StateActionInvoke(
			CacheStates.WriteState,
			errbyte.Empty.ResultsWithError(cannotWriteErrWrap))
	}

	// can be null
	if isNull {
		// on null write empty
		return it.writeToFile([]byte(""))
	}

	toJsonResult := corejson.
		AnyTo.
		SerializedJsonResult(
			fromAnyItem)

	if toJsonResult == nil || toJsonResult.HasError() {
		serializeErrWrap := errnew.Error.Default(
			errtype.Serialize,
			toJsonResult.MeaningfulError())

		it.StateActionInvoke(
			CacheStates.WriteState,
			errbyte.Empty.ResultsWithError(serializeErrWrap))

		return serializeErrWrap
	}

	if it.IsRemoveFileBeforeWrite {
		// swallow is fine here.
		removeErr := os.RemoveAll(it.AbsFilePath)
		removeErrWrap := errnew.Path.Error(
			errtype.RemoveFailed,
			removeErr,
			it.AbsFilePath)
		// next error will be caught in write.
		it.StateActionInvoke(
			CacheStates.CacheExpireState,
			errbyte.New.Results.Create(
				removeErrWrap,
				toJsonResult.SafeValues()))
	}

	return it.writeToFile(toJsonResult.Bytes)
}

func (it *CacheFile) writeToFile(rawBytes []byte) *errorwrapper.Wrapper {
	writeErrWrap := fs.WriteAllParams(
		true,
		true,
		true,
		it.ChmodWrapper.IsSkipOnInvalid,
		it.ChmodWrapper.IsKeepExistingChmod,
		it.ChmodWrapper.DirChmod,
		it.ChmodWrapper.FileChmod,
		it.AbsFilePath,
		rawBytes)

	it.StateActionInvoke(
		CacheStates.WriteState,
		errbyte.New.Results.Create(
			writeErrWrap,
			rawBytes))

	it.invalidateCachePropertiesForReloading()

	return writeErrWrap
}

func (it *CacheFile) isCacheIntegrityAlrightInternal(
	toPtr interface{},
) bool {
	if it.isNotFileExistInternal() {
		return false
	}

	errWrap := it.readFromFileInternal(
		toPtr)

	return errWrap.IsEmpty()
}
