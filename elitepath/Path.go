package elitepath

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
	"gitlab.com/evatix-go/pathhelper/deletepaths"
	"gitlab.com/evatix-go/pathhelper/fileinfopath"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompareinternal"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathfixer"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
	"gitlab.com/evatix-go/pathhelper/pathstatlinux"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
	"gitlab.com/evatix-go/pathhelper/pathwrapper"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

type Path struct {
	pathfixer.Location
	fixedPathSlice []string
}

func NewPathExpandNormalizeDefault(path string) *Path {
	return &Path{
		Location: *pathfixer.NewLocationUsingOptions(path, pathfixer.PathOptions{
			IsNormalize:    true,
			IsExpandEnvVar: true,
		}),
	}
}

func NewPathDefault(path string) *Path {
	return &Path{
		Location: *pathfixer.NewLocation(path),
	}
}

func (it *Path) Join(location string) string {
	return pathjoin.JoinSimple(
		it.CompiledPath(),
		location)
}

func (it *Path) Join2(location1, location2 string) string {
	return pathjoin.JoinSimpleThree(
		it.CompiledPath(),
		location1,
		location2)
}

func (it *Path) Joins(isNormalize bool, locations ...string) string {
	return pathjoin.JoinFixedIf(
		isNormalize,
		it.CompiledPath(),
		locations...)
}

func (it *Path) JoinsToPath(
	relativePaths ...string,
) *Path {
	combinedPath := it.Combine(relativePaths...)

	return it.ClonePathUsingNew(
		combinedPath)
}

func (it *Path) Combine(relativePaths ...string) string {
	if it == nil {
		return pathjoin.JoinBaseDirWithSep(
			true,
			false,
			false,
			osconsts.PathSeparator,
			"",
			relativePaths...)
	}

	return pathjoin.JoinBaseDirWithSep(
		true,
		it.IsExpandEnvVar,
		it.IsNormalize,
		osconsts.PathSeparator,
		it.Location.Path,
		relativePaths...)
}

func (it *Path) enhancePathsToSliceStrings(enhancePaths ...*Path) []string {
	slice := make(
		[]string,
		0,
		len(enhancePaths))

	if len(enhancePaths) == 0 {
		return slice
	}

	for _, enhancePath := range enhancePaths {
		if enhancePath == nil {
			continue
		}

		slice = append(
			slice,
			enhancePath.CompiledPath())
	}

	return slice
}

func (it *Path) RwxWrapper() *pathchmod.RwxWrapperWithError {
	return pathchmod.ExistingRwxWrapperWithError(it.CompiledPath())
}

// OsFile use fs.Flag to use appropriate file flags
//
// Must call fs.OsFile.AttachDeferCloseOnRequire()
func (it *Path) OsFile(
	existingErrorWrapper *errorwrapper.Wrapper, // can be nil
	osFlag int,
	fileMode os.FileMode,
) *fs.OsFile {
	return fs.GetOsFile(
		existingErrorWrapper,
		osFlag,
		fileMode,
		it.CompiledPath())
}

func (it *Path) CombineWithElitePaths(
	enhancePaths ...*Path,
) string {
	relativePaths := it.enhancePathsToSliceStrings(
		enhancePaths...)

	return it.Combine(relativePaths...)
}

func (it *Path) CombineWithElitePathsToElitePath(
	enhancePaths ...*Path,
) *Path {
	finalPath := it.CombineWithElitePaths(
		enhancePaths...)

	return it.ClonePathUsingNew(finalPath)
}

func (it *Path) IsFilterMatch(filter *Filter) bool {
	return filter.IsMatch(it)
}

func (it *Path) DeletePath(isSkipOnNonExist bool) *errorwrapper.Wrapper {
	if isSkipOnNonExist {
		return deletepaths.RecursiveOnExist(it.CompiledPath())
	}

	return deletepaths.Recursive(it.CompiledPath())
}

func (it *Path) Move(toPath string) *errorwrapper.Wrapper {
	err := os.Rename(it.CompiledPath(), toPath)

	return errnew.
		Path.
		Error(errtype.PathMove, err, toPath)
}

func (it *Path) SimpleStat() *pathchmod.SimpleStat {
	if it.IsEmptyPath() {
		return nil
	}

	return pathchmod.GetSimpleStat(it.CompiledPath())
}

