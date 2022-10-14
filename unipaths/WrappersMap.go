package unipaths

import (
	"errors"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

type WrappersMap struct {
	separator   string
	isFinalized bool
	sync.Mutex
	items map[string]*unipath.Wrapper
}

func (it *WrappersMap) IsFinalized() bool {
	return it.isFinalized
}

func (it *WrappersMap) Lock() {
	it.Lock()
}

func (it *WrappersMap) Unlock() {
	it.Unlock()
}

func (it *WrappersMap) AddLock(
	key,
	anyPath string,
) *WrappersMap {
	it.Lock()
	defer it.Unlock()

	return it.addLock(key, anyPath)
}

func (it *WrappersMap) addLock(
	key,
	anyPath string,
) *WrappersMap {
	wrapper, has := it.items[key]

	if !has {
		wrapper = unipath.New(it.separator)
		it.items[key] = wrapper
	}

	wrapper.AddLock(anyPath)

	return it
}

func (it *WrappersMap) Add(
	key,
	anyPath string,
) *WrappersMap {
	wrapper, has := it.items[key]

	if !has {
		wrapper = unipath.New(it.separator)
		it.items[key] = wrapper
	}

	wrapper.Add(anyPath)

	return it
}

func (it *WrappersMap) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *WrappersMap) HasItems() bool {
	return it.Length() > constants.Zero
}

// Has same one needs to be inserted
func (it *WrappersMap) Has(
	key,
	pathSplit string,
) bool {
	wrapper, has := it.items[key]

	return has && wrapper.Has(pathSplit)
}

func (it *WrappersMap) IsEmpty() bool {
	return it.Length() == constants.Zero
}

func (it *WrappersMap) IsEqual(wrappersMap *WrappersMap) bool {
	if wrappersMap == nil && it == nil {
		return true
	}

	if wrappersMap == nil || it == nil {
		return false
	}

	if wrappersMap == it {
		return true
	}

	if wrappersMap.isFinalized != it.isFinalized {
		return false
	}

	if wrappersMap.Length() != it.Length() {
		return false
	}

	if &it.items == &wrappersMap.items {
		return true
	}

	for key, receiverWrapper := range it.items {
		anotherWrapper, has := wrappersMap.items[key]

		if !has {
			return false
		}

		if !anotherWrapper.IsEqual(receiverWrapper) {
			return false
		}
	}

	return true
}

// FinalizeAll all wrappers
func (it *WrappersMap) FinalizeAll() {
	if it.isFinalized {
		return
	}

	// set finalize error
	it.isFinalized = true
	for _, wrapper := range it.items {
		wrapper.Finalize()
	}
}

func (it *WrappersMap) GetFinalizePath(
	key string,
) *errstr.Result {
	wrapper, has := it.items[key]

	if has {
		return wrapper.GetFinalizePath()
	}

	return errstr.New.Result.ErrorWrapper(
		errnew.Ref.OnlyOne(
			errtype.NotContainsExpectation,
			"Key",
			key))
}

func (it *WrappersMap) GetFinalizePaths() *errstr.ResultsWithErrorCollection {
	if !it.IsFinalized() {
		return errstr.New.ResultsWithErrorCollection.Error(
			errtype.Unexpected,
			errors.New(nonFinalizePathsCannotBeRetrievedMessage))
	}

	length := it.Length()
	list := make(
		[]string,
		length,
		length)

	errCollection := errwrappers.Empty()

	i := constants.Zero
	for _, wrapper := range it.items {
		finalizedResult := wrapper.GetFinalizePath()
		errCollection.AddWrapperPtr(finalizedResult.ErrorWrapper)
		list[i] = finalizedResult.Value

		i++
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        list,
		ErrorWrappers: errCollection,
	}
}

func (it *WrappersMap) Items() map[string]*unipath.Wrapper {
	return it.items
}

func (it *WrappersMap) ListPtr() []*unipath.Wrapper {
	list := make([]*unipath.Wrapper, it.Length())

	i := constants.Zero
	for _, wrapper := range it.items {
		list[i] = wrapper
		i++
	}

	return list
}

func (it *WrappersMap) ToStringsPtr(
	separator string,
	isNormalize bool,
) *[]string {
	list := make([]string, it.Length())

	i := constants.Zero
	for _, wrapper := range it.items {
		list[i] = wrapper.ToString(
			separator,
			isNormalize)
		i++
	}

	return &list
}

func (it *WrappersMap) StringsPtr() *[]string {
	list := make([]string, it.Length())

	i := constants.Zero
	for _, wrapper := range it.items {
		list[i] = wrapper.String()
		i++
	}

	return &list
}

func (it *WrappersMap) Get(
	key string,
) *unipath.Wrapper {
	return it.items[key]
}
