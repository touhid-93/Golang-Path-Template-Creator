package pathext

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/core/extensionsconst"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

type Wrapper struct {
	fullPath                 string
	fileInfo                 *os.FileInfo
	fileInfoError            *errorwrapper.Wrapper
	dotExtension, extension  *string
	filteringExt             *string
	dotIndex                 *int
	fileNameWithExtension    *string
	baseDir                  *string
	fileNameWithoutExtension *string
}

func New(path string) Wrapper {
	return Wrapper{
		fullPath: path,
	}
}

func NewPtr(path string) *Wrapper {
	return &Wrapper{
		fullPath: path,
	}
}

func (it *Wrapper) IsPathEquals(path string, ignoreCase bool) bool {
	if ignoreCase {
		return strings.EqualFold(it.fullPath, path)
	}

	return it.fullPath == path
}

func (it *Wrapper) IsPathContains(path string) bool {
	return strings.Contains(it.fullPath, path)
}

func (it *Wrapper) ExtDotIndex() int {
	if it.dotIndex != nil {
		return *it.dotIndex
	}

	invalid := constants.InvalidNotFoundCase
	p := it.fullPath
	if p == "" {
		it.dotIndex = &invalid

		return *it.dotIndex
	}

	// doesn't look good on-line break
	for i := len(p) - 1; i >= 0 && !(p[i] == constants.BackwardChar || p[i] == constants.ForwardChar); i-- {
		if p[i] == constants.DotChar {
			it.dotIndex = &i

			return i
		}
	}

	it.dotIndex = &invalid

	return *it.dotIndex
}

func (it *Wrapper) BaseDir() string {
	if it.baseDir != nil {
		return *it.baseDir
	}

	it.initializeProperties()

	return *it.baseDir
}

func (it *Wrapper) FileNameWithExtension() string {
	if it.fileNameWithExtension != nil {
		return *it.fileNameWithExtension
	}

	it.initializeProperties()

	return *it.fileNameWithExtension
}

func (it *Wrapper) FileNameWithoutExtension() string {
	if it.fileNameWithoutExtension != nil {
		return *it.fileNameWithoutExtension
	}

	it.initializeProperties()

	return *it.fileNameWithoutExtension
}

func (it *Wrapper) initializeProperties() {
	if it.baseDir != nil {
		return
	}

	baseDir, fileNameWithExtension := splitinternal.GetWithoutSlash(
		it.fullPath)
	fileNameWithoutExt := fileNameWithExtension

	if it.HasExtension() {
		fileNameWithoutExt = strings.Replace(
			fileNameWithExtension,
			*it.DotExtension(),
			"",
			1)
	}

	it.fileNameWithExtension = &fileNameWithExtension
	it.baseDir = &baseDir
	it.fileNameWithoutExtension = &fileNameWithoutExt
}

// FilteringExt is ext and one char more from left.
// Panics if char is not there
func (it *Wrapper) FilteringExt() string {
	if it.filteringExt != nil {
		return *it.filteringExt
	}

	filteringExt := it.GetMoreThanExt(1)
	it.filteringExt = &filteringExt

	return *it.filteringExt
}

func (it *Wrapper) FileInfoWrapper() *os.FileInfo {
	if it.fileInfo != nil ||
		it.fileInfoError != nil &&
			it.fileInfoError.HasError() {
		return it.fileInfo
	}

	fileInfo, err :=
		os.Stat(it.fullPath)

	if err != nil {
		it.fileInfoError =
			errorwrapper.NewPath(
				codestack.SkipNone,
				errtype.FileInvalidOrMissing,
				err,
				it.fullPath)
	} else {
		it.fileInfo = &fileInfo
	}

	return it.fileInfo
}

func (it *Wrapper) IsFile() bool {
	fileInfo := it.FileInfoWrapper()
	err := it.fileInfoError

	if err != nil && err.HasError() {
		return false
	}

	return fileInfo != nil &&
		!(*fileInfo).IsDir()
}

func (it *Wrapper) IsDir() bool {
	fileInfo := it.FileInfoWrapper()
	err := it.fileInfoError

	if err != nil && err.HasError() {
		return false
	}

	return fileInfo != nil &&
		(*fileInfo).IsDir()
}

