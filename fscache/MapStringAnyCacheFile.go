package fscache

import (
	"sort"

	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type MapStringAnyCacheFile struct {
	CacheFile
	mapStringAny map[string]interface{}
}

func (it *MapStringAnyCacheFile) Read() (map[string]interface{}, *errorwrapper.Wrapper) {
	if it == nil {
		return map[string]interface{}{}, errnew.
			Null.
			Simple(it)
	}

	if it.mapStringAny != nil && it.IsCompiledSafe() {
		return it.mapStringAny, nil
	}

	if it.IsCompiled() {
		// has error
		return it.mapStringAny, it.compileErr
	}

	it.setDefaultGeneratorOnInvalidFunc()

	// generate
	var newHashset map[string]interface{}

	errWrap := it.GetOnce(&newHashset)
	it.mapStringAny = newHashset

	return newHashset, errWrap
}

func (it *MapStringAnyCacheFile) IsFileIntegrityAlright() bool {
	it.setDefaultGeneratorOnInvalidFunc()

	var hashset map[string]interface{}

	return it.IsCacheIntegrityAlright(&hashset)
}

func (it *MapStringAnyCacheFile) CoreMapAnyItems() (
	*coredynamic.MapAnyItems,
	*errorwrapper.Wrapper,
) {
	mapStringAnyItems, errWrap := it.Read()

	return coredynamic.NewMapAnyItemsUsingItems(mapStringAnyItems), errWrap
}

func (it *MapStringAnyCacheFile) EnumDynamicMap() (
	enumimpl.DynamicMap,
	*errorwrapper.Wrapper,
) {
	mapStringAnyItems, errWrap := it.Read()

	return mapStringAnyItems, errWrap
}

func (it *MapStringAnyCacheFile) SafeMapStringAny() map[string]interface{} {
	currentMap, _ := it.Read()

	return currentMap
}

func (it *MapStringAnyCacheFile) SafeMayAnyItems() *coredynamic.MapAnyItems {
	currentMap, _ := it.Read()

	return coredynamic.NewMapAnyItemsUsingItems(currentMap)
}

func (it *MapStringAnyCacheFile) SafeDynamicMap() enumimpl.DynamicMap {
	currentMap, _ := it.Read()

	return currentMap
}

func (it *MapStringAnyCacheFile) Strings() []string {
	dynamicMap := it.SafeDynamicMap()

	return dynamicMap.Strings()
}

func (it *MapStringAnyCacheFile) String() string {
	dynamicMap := it.SafeDynamicMap()

	return dynamicMap.String()
}

func (it *MapStringAnyCacheFile) Length() int {
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

func (it *MapStringAnyCacheFile) IsEmpty() bool {
	return it.Length() == 0
}

func (it *MapStringAnyCacheFile) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *MapStringAnyCacheFile) GetValue(key string) (val interface{}) {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return currentMap[key]
}

func (it *MapStringAnyCacheFile) GetValueWithStat(key string) (val interface{}, isFound bool) {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	val, isFound = currentMap[key]

	return val, isFound
}

func (it *MapStringAnyCacheFile) HasKey(key string) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := currentMap[key]

	return has
}

func (it *MapStringAnyCacheFile) HasAllKey(
	keys ...string,
) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if len(keys) == 0 {
		return len(currentMap) == 0
	}

	for _, key := range keys {
		_, has := currentMap[key]

		if !has {
			return false
		}
	}

	return true
}

func (it *MapStringAnyCacheFile) IsMissing(key string) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := currentMap[key]

	return !has
}

func (it *MapStringAnyCacheFile) IsAnyMissing(keys ...string) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	if len(keys) == 0 {
		return len(currentMap) == 0
	}

	for _, key := range keys {
		_, has := currentMap[key]

		if !has {
			return true
		}
	}

	return false
}

func (it *MapStringAnyCacheFile) AddOrUpdateSave(
	key string,
	val interface{},
) (savingErrWrap *errorwrapper.Wrapper) {
	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	currentMap[key] = val

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) Save(
	hashset map[string]bool,
) *errorwrapper.Wrapper {
	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	errWrap := it.saveInternal(hashset)
	it.invalidateHashsetCache()

	return errWrap
}

