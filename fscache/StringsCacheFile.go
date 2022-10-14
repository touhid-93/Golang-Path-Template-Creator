package fscache

import (
	"sort"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type StringsCacheFile struct {
	CacheFile
	items []string
}

func (it *StringsCacheFile) CoreStringSlice() (*corestr.SimpleSlice, *errorwrapper.Wrapper) {
	return it.ReadSlice()
}

func (it *StringsCacheFile) ReadSlice() (*corestr.SimpleSlice, *errorwrapper.Wrapper) {
	if it == nil {
		return corestr.Empty.SimpleSlice(), errnew.Null.Simple(it)
	}

	slice, errWrap := it.Read()

	if errWrap.HasAnyError() {
		return corestr.Empty.SimpleSlice(), errWrap
	}

	return corestr.New.SimpleSlice.Strings(slice), nil
}

func (it *StringsCacheFile) Read() ([]string, *errorwrapper.Wrapper) {
	if it == nil {
		return []string{}, errnew.
			Null.
			Simple(it)
	}

	if it.IsCompiled() {
		// has error
		return it.items, it.compileErr
	}

	it.setDefaultGeneratorOnInvalidFunc()

	// generate
	var items []string
	errWrap := it.GetOnce(&items)
	it.items = items

	return items, errWrap
}

func (it *StringsCacheFile) IsFileIntegrityAlright() bool {
	it.setDefaultGeneratorOnInvalidFunc()

	var items []string

	return it.IsCacheIntegrityAlright(&items)
}

func (it *StringsCacheFile) CoreHashset() (
	*corestr.Hashset,
	*errorwrapper.Wrapper,
) {
	currentItems, errWrap := it.Read()

	return corestr.
		New.
		Hashset.
		StringsPtr(&currentItems), errWrap
}

func (it *StringsCacheFile) EnumDynamicMap() (
	enumimpl.DynamicMap,
	*errorwrapper.Wrapper,
) {
	currentMap, errWrap := it.Read()
	newMap := make(map[string]interface{}, len(currentMap))

	for index, name := range currentMap {
		newMap[name] = index
	}

	return newMap, errWrap
}

func (it *StringsCacheFile) SafeStrings() []string {
	currentItems, _ := it.Read()

	return currentItems
}

func (it *StringsCacheFile) SafeHashmap() map[string]string {
	currentItems, _ := it.Read()

	newMap := make(map[string]string, len(currentItems))

	for _, item := range currentItems {
		newMap[item] = item
	}

	return newMap
}

func (it *StringsCacheFile) SafeHashset() map[string]bool {
	currentItems, _ := it.Read()

	newMap := make(map[string]bool, len(currentItems))

	for _, item := range currentItems {
		newMap[item] = true
	}

	return newMap
}

func (it *StringsCacheFile) SafeMayAnyItems() *coredynamic.MapAnyItems {
	currentItems, _ := it.Read()
	newMap := coredynamic.NewMapAnyItems(len(currentItems))

	for index, name := range currentItems {
		newMap.Add(name, index)
	}

	return newMap
}

func (it *StringsCacheFile) SafeDynamicMap() enumimpl.DynamicMap {
	currentMap, _ := it.EnumDynamicMap()

	return currentMap
}

func (it *StringsCacheFile) Strings() []string {
	currentItems, _ := it.Read()

	return currentItems
}

func (it *StringsCacheFile) String() string {
	dynamicMap := it.SafeDynamicMap()

	return dynamicMap.String()
}

func (it *StringsCacheFile) Length() int {
	if it == nil {
		return 0
	}

	hashset, errWrap := it.Read()

	if errWrap.HasAnyError() {
		return 0
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return len(hashset)
}

func (it *StringsCacheFile) IsEmpty() bool {
	return it.Length() == 0
}

func (it *StringsCacheFile) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *StringsCacheFile) GetValueByIndex(index int) (val string) {
	currentItems, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return currentItems[index]
}

func (it *StringsCacheFile) GetSafeValueByIndex(index int) (val string) {
	currentItems, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return stringslice.SafeIndexAt(
		currentItems, index)
}

func (it *StringsCacheFile) AppendLinesSave(
	items ...string,
) (savingErrWrap *errorwrapper.Wrapper) {
	currentItems, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	currentItems = append(currentItems, items...)

	return it.saveStringsInternal(currentItems)
}

func (it *StringsCacheFile) Save(
	stringItems ...string,
) *errorwrapper.Wrapper {
	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.saveStringsInternal(stringItems)
}

func (it *StringsCacheFile) AppendSave(
	stringItems ...string,
) *errorwrapper.Wrapper {
	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return it.saveStringsInternal(stringItems)
}

func (it *StringsCacheFile) saveStringsInternal(
	currentItems []string,
) *errorwrapper.Wrapper {
	errWrap := it.saveInternal(currentItems)
	it.invalidateHashsetCache()

	return errWrap
}

func (it *StringsCacheFile) AddOrUpdateMetaCollectionSave(
	metaAttrCompiler loggerinf.MetaAttributesCompiler,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if metaAttrCompiler == nil {
		return nil
	}

	return it.Save(metaAttrCompiler.CompileStacks()...)
}

func (it *StringsCacheFile) StringsSorted() (
	allKeysSorted []string,
	readErrWrap *errorwrapper.Wrapper,
) {
	allKeys, readErrWrap := it.Read()
	sort.Strings(allKeys)

	return allKeys, readErrWrap
}

func (it *StringsCacheFile) SafeStringsSorted() (allKeysSorted []string) {
	allKeys, _ := it.Read()
	sort.Strings(allKeys)

	return allKeys
}

func (it *StringsCacheFile) invalidateHashsetCache() {
	if it == nil {
		return
	}

	it.items = nil
}

func (it *StringsCacheFile) setDefaultGeneratorOnInvalidFunc() {
	if it.IsOnInvalidGeneratorDefined() {
		return
	}

	it.GetSetInvalidGeneratorOnEmpty(stringsOnInvalidDefaultFunc)
}
