package nginxlinuxpath

import (
	"log"
	"os"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
)

type NginxDirMap struct {
	currentNginxRoot string
	items            map[string]*NginxDir
	sync.Mutex
}

func NewNginxDirMap(configRoot string) *NginxDirMap {
	return NewNginxDirMapCap(
		configRoot,
		constants.ArbitraryCapacity50)
}

func NewNginxDirMapCap(configRoot string, capacity int) *NginxDirMap {
	return &NginxDirMap{
		currentNginxRoot: configRoot,
		items:            make(map[string]*NginxDir, capacity),
	}
}

func (it *NginxDirMap) Length() int {
	if it == nil {
		return constants.Zero
	}

	return len(it.items)
}

func (it *NginxDirMap) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.Length()
}

func (it *NginxDirMap) Count() int {
	return it.Length()
}

func (it *NginxDirMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it *NginxDirMap) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.Length() == 0
}

func (it *NginxDirMap) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *NginxDirMap) HasAnyItemLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.Length() > 0
}

func (it *NginxDirMap) LastIndex() int {
	return it.Length() - 1
}

func (it *NginxDirMap) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *NginxDirMap) Keys() []string {
	keys, err := coredynamic.MapKeysStringSliceAny(it)

	if err != nil {
		log.Fatalln(err)
	}

	return keys
}

func (it *NginxDirMap) KeysSorted() []string {
	return coredynamic.MapKeysStringSliceAnySortedMust(it)
}

func (it *NginxDirMap) AddOrUpdate(key string, nginxDir *NginxDir) *NginxDir {
	it.items[key] = nginxDir

	return nginxDir
}

func (it *NginxDirMap) AddOrUpdateLock(key string, nginxDir *NginxDir) *NginxDir {
	it.Lock()
	defer it.Unlock()

	it.items[key] = nginxDir

	return nginxDir
}

func (it *NginxDirMap) ClearAll() {
	if it == nil {
		return
	}

	it.items = map[string]*NginxDir{}
}

func (it *NginxDirMap) ClearAllLock() {
	if it == nil {
		return
	}

	it.Lock()
	defer it.Unlock()

	it.items = map[string]*NginxDir{}
}

func (it *NginxDirMap) DeleteKey(key string) {
	if it == nil {
		return
	}

	_, has := it.items[key]

	if has {
		delete(it.items, key)
	}
}

func (it *NginxDirMap) DeleteKeyLock(key string) {
	if it == nil {
		return
	}

	it.Lock()
	defer it.Unlock()

	_, has := it.items[key]

	if has {
		delete(it.items, key)
	}
}

func (it *NginxDirMap) HasKey(key string) bool {
	_, has := it.items[key]

	return has
}

func (it *NginxDirMap) IsMissingKey(key string) bool {
	_, has := it.items[key]

	return !has
}

func (it *NginxDirMap) GetSet(
	isNormalize bool,
	dirChmod os.FileMode,
	userName string,
) *NginxDir {
	if it == nil {
		return NewNginxDir(
			isNormalize,
			dirChmod,
			it.currentNginxRoot,
			userName)
	}

	nginxDir, has := it.items[userName]

	if has {
		return nginxDir
	}

	newNginxDir := NewNginxDir(
		isNormalize,
		dirChmod,
		it.currentNginxRoot,
		userName)

	return it.AddOrUpdate(
		userName,
		newNginxDir)
}

func (it *NginxDirMap) GetSetLock(
	isNormalize bool,
	dirChmod os.FileMode,
	userName string,
) *NginxDir {
	it.Lock()
	defer it.Unlock()

	return it.GetSet(
		isNormalize,
		dirChmod,
		userName)
}

func (it *NginxDirMap) GetSetDefault(
	isNormalize bool,
	userName string,
) *NginxDir {
	return it.GetSet(
		isNormalize,
		DefaultDirChmod,
		userName)
}

func (it *NginxDirMap) GetSetDefaultLock(
	isNormalize bool,
	userName string,
) *NginxDir {
	it.Lock()
	defer it.Unlock()

	return it.GetSet(
		isNormalize,
		DefaultDirChmod,
		userName)
}