func (it *Wrapper) GetMoreThanExt(moreIndex int) string {
	dotIndex := it.ExtDotIndex()

	if dotIndex == constants.InvalidValue {
		return ""
	}

	newIndex := dotIndex - moreIndex

	return it.fullPath[newIndex:]
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (it *Wrapper) DotExtension() *string {
	if it.dotExtension == nil {
		dotExt := it.GetMoreThanExt(
			constants.Zero)
		it.dotExtension = &dotExt
	}

	return it.dotExtension
}

func (it *Wrapper) HasExtension() bool {
	return it.ExtDotIndex() > -1
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (it *Wrapper) Extension() *string {
	if it.extension != nil {
		return it.extension
	}

	dotExt := *it.DotExtension()

	if len(dotExt) > 0 && dotExt[0] == constants.Dot[0] {
		ext := dotExt[1:]
		it.extension = &ext
	} else {
		it.extension = &dotExt
	}

	return it.extension
}

func (it *Wrapper) IsExtension(extension string) bool {
	return *it.Extension() == extension
}

func (it *Wrapper) IsExtOrDotExt(extOrDotExt string) bool {
	return *it.Extension() == extOrDotExt ||
		*it.DotExtension() == extOrDotExt
}

func (it *Wrapper) IsDotExtension(dotExtension string) bool {
	return *it.DotExtension() == dotExtension
}

func (it *Wrapper) IsAnyOfExtension(extensions ...string) bool {
	if extensions == nil {
		return false
	}

	return it.IsAnyOfExtensionPtr(&extensions)
}

func (it *Wrapper) IsAnyOfExtensionPtr(
	extensions *[]string,
) bool {
	if extensions == nil || len(*extensions) == 0 {
		return false
	}

	currentExt := *it.Extension()
	for _, ext := range *extensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}

func (it *Wrapper) IsAnyOfDotExtension(
	dotExtensions ...string,
) bool {
	if dotExtensions == nil {
		return false
	}

	return it.IsAnyOfDotExtensionPtr(&dotExtensions)
}

func (it *Wrapper) IsAnyOfDotExtensionPtr(
	dotExtensions *[]string,
) bool {
	if dotExtensions == nil || len(*dotExtensions) == 0 {
		return false
	}

	currentExt := *it.DotExtension()
	for _, ext := range *dotExtensions {
		if currentExt == ext {
			return true
		}
	}

	return false
}

func (it *Wrapper) IsExtensionFilterMatch(
	extensionFilter string,
) bool {
	if extensionFilter == extensionsconst.AllFiles {
		return true
	}

	length := len(extensionFilter)
	if length == 0 {
		return false
	}

	currentDotExt := *it.DotExtension()

	if length == 2 {
		// *.
		firstChar := extensionFilter[coreindexes.First]
		secondChar := extensionFilter[coreindexes.Second]

		// a file name ends with dot.
		isFileEndsWithDot := firstChar == constants.AsteriskChar &&
			secondChar == constants.DotChar &&
			currentDotExt == constants.Dot

		return isFileEndsWithDot ||
			firstChar == constants.DotChar &&
				currentDotExt == extensionFilter
	}

	if length < 3 {
		return false
	}

	firstChar := extensionFilter[coreindexes.First]
	secondChar := extensionFilter[coreindexes.Second]

	// *.ext
	if firstChar == constants.AsteriskChar &&
		secondChar == constants.DotChar {
		dotExt := extensionFilter[coreindexes.Second:]

		if dotExt == currentDotExt {
			return true
		}
	}

	// .whatever
	if firstChar == constants.DotChar &&
		extensionFilter == currentDotExt {
		return true
	}

	// {don't care}what.ever, ends with extension
	return stringutil.IsEndsWith(
		it.fullPath,
		extensionFilter,
		true)
}

func (it *Wrapper) IsExtensionFiltersMatch(
	extensionsFilter *[]string,
	extensionsLength int,
) bool {
	if extensionsLength == 0 {
		return false
	}

	for _, extFilter := range *extensionsFilter {
		if it.IsExtensionFilterMatch(extFilter) {
			return true
		}
	}

	return false
}

func (it *Wrapper) IsNameWithExtensionMatches(
	fullPath string,
) bool {
	_, fileName := splitinternal.Get(fullPath)

	return stringutil.IsEndsWith(
		it.fullPath,
		fileName,
		true)
}