func (it *Path) PathLinuxStat() *pathstatlinux.Info {
	if it.IsEmptyPath() {
		return pathstatlinux.Get(
			constants.EmptyString)
	}

	return pathstatlinux.Get(it.CompiledPath())
}

func (it *Path) CheckSummer(
	isAsync,
	isRecursive bool,
	hashMethod hashas.Variant,
) *checksummer.Instance {
	src := it.CompiledPath()
	if it.IsEmptyPath() {
		return checksummer.Invalid(
			false,
			src,
			hashMethod,
			errcore.InvalidEmptyPathType.ErrorNoRefs(src))
	}

	return checksummer.New(
		isAsync,
		isRecursive,
		src,
		hashMethod)
}

func (it *Path) IsChecksumEqual(
	isRecursive bool,
	hashMethod hashas.Variant,
	comparingPath string,
) bool {
	src := it.CompiledPath()
	if it.IsEmptyPath() && comparingPath == "" {
		return true
	}

	if it.IsEmptyPath() || comparingPath == "" {
		return false
	}

	sourceCheckSummer := checksummer.New(
		true,
		isRecursive,
		src,
		hashMethod)

	destinationCheckSummer := checksummer.New(
		true,
		isRecursive,
		comparingPath,
		hashMethod)

	return sourceCheckSummer.IsEqual(
		true,
		destinationCheckSummer)
}

func (it *Path) VerifyCheckSumTreeError(
	isRecursive bool,
	hashMethod hashas.Variant,
	comparingPath string,
) *errwrappers.Collection {
	src := it.CompiledPath()
	if it.IsEmptyPath() {
		return errwrappers.NewCap1().AddWrapperPtr(errnew.Path.Empty())
	}

	sourceCheckSummer := checksummer.New(
		true,
		isRecursive,
		src,
		hashMethod)

	destinationCheckSummer := checksummer.New(
		true,
		isRecursive,
		comparingPath,
		hashMethod)

	return sourceCheckSummer.VerifyError(
		true,
		true,
		destinationCheckSummer)
}

func (it *Path) PathWrapper() pathwrapper.Wrapper {
	if it.IsEmptyPath() {
		return constants.EmptyString
	}

	return pathwrapper.Wrapper(it.CompiledPath())
}

func (it *Path) PathExtWrapper() *pathext.Wrapper {
	if it.IsEmptyPath() {
		return pathext.NewPtr(
			constants.EmptyString)
	}

	return pathext.NewPtr(it.CompiledPath())
}

func (it *Path) FileNameWithExt() string {
	return splitinternal.GetName(
		it.CompiledPath())
}

func (it *Path) FileNameWithoutExt() string {
	return splitinternal.GetFileNameWithoutExt(
		it.CompiledPath())
}

func (it *Path) BothExt() (dotExt, ext string) {
	return splitinternal.GetBothExtension(
		it.CompiledPath())
}

func (it *Path) IsExtension(extCompare string) bool {
	_, ext := it.BothExt()

	return strings.EqualFold(ext, extCompare)
}

func (it *Path) IsDotExtension(dotExtCompare string) bool {
	dotExt, _ := it.BothExt()

	return strings.EqualFold(
		dotExt,
		dotExtCompare)
}

func (it *Path) ParentDirPath() *Path {
	return it.ClonePathUsingNew(
		splitinternal.GetBaseDir(
			it.CompiledPath()))
}

func (it *Path) ParentDir() string {
	return splitinternal.GetBaseDir(
		it.CompiledPath())
}

func (it *Path) BaseDir() string {
	return splitinternal.GetBaseDir(
		it.CompiledPath())
}

// AllPaths returns all immediate paths but not nested or recursive paths.
func (it *Path) AllPaths() *errstr.Results {
	pathWrapper := it.PathWrapper()

	return pathWrapper.GetAllPathsDefault()
}

func (it *Path) AllPathsSimpleStat() (*pathchmod.SimpleStats, *errorwrapper.Wrapper) {
	if it.IsFile() {
		return pathchmod.
				NewSimpleStats(1).
				Add(it.CompiledPath()),
			nil
	}

	allPaths := it.AllPaths()

	if allPaths.HasError() {
		return pathchmod.NewSimpleStats(0),
			allPaths.ErrorWrapper
	}

	return pathchmod.NewSimpleStatsUsingItems(
			allPaths.SafeValues()...),
		nil
}

