package checksummer

import (
	"bytes"
	"encoding/hex"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type Instance struct {
	rawHashes                     map[string][]byte
	stringHashesMap               map[string]string
	strippedPathWithActualPathMap map[string]string // stripped path key -> actual path
	hashType                      hashas.Variant
	isRecursive                   bool
	isFile                        bool
	rootPath                      string
	rawSingleHash                 []byte
	singleHash                    *string
	stringHashes                  *corestr.SimpleSlice
	hashesHashset                 *corestr.Hashset
	ErrorWrapper                  *errorwrapper.Wrapper
}

func New(
	isAsync,
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) *Instance {
	stat := chmodhelper.
		GetPathExistStat(root)

	if !stat.IsExist {
		return Invalid(
			isRecursive,
			root,
			hashType,
			stat.MeaningFullError())
	}

	isFile := stat.
		IsFile()

	checkSumMap, err := hashAll(
		isAsync,
		isRecursive,
		root,
		hashType)

	errWrap := errnew.
		Path.
		Error(
			errtype.PathMissingOrInvalid,
			err,
			root)

	instance := &Instance{
		rawHashes:    checkSumMap,
		rootPath:     root,
		isRecursive:  isRecursive,
		isFile:       isFile,
		hashType:     hashType,
		ErrorWrapper: errWrap,
	}

	return instance

}

func NewSync(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) *Instance {
	return New(
		false,
		isRecursive,
		root,
		hashType)
}

func Invalid(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
	err error,
) *Instance {
	pathErr := errnew.
		Path.
		Error(
			errtype.PathMissingOrInvalid,
			err,
			root)

	return &Instance{
		hashType:     hashType,
		isRecursive:  isRecursive,
		rootPath:     root,
		ErrorWrapper: pathErr,
	}
}

func NewAsync(
	isRecursive bool,
	root string,
	hashType hashas.Variant,
) *Instance {
	return New(
		true,
		isRecursive,
		root,
		hashType)
}

// HashesHashset Key => HashString (Hex), Value => true
func (it *Instance) HashesHashset() *corestr.Hashset {
	if it.hashesHashset != nil {
		return it.hashesHashset
	}

	slice := it.StringHashes()
	it.hashesHashset = corestr.
		New.Hashset.Strings(
		slice.Items)

	return it.hashesHashset
}

// StringHashesMap Key => Path, Value => HashString (Hex)
func (it *Instance) StringHashesMap() map[string]string {
	if it.stringHashesMap != nil {
		return it.stringHashesMap
	}

	stringHashes := make(
		map[string]string,
		len(it.rawHashes))

	for filename, hashBytes := range it.rawHashes {
		stringHashes[filename] = hex.EncodeToString(hashBytes)
	}

	it.stringHashesMap = stringHashes

	return it.stringHashesMap
}

// StringHashes all the hashes string slice
func (it *Instance) StringHashes() *corestr.SimpleSlice {
	if it.stringHashes != nil {
		return it.stringHashes
	}

	hashMap := it.StringHashesMap()
	slice := corestr.New.SimpleSlice.Cap(it.Length())

	for _, hashString := range hashMap {
		slice.Add(hashString)
	}

	it.stringHashes = slice

	return it.stringHashes
}

func (it *Instance) SingleHash() string {
	if it.singleHash != nil {
		return *it.singleHash
	}

	singleStringHash := hex.EncodeToString(
		it.RawSingleHash())

	it.singleHash = &singleStringHash

	return *it.singleHash
}

// RawSingleHash
// Single
// Bytes hash for all the hashes combined.
//
// Runs lazy
func (it *Instance) RawSingleHash() []byte {
	if it.rawSingleHash != nil {
		return it.rawSingleHash
	}

	hashGen, errWrap := it.hashType.NewHash()
	errWrap.HandleError()

	for _, rawHash := range it.rawHashes {
		hashGen.Write(rawHash)
	}

	it.rawSingleHash = hashGen.Sum(nil)

	return it.rawSingleHash
}

func (it *Instance) HasHashPath(filePath string) bool {
	_, isFound := it.rawHashes[filePath]

	return isFound
}

func (it *Instance) HasHashString(hexHashString string) bool {
	return it.HashesHashset().Has(hexHashString)
}

// HasHashBytes Converts hashBytes to hexHashString then use it. HasHashString
func (it *Instance) HasHashBytes(hashBytes []byte) bool {
	hexHash := hex.EncodeToString(
		hashBytes)

	return it.HasHashString(hexHash)
}

func (it *Instance) GetMap() map[string][]byte {
	return it.rawHashes
}

func (it *Instance) ContainsMap(other map[string][]byte) bool {
	for key, otherHash := range other {
		myHash, isFound := it.rawHashes[key]

		if !isFound {
			return false
		}

		if !bytes.Equal(otherHash, myHash) {
			return false
		}
	}

	return true
}

func (it *Instance) isEqualAllCompare(
	isTrimRoot bool,
	other *Instance,
) bool {
	if it == nil && other == nil {
		return true
	}

	if it == nil || other == nil {
		return false
	}

	if len(it.rawHashes) != len(other.rawHashes) {
		return false
	}

	for otherFile, otherHash := range other.rawHashes {
		myHash, isFound := it.getHash(isTrimRoot, otherFile)

		if !isFound {
			return false
		}

		if !bytes.Equal(myHash, otherHash) {
			return false
		}
	}

	return true
}

func (it *Instance) getHash(
	isPathStripped bool,
	path string,
) (hash []byte, isFound bool) {
	if !isPathStripped {
		fileHash, fileIsFound := it.rawHashes[path]

		return fileHash, fileIsFound
	}

	strippedMap := it.StrippedPathMap()
	actualPath, isPathFound := strippedMap[path]

	if !isPathFound {
		return nil, false
	}

	fileHash, fileIsFound := it.rawHashes[actualPath]

	return fileHash, fileIsFound
}

func (it *Instance) verifyEqualAllCompareStripped(
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	other *Instance,
) *errwrappers.Collection {
	stateTracker := errCollection.StateTracker()

	for otherStrippedFile, otherActualFile := range other.StrippedPathMap() {
		myHash, isFound := it.getHash(true, otherStrippedFile)
		otherHash, otherIsFound := other.getHash(false, otherActualFile)

		if !otherIsFound {
			errCollection.AddPathIssueMessages(
				errtype.FileInvalidOrMissing,
				otherActualFile,
				"file missing from map")
		}

		if !isFound {
			errCollection.AddPathIssueMessages(
				errtype.FileInvalidOrMissing,
				otherActualFile,
				"hash not found in the dictionary or map.")
		}

		if !bytes.Equal(myHash, otherHash) {
			errCollection.AddPathIssueMessages(
				errtype.CheckSumMismatch,
				otherActualFile,
				"hash not found",
				"Src file:",
				pathjoin.JoinNormalized(it.rootPath, otherStrippedFile),
				"Dest file:",
				otherActualFile)
		}

		if !isContinueOnErr && stateTracker.HasChangesCollection() {
			return errCollection
		}
	}

	return errCollection
}

func (it *Instance) verifyEqualAllCompare(
	isPathStripped,
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	other *Instance,
) *errwrappers.Collection {
	if isPathStripped {
		return it.verifyEqualAllCompareStripped(isContinueOnErr, errCollection, other)
	}

	stateTracker := errCollection.StateTracker()

	for otherFile, otherHash := range other.rawHashes {
		myHash, isFound := it.getHash(isPathStripped, otherFile)

		if !isFound {
			errCollection.AddPathIssueMessages(
				errtype.FileInvalidOrMissing,
				otherFile,
				"hash not found in the dictionary or map.")
		}

		if !bytes.Equal(myHash, otherHash) {
			errCollection.AddPathIssueMessages(
				errtype.CheckSumMismatch,
				otherFile,
				"hash not found",
				"Src file:",
				pathjoin.JoinNormalized(it.rootPath, otherFile),
				"Dest file:",
				pathjoin.JoinNormalized(other.rootPath, otherFile))
		}

		if !isContinueOnErr && stateTracker.HasChangesCollection() {
			return errCollection
		}
	}

	return errCollection
}

func (it *Instance) IsEqual(isTrimRoot bool, other *Instance) bool {
	if it == nil && other == nil {
		return true
	}

	if it == nil || other == nil {
		return false
	}

	// What should We do with errors?
	if it.HasError() || other.HasError() {
		return false
	}

	if len(it.rawHashes) != len(other.rawHashes) {
		return false
	}

	return it.isEqualAllCompare(isTrimRoot, other)
}

func (it *Instance) HasHash(filePath string) bool {
	_, isFound := it.rawHashes[filePath]

	return isFound
}

func (it *Instance) IsHashMatch(filePath string, fileHash []byte) bool {
	checksumBytes, isFound := it.rawHashes[filePath]

	if !isFound {
		return false
	}

	return bytes.Equal(checksumBytes, fileHash)
}

func (it *Instance) IsHashMatchString(filePath, fileHash string) bool {
	checksumBytes, isFound := it.rawHashes[filePath]

	if !isFound {
		return false
	}

	return hex.EncodeToString(checksumBytes) == fileHash
}

// StrippedPathMap Key => Root Stripped Path, Value => Actual Path
func (it *Instance) StrippedPathMap() map[string]string {
	if it == nil {
		return nil
	}

	if it.strippedPathWithActualPathMap != nil {
		return it.strippedPathWithActualPathMap
	}

	strippedPathToActualPath := make(
		map[string]string,
		it.Length())
	for actualPath := range it.rawHashes {
		trimmedPath := strings.TrimPrefix(
			actualPath,
			it.rootPath)
		strippedPathToActualPath[trimmedPath] = actualPath
	}

	it.strippedPathWithActualPathMap = strippedPathToActualPath

	return it.strippedPathWithActualPathMap
}

func (it *Instance) VerifyError(
	isContinueOnErr bool,
	isTrimRoot bool,
	other *Instance,
) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	if it == nil && other == nil {
		return errCollection
	}

	if it == nil || other == nil {
		return errCollection.AddUsingMessages(
			errtype.ValidationMismatch,
			"One of the instance is nil.")
	}

	// What should We do with errors?
	if it.HasError() || other.HasError() {
		return errCollection.AddTypeRefQuick(
			errtype.ValidationMismatch,
			"Either one of the instance has existing error.",
			"Source:",
			it.ErrorWrapper.String(),
			"other:",
			other.ErrorWrapper.String())
	}

	if len(it.rawHashes) != len(other.rawHashes) {
		return errCollection.AddTypeRefQuick(
			errtype.LengthMismatch,
			"Length mismatch.",
			"Source:",
			len(it.rawHashes),
			"other:",
			len(other.rawHashes))
	}

	return it.verifyEqualAllCompare(
		isTrimRoot,
		isContinueOnErr,
		errCollection,
		other)
}

func (it *Instance) IsFile() bool {
	return it.isFile
}

func (it *Instance) IsDir() bool {
	return !it.isFile
}

func (it *Instance) Length() int {
	if it == nil {
		return 0
	}

	return len(it.rawHashes)
}

func (it *Instance) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Instance) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}
