package fscache

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type HashsetCacheFile struct {
	CacheFile
	hashset map[string]bool
}

func (it *HashsetCacheFile) Hashset() (map[string]bool, *errorwrapper.Wrapper) {
	if it == nil {
		return map[string]bool{}, errnew.
			Null.
			Simple(it)
	}

	if it.hashset != nil && it.IsCompiledSafe() {
		return it.hashset, nil
	}

	if it.IsCompiled() {
		// has error
		return it.hashset, it.compileErr
	}

	it.setDefaultGeneratorOnInvalidFunc()

	// generate
	var newHashset map[string]bool

	errWrap := it.GetOnce(&newHashset)
	it.hashset = newHashset

	return newHashset, errWrap
}

func (it *HashsetCacheFile) Read() (map[string]bool, *errorwrapper.Wrapper) {
	return it.Hashset()
}

func (it *HashsetCacheFile) IsHashsetFileIntegrityAlright() bool {
	it.setDefaultGeneratorOnInvalidFunc()

	var hashset map[string]bool

	return it.IsCacheIntegrityAlright(&hashset)
}

func (it *HashsetCacheFile) CoreHashset() (
	*corestr.Hashset,
	*errorwrapper.Wrapper,
) {
	hashset, errWrap := it.Hashset()

	return corestr.New.Hashset.UsingMap(hashset), errWrap
}

func (it *HashsetCacheFile) SafeHashset() map[string]bool {
	hashset, _ := it.Hashset()

	return hashset
}

func (it *HashsetCacheFile) Length() int {
	if it == nil {
		return 0
	}

	hashset, errWrap := it.Hashset()

	if errWrap.HasAnyError() {
		return 0
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	return len(hashset)
}

func (it *HashsetCacheFile) IsEmpty() bool {
	return it.Length() == 0
}

func (it *HashsetCacheFile) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *HashsetCacheFile) HasKey(key string) bool {
	hashset, _ := it.Hashset()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := hashset[key]

	return has
}

func (it *HashsetCacheFile) IsMissing(key string) bool {
	hashset, _ := it.Hashset()

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := hashset[key]

	return !has
}

func (it *HashsetCacheFile) AddOrUpdateSave(
	key string,
) (isAddedNewly bool, savingErrWrap *errorwrapper.Wrapper) {
	hashset, existingErrWrap := it.Hashset()

	if existingErrWrap.HasError() {
		return false, existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	_, has := hashset[key]

	if has {
		return false, nil
	}

	hashset[key] = true

	return true, it.saveHashsetInternal(hashset)
}

func (it *HashsetCacheFile) Save(
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

func (it *HashsetCacheFile) saveHashsetInternal(
	hashset map[string]bool,
) *errorwrapper.Wrapper {
	errWrap := it.saveInternal(hashset)
	it.invalidateHashsetCache()

	return errWrap
}

func (it *HashsetCacheFile) AddOrUpdateCoreHashsetSave(
	savingCoreHashset *corestr.Hashset,
) (
	isAnyAddedNewly bool,
	coreHashset *corestr.Hashset,
	savingErrWrap *errorwrapper.Wrapper,
) {
	isAnyAddedNewly, savingErrWrap = it.AddOrUpdateHashsetSave(
		savingCoreHashset.Items())

	currentHashset, errWrap := it.Hashset()

	return isAnyAddedNewly,
		corestr.New.Hashset.UsingMap(currentHashset),
		errnew.Merge.New(savingErrWrap, errWrap)
}

func (it *HashsetCacheFile) AddOrUpdateHashsetSave(
	anotherHashset map[string]bool,
) (
	isAnyAddedNewly bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(anotherHashset) == 0 {
		return false, nil
	}

	hashset, existingErrWrap := it.Hashset()

	if existingErrWrap.HasError() {
		return false, existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for key, isResult := range anotherHashset {
		currentVal, has := hashset[key]

		if has && currentVal == isResult {
			continue
		}

		hashset[key] = isResult
		isAnyAddedNewly = true
	}

	if isAnyAddedNewly {
		return isAnyAddedNewly, it.saveHashsetInternal(hashset)
	}

	return false, nil
}

func (it *HashsetCacheFile) AddOrUpdateManySave(
	keys ...string,
) (
	isAnyAddedNewly bool,
	savingErrWrap *errorwrapper.Wrapper,
) {
	if len(keys) == 0 {
		return false, nil
	}

	hashset, existingErrWrap := it.Hashset()

	if existingErrWrap.HasError() {
		return false, existingErrWrap
	}

	if it.isLockAcquire() {
		globalMutex.Lock()
		defer globalMutex.Unlock()
	}

	for _, key := range keys {
		_, has := hashset[key]

		if has {
			continue
		}

		hashset[key] = true
		isAnyAddedNewly = true
	}

	if isAnyAddedNewly {
		return isAnyAddedNewly, it.saveHashsetInternal(hashset)
	}

	return false, nil
}

func (it *HashsetCacheFile) invalidateHashsetCache() {
	if it == nil {
		return
	}

	it.hashset = nil
}

func (it *HashsetCacheFile) setDefaultGeneratorOnInvalidFunc() {
	if it.IsOnInvalidGeneratorDefined() {
		return
	}

	it.GetSetInvalidGeneratorOnEmpty(hashsetOnInvalidDefaultFunc)
}