func (it *Path) AllFilesSimpleStat() (*pathchmod.SimpleStats, *errorwrapper.Wrapper) {
	if it.IsFile() {
		return pathchmod.
				NewSimpleStats(1).
				Add(it.CompiledPath()),
			nil
	}

	files := it.Files()

	if files.HasError() {
		return pathchmod.NewSimpleStats(0),
			files.ErrorWrapper
	}

	return pathchmod.NewSimpleStatsUsingItems(
			files.SafeValues()...),
		nil
}

func (it *Path) AllDirsSimpleStat() (*pathchmod.SimpleStats, *errorwrapper.Wrapper) {
	paths := it.Directories()

	if paths.HasError() {
		return pathchmod.NewSimpleStats(0),
			paths.ErrorWrapper
	}

	return pathchmod.NewSimpleStatsUsingItems(
			paths.SafeValues()...),
		nil
}

// Files it doesn't return recursive files but just immediate nested files
func (it *Path) Files() *errstr.Results {
	if it.IsFile() {
		return errstr.New.Results.SpreadValuesOnly(it.CompiledPath())
	}

	pathWrapper := it.PathWrapper()

	return pathWrapper.GetFilesDefault()
}

// Directories it doesn't return recursive directories but just immediate nested directories
func (it *Path) Directories() *errstr.Results {
	pathWrapper := it.PathWrapper()

	return pathWrapper.GetDirectoriesDefault()
}

func (it *Path) LocationInfo() *pathhelper.LocationInfo {
	if it.IsEmptyPath() {
		return &pathhelper.LocationInfo{
			RawLocation:           "",
			FileNameWithExtension: "",
			BaseDir:               "",
			FileName:              "",
			DotExtension:          "",
			Extension:             "",
		}
	}

	return pathhelper.GetLocationInfo(
		it.CompiledPath())
}

func (it *Path) NotDirError() *errorwrapper.Wrapper {
	if it.IsDir() {
		return nil
	}

	return errnew.
		Path.
		Messages(
			errtype.InvalidDir,
			it.CompiledPath(),
			"not a valid directory")
}

func (it *Path) NotFileError() *errorwrapper.Wrapper {
	if it.IsFile() {
		return nil
	}

	return errnew.
		Path.
		Messages(
			errtype.FileInvalid,
			it.CompiledPath(),
			"not a valid file")
}

func (it *Path) ReadFileBytesMust() []byte {
	rs := it.ReadFileBytes()
	rs.ErrorWrapper.HandleError()

	return rs.SafeValues()
}

func (it *Path) ReadFileBytes() *errbyte.Results {
	return fs.ReadFileUsingLock(it.CompiledPath())
}

func (it *Path) ReadFileStringMust() string {
	rs := it.ReadFileString()
	rs.ErrorWrapper.HandleError()

	return rs.Value
}

func (it *Path) ReadFileString() *errstr.Result {
	return fs.ReadFileStringUsingLock(it.CompiledPath())
}

func (it *Path) ReadLinesMust() []string {
	rs := it.ReadLines()
	rs.ErrorWrapper.HandleError()

	return rs.SafeValues()
}

func (it *Path) ReadLines() *errstr.Results {
	return fs.ReadFileLinesUsingLock(it.CompiledPath())
}

func (it *Path) WriteLines(lines []string) *errorwrapper.Wrapper {
	return fs.WriteStringLinesToFileUsingLock(
		true,
		it.CompiledPath(),
		lines)
}

func (it *Path) WriteJsonResult(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	return fs.WriteJsonResult(
		true,
		false,
		jsonResult,
		it.CompiledPath(),
	)
}

func (it *Path) WriteJsonResultWithoutChecking(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	return fs.WriteJsonResultWithoutChecking(
		jsonResult,
		it.CompiledPath(),
	)
}

func (it *Path) RecursivePathsAll(
	isRelativePath bool,
	isExcludeRootName bool,
	excludeRootNames ...string,
) *corestr.SimpleSlice {
	if !it.IsPathExist() {
		return corestr.Empty.SimpleSlice()
	}

	src := it.CompiledPath()

	instruction := pathrecurseinfo.Instruction{
		Root:               src,
		ExcludingRootNames: excludeRootNames,
		IsIncludeFilesOnly: false,
		IsRelativePath:     isRelativePath,
		IsIncludeDirsOnly:  false,
		IsIncludeAll:       true,
		IsExcludeRoot:      isExcludeRootName,
		IsRecursive:        true,
		IsNormalize:        false,
	}

	return instruction.
		Result().
		PathsResult.
		ExpandingPaths
}

