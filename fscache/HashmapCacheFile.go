package fscache

import (
	"sort"

	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

type HashmapCacheFile struct {
	CacheFile
	mapStringAny map[string]string
}

func (it *HashmapCacheFile) Read() (map[string]string, *errorwrapper.Wrapper) {
	if it == nil {
		return map[string]string{}, errnew.
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
	var newHashset map[string]string

	errWrap := it.GetOnce(&newHashset)
	it.mapStringAny = newHashset

	return newHashset, errWrap
}

func (it *HashmapCacheFile) IsFileIntegrityAlright() bool {
	it.setDefaultGeneratorOnInvalidFunc()

	var hashset map[string]string

	return it.IsCacheIntegrityAlright(&hashset)
}

func (it *HashmapCacheFile) CoreHashmap() (
	*corestr.Hashmap,
	*errorwrapper.Wrapper,
) {
	currentMappedItems, errWrap := it.Read()

	return corestr.New.Hashmap.UsingMap(currentMappedItems), errWrap
}

func (it *HashmapCacheFile) EnumDynamicMap() (
	enumimpl.DynamicMap,
	*errorwrapper.Wrapper,
) {
	currentMap, errWrap := it.Read()
	newMap := make(map[string]interface{}, len(currentMap))

	for key, val := range currentMap {
		newMap[key] = val
	}

	return newMap, errWrap
}

func (it *HashmapCacheFile) SafeHashmap() map[string]string {
	currentMap, _ := it.Read()

	return currentMap
}

func (it *HashmapCacheFile) SafeMayAnyItems() *coredynamic.MapAnyItems {
	currentMap, _ := it.Read()
	newMap := coredynamic.NewMapAnyItems(len(currentMap))

	for key, val := range currentMap {
		newMap.Add(key, val)
	}

	return newMap
}

func (it *HashmapCacheFile) SafeDynamicMap() enumimpl.DynamicMap {
	currentMap, _ := it.EnumDynamicMap()

	return currentMap
}

func (it *HashmapCacheFile) Strings() []string {
	dynamicMap := it.SafeDynamicMap()

	return dynamicMap.Strings()
}

func (it *HashmapCacheFile) String() string {
	dynamicMap := it.SafeDynamicMap()

	return dynamicMap.String()
}

func (it *HashmapCacheFile) Length() int {
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

func (it *HashmapCacheFile) IsEmpty() bool {
	return it.Length() == 0
}

func (it *HashmapCacheFile) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *HashmapCacheFile) GetValue(key string) (val string) {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return currentMap[key]
}

func (it *HashmapCacheFile) GetValueWithStat(key string) (val string, isFound bool) {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	val, isFound = currentMap[key]

	return val, isFound
}

func (it *HashmapCacheFile) HasKey(key string) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := currentMap[key]

	return has
}

func (it *HashmapCacheFile) HasAllKey(
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

func (it *HashmapCacheFile) IsMissing(key string) bool {
	currentMap, _ := it.Read()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := currentMap[key]

	return !has
}

func (it *HashmapCacheFile) IsAnyMissing(keys ...string) bool {
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

func (it *HashmapCacheFile) AddOrUpdateSave(
	key string,
	val string,
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

func (it *HashmapCacheFile) Save(
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

func (it *HashmapCacheFile) saveMapInternal(
	currentMap map[string]string,
) *errorwrapper.Wrapper {
	errWrap := it.saveInternal(currentMap)
	it.invalidateHashsetCache()

	return errWrap
}

func (it *HashmapCacheFile) AddOrUpdateMapAnyItemsSave(
	savingMap *coredynamic.MapAnyItems,
) (
	isAnyNewlyAdded bool,
	currentMap *coredynamic.MapAnyItems,
	savingErrWrap *errorwrapper.Wrapper,
) {
	isAnyNewlyAdded, savingErrWrap = it.AddOrUpdateEnumDynamicMapSave(
		savingMap.Items)

	if savingErrWrap.HasError() {
		return isAnyNewlyAdded, nil, savingErrWrap
	}

	return isAnyNewlyAdded,
		it.SafeMayAnyItems(),
		savingErrWrap
}

func (it *HashmapCacheFile) AddOrUpdateEnumDynamicMapSave(
	enumDynamicMap enumimpl.DynamicMap,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(enumDynamicMap) == 0 {
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

	for key, val := range enumDynamicMap {
		valString := converters.AnyToValueString(val)
		currentVal, has := currentMap[key]

		if has && currentVal == valString {
			continue
		}

		isAnyNewlyAdded = true
		currentMap[key] = valString
	}

	if isAnyNewlyAdded {
		return true, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) AddOrUpdateMapSave(
	mapStringAnyItem map[string]string,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(mapStringAnyItem) == 0 {
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

	for key, val := range mapStringAnyItem {
		currentVal, has := currentMap[key]

		if has && currentVal == val {
			continue
		}

		isAnyNewlyAdded = true
		currentMap[key] = val
	}

	if isAnyNewlyAdded {
		return true, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) AddOrUpdateKeyValuesSave(
	keyValues ...corestr.KeyAnyValuePair,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(keyValues) == 0 {
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

	for _, keyAnyVal := range keyValues {
		currentVal, has := currentMap[keyAnyVal.Key]

		if has && currentVal == keyAnyVal.ValueString() {
			continue
		}

		isAnyNewlyAdded = true
		currentMap[keyAnyVal.Key] = keyAnyVal.ValueString()
	}

	if isAnyNewlyAdded {
		return true, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) AddOrUpdateReferencesSave(
	references ...ref.Value,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(references) == 0 {
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

	for _, referenceKeyVal := range references {
		key := referenceKeyVal.KeyName()
		searchingVal := referenceKeyVal.ValueString()
		currentVal, has := currentMap[key]

		if has && currentVal == searchingVal {
			continue
		}

		isAnyNewlyAdded = true
		currentMap[key] = searchingVal
	}

	if isAnyNewlyAdded {
		return true, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) AddOrUpdateReferencesFullStringSave(
	references ...ref.Value,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(references) == 0 {
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

	for _, referenceKeyVal := range references {
		key := referenceKeyVal.KeyName()
		searchingVal := referenceKeyVal.ValueString()
		currentVal, has := currentMap[key]

		if has && currentVal == searchingVal {
			continue
		}

		isAnyNewlyAdded = true
		currentMap[key] = searchingVal
	}

	if isAnyNewlyAdded {
		return true, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) AddOrUpdateMetaCollectionSave(
	metaAttrCompiler loggerinf.MetaAttributesCompiler,
) (
	isAnyNewlyAdded bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if metaAttrCompiler == nil {
		return false, nil
	}

	return it.AddOrUpdateEnumDynamicMapSave(metaAttrCompiler.CompileMap())
}

func (it *HashmapCacheFile) AllKeys() (
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

func (it *HashmapCacheFile) AllKeysSorted() (
	allKeysSorted []string,
	readErrWrap *errorwrapper.Wrapper,
) {
	allKeys, readErrWrap := it.AllKeys()
	sort.Strings(allKeys)

	return allKeys, readErrWrap
}

func (it *HashmapCacheFile) AddOrUpdateManySave(
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

		currentMap[key] = "true"
		isAnyAddedNewly = true
	}

	if isAnyAddedNewly {
		return isAnyAddedNewly, it.saveMapInternal(currentMap)
	}

	return false, nil
}

func (it *HashmapCacheFile) invalidateHashsetCache() {
	if it == nil {
		return
	}

	it.mapStringAny = nil
}

func (it *HashmapCacheFile) setDefaultGeneratorOnInvalidFunc() {
	if it.IsOnInvalidGeneratorDefined() {
		return
	}

	it.GetSetInvalidGeneratorOnEmpty(hashmapOnInvalidDefaultFunc)
}