func (it *MapStringAnyCacheFile) saveMapInternal(
	currentMap map[string]interface{},
) *errorwrapper.Wrapper {
	errWrap := it.saveInternal(currentMap)
	it.invalidateHashsetCache()

	return errWrap
}

func (it *MapStringAnyCacheFile) AddOrUpdateMapAnyItemsSave(
	savingMap *coredynamic.MapAnyItems,
) (
	currentMap *coredynamic.MapAnyItems,
	savingErrWrap *errorwrapper.Wrapper,
) {
	savingErrWrap = it.AddOrUpdateMapSave(savingMap.Items)

	rawCurrentMap, errWrap := it.Read()

	return coredynamic.NewMapAnyItemsUsingItems(rawCurrentMap),
		errnew.Merge.New(savingErrWrap, errWrap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateEnumDynamicMapSave(
	enumDynamicMap enumimpl.DynamicMap,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	return it.AddOrUpdateMapSave(enumDynamicMap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateMapSave(
	mapStringAnyItem map[string]interface{},
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(mapStringAnyItem) == 0 {
		return nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for key, val := range mapStringAnyItem {
		currentMap[key] = val
	}

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateKeyValuesSave(
	keyValues ...corestr.KeyAnyValuePair,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(keyValues) == 0 {
		return nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for _, keyAnyVal := range keyValues {
		currentMap[keyAnyVal.Key] = keyAnyVal.Value
	}

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateReferencesSave(
	references ...ref.Value,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(references) == 0 {
		return nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for _, referenceKeyVal := range references {
		currentMap[referenceKeyVal.KeyName()] = referenceKeyVal.Value
	}

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateReferencesFullStringSave(
	references ...ref.Value,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(references) == 0 {
		return nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for _, referenceKeyVal := range references {
		currentMap[referenceKeyVal.KeyName()] = referenceKeyVal.FullString()
	}

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) AddOrUpdateMetaCollectionSave(
	metaAttrCompiler loggerinf.MetaAttributesCompiler,
) (
	savingErrWrap *errorwrapper.Wrapper,
) {
	if metaAttrCompiler == nil {
		return nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for key, val := range metaAttrCompiler.CompileMap() {
		currentMap[key] = val
	}

	return it.saveMapInternal(currentMap)
}

func (it *MapStringAnyCacheFile) AllKeys() (
	allKeys []string,
	readErrWrap *errorwrapper.Wrapper,
) {
	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return allKeys, existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	allKeys = make([]string, len(currentMap))
	index := 0

	for key := range currentMap {
		allKeys[index] = key
		index++
	}

	return allKeys, nil
}

func (it *MapStringAnyCacheFile) AllKeysSorted() (
	allKeysSorted []string,
	readErrWrap *errorwrapper.Wrapper,
) {
	allKeys, readErrWrap := it.AllKeys()
	sort.Strings(allKeys)

	return allKeys, readErrWrap
}

func (it *MapStringAnyCacheFile) AddOrUpdateManySave(
	keys ...string,
) (
	isAnyAddedNewly bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(keys) == 0 {
		return false, nil
	}

	currentMap, existingErrWrap := it.Read()

	if existingErrWrap.HasError() {
		return false, existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for _, key := range keys {
		_, has := currentMap[key]

		if has {
			continue
		}

		currentMap[key] = true
		isAnyAddedNewly = true
	}

	if isAnyAddedNewly {
		return isAnyAddedNewly, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *MapStringAnyCacheFile) invalidateHashsetCache() {
	if it == nil {
		return
	}

	it.mapStringAny = nil
}

func (it *MapStringAnyCacheFile) setDefaultGeneratorOnInvalidFunc() {
	if it.IsOnInvalidGeneratorDefined() {
		return
	}

	it.GetSetInvalidGeneratorOnEmpty(mapStringAnyOnInvalidDefaultFunc)
}