func (it *Path) RecursiveFilePaths(
	isRelativePath bool,
	excludeRootNames ...string,
) *corestr.SimpleSlice {
	src := it.CompiledPath()

	instruction := pathrecurseinfo.Instruction{
		Root:               src,
		ExcludingRootNames: excludeRootNames,
		IsIncludeFilesOnly: true,
		IsRelativePath:     isRelativePath,
		IsIncludeDirsOnly:  false,
		IsIncludeAll:       false,
		IsExcludeRoot:      false,
		IsRecursive:        true,
		IsNormalize:        false,
	}

	return instruction.
		Result().
		PathsResult.
		ExpandingPaths
}

func (it *Path) RecursiveDirPaths(
	isRelativePath bool,
	excludeRootNames ...string,
) *corestr.SimpleSlice {
	src := it.CompiledPath()

	instruction := pathrecurseinfo.Instruction{
		Root:               src,
		ExcludingRootNames: excludeRootNames,
		IsIncludeFilesOnly: false,
		IsRelativePath:     isRelativePath,
		IsIncludeDirsOnly:  true,
		IsIncludeAll:       false,
		IsExcludeRoot:      false,
		IsRecursive:        true,
		IsNormalize:        false,
	}

	return instruction.
		Result().
		PathsResult.
		ExpandingPaths
}

func (it *Path) DirFilesPaths(isRecursive bool) *errstr.Results {
	return recursivepaths.FilesPlusDirsByName(
		isRecursive,
		it.CompiledPath())
}

func (it *Path) CopyTo(
	isClearBefore,
	isRecursive bool,
	toPath string,
) *errorwrapper.Wrapper {
	src := it.CompiledPath()

	return copyrecursive.DoOptions(src, toPath, copyrecursive.Options{
		IsSkipOnExist:      false,
		IsRecursive:        isRecursive,
		IsMove:             false,
		IsClearDestination: isClearBefore,
		IsUseShellOrCmd:    true,
		IsNormalize:        false,
		IsExpandVar:        false,
	})
}

func (it *Path) MoveTo(
	isClearBefore,
	isRecursive bool,
	toPath string,
) *errorwrapper.Wrapper {
	src := it.CompiledPath()

	return copyrecursive.DoOptions(src, toPath, copyrecursive.Options{
		IsSkipOnExist:      false,
		IsRecursive:        isRecursive,
		IsMove:             true,
		IsClearDestination: isClearBefore,
		IsUseShellOrCmd:    true,
		IsNormalize:        false,
		IsExpandVar:        false,
	})
}

func (it *Path) CopyChmod(isSkipOnWidows bool, toPath string) *errorwrapper.Wrapper {
	if isSkipOnWidows && osconsts.IsWindows {
		return nil
	}

	return fs.CopyChmod(it.CompiledPath(), toPath)
}

func (it *Path) CopyChown(isSkipOnWidows bool, toPath string) *errorwrapper.Wrapper {
	if isSkipOnWidows && osconsts.IsWindows {
		return nil
	}

	return pathsysinfo.ChownCopy(it.CompiledPath(), toPath)
}

func (it *Path) CopyChmodChown(isSkipOnWidows bool, toPath string) *errorwrapper.Wrapper {
	if isSkipOnWidows && osconsts.IsWindows {
		return nil
	}

	chmodErr := fs.CopyChmod(
		it.CompiledPath(),
		toPath)

	if chmodErr.HasError() {
		return chmodErr
	}

	return pathsysinfo.ChownCopy(
		it.CompiledPath(),
		toPath)
}

func (it *Path) ChownPathUserGroupId() *pathsysinfo.PathUserGroupId {
	return pathsysinfo.GetPathUserGroupId(it.CompiledPath())
}

func (it *Path) FileInfoWithPath() *fileinfopath.Instance {
	return fileinfopath.New(it.CompiledPath())
}

func (it *Path) CompareFileInfo(right os.FileInfo) corecomparator.Compare {
	return pathcompareinternal.FileInfoLastModified(it.FileInfo(), right)
}

func (it *Path) CompareSize(anotherInstance *Path) corecomparator.Compare {
	return pathcompareinternal.SizePtr(it.Size(), anotherInstance.Size())
}

func (it *Path) CompareLastModified(anotherInstance *Path) corecomparator.Compare {
	return pathcompareinternal.LastModifiedPtr(it.LastModifiedAt(), anotherInstance.LastModifiedAt())
}

func (it *Path) ChmodCondition() *chmodins.Condition {
	condition := chmodins.Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}

	return &condition
}

func (it *Path) ApplyFileMode(mode os.FileMode) *errorwrapper.Wrapper {
	condition := it.ChmodCondition()

	err := chmodhelper.ChmodApply.PathsUsingFileModeConditions(
		mode,
		condition,
		it.Path)

	return errnew.Type.Error(
		errtype.ChmodApplyFailed,
		err)
}

func (it Path) String() string {
	return it.CompiledPath()
}

func (it *Path) ReadFileUnmarshal(
	unmarshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	return fs.
		JsonReadUnmarshal(
			it.CompiledPath(),
			unmarshallingObjectRef)
}

func (it *Path) ReadJsonParseSelfInjector(
	injector corejson.JsonParseSelfInjector,
) *errorwrapper.Wrapper {
	return fs.
		ReadJsonParseSelfInjector(
			it.CompiledPath(),
			injector)
}

func (it *Path) ApplyRwxInstruction(
	rwx *chmodins.RwxInstruction,
) *errorwrapper.Wrapper {
	if it.IsEmptyPath() || rwx == nil {
		return nil
	}

	return pathchmod.ApplyChmodRwxOwnerGroupOther(
		rwx.IsRecursive,
		rwx.IsSkipOnInvalid,
		rwx.IsContinueOnError,
		&rwx.RwxOwnerGroupOther,
		it.FixedPathAsSlice())
}

func (it *Path) ApplyPathVerifiers(
	errorCollection *errwrappers.Collection,
	pathVerifiers *pathinsfmt.PathVerifiers,
) (isSuccess bool) {
	if it.IsEmptyPath() || pathVerifiers == nil {
		return true
	}

	return pathmodifierverify.ApplyUsingFlatPaths(
		false,
		pathVerifiers,
		errorCollection,
		it.FixedPathAsSlice())
}

func (it *Path) ApplyPathVerifier(
	isSkipOnInvalid bool,
	errorCollection *errwrappers.Collection,
	pathVerifier *pathinsfmt.PathVerifier,
) (isSuccess bool) {
	if it.IsEmptyPath() || pathVerifier == nil {
		return true
	}

	return pathmodifierverify.ApplyVerifierDirect(
		false,
		false,
		isSkipOnInvalid,
		pathVerifier,
		errorCollection,
		it.FixedPathAsSlice()...)
}

func (it *Path) ApplyChown(
	chown *pathinsfmt.Chown,
) *errorwrapper.Wrapper {
	if it.IsEmptyPath() || chown == nil {
		return nil
	}

	return namegroup.Apply(
		chown.IsRecursive,
		false,
		&chown.UserGroupName,
		it.CompiledPath())
}

func (it *Path) FixedPathAsSlice() []string {
	if it.IsEmptyPath() {
		return []string{}
	}

	if it.fixedPathSlice != nil {
		return it.fixedPathSlice
	}

	it.fixedPathSlice = []string{
		it.CompiledPath(),
	}

	return it.fixedPathSlice
}

func (it *Path) ClonePath() *Path {
	if it == nil {
		return nil
	}

	return &Path{
		Location: *it.Location.ClonePath(),
	}
}

func (it *Path) ClonePathUsingNew(newLocation string) *Path {
	if it == nil {
		return &Path{
			Location: *pathfixer.NewLocation(newLocation),
		}
	}

	newPath := Path{
		Location: *it.Location.ClonePath(),
	}

	newPath.Path = newLocation

	return &newPath
}

func (it *Path) LazyPath() *LazyPath {
	return &LazyPath{
		Path: it,
	}
}

func (it *Path) IsEqual(another *Path) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.Location.IsEqual(
		&another.Location)
}

func (it *Path) IsEqualWithoutOptions(another *Path) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.CompiledPath() == another.CompiledPath()
}
