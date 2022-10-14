package unipaths

import (
	"errors"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

type Wrappers struct {
	separator   string
	isFinalized bool
	sync.Mutex
	items []*unipath.Wrapper
}

func (it *Wrappers) IsFinalized() bool {
	return it.isFinalized
}

func (it *Wrappers) Lock() {
	it.Lock()
}

func (it *Wrappers) Unlock() {
	it.Unlock()
}

func (it *Wrappers) AddPathAs(
	givenPath string,
) *Wrappers {
	wrapper := unipath.NewUsingPath(
		givenPath,
		it.separator)

	it.items = append(
		it.items,
		wrapper)

	return it
}

func (it *Wrappers) AddPathsLock(
	givenPaths ...string,
) *Wrappers {
	it.Lock()
	defer it.Unlock()

	if givenPaths == nil {
		return it
	}

	return it.
		AddPathsPtr(&givenPaths)
}

func (it *Wrappers) AddPaths(
	givenPaths ...string,
) *Wrappers {
	if givenPaths == nil {
		return it
	}

	return it.
		AddPathsPtr(&givenPaths)
}

func (it *Wrappers) AddPathsPtr(
	givenPaths *[]string,
) *Wrappers {
	if givenPaths == nil {
		return it
	}

	for _, currentPath := range *givenPaths {
		wrapper := unipath.NewUsingPath(
			currentPath,
			it.separator)

		it.items = append(
			it.items,
			wrapper,
		)
	}

	return it
}

func (it *Wrappers) AddWrapper(
	wrapper *unipath.Wrapper,
) *Wrappers {
	if wrapper == nil {
		return it
	}

	it.items = append(
		it.items,
		wrapper)

	return it
}

func (it *Wrappers) AddWrapperLock(
	wrapper *unipath.Wrapper,
) *Wrappers {
	it.Lock()
	defer it.Unlock()

	if wrapper == nil {
		return it
	}

	it.items = append(
		it.items,
		wrapper)

	return it
}

func (it *Wrappers) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *Wrappers) HasItems() bool {
	return it.Length() > 0
}

func (it *Wrappers) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Wrappers) IsEqual(wrappers *Wrappers) bool {
	if wrappers == nil && it == nil {
		return true
	}

	if wrappers == nil || it == nil {
		return false
	}

	if wrappers == it {
		return true
	}

	if wrappers.isFinalized != it.isFinalized {
		return false
	}

	if wrappers.Length() != it.Length() {
		return false
	}

	if &it.items == &wrappers.items {
		return true
	}

	for index, receiverWrapper := range it.items {
		anotherWrapper := wrappers.items[index]

		if anotherWrapper == nil && receiverWrapper == nil {
			continue
		}

		if anotherWrapper == nil || receiverWrapper == nil {
			return false
		}

		if !anotherWrapper.IsEqual(receiverWrapper) {
			return false
		}
	}

	return true
}

func (it *Wrappers) FinalizeAll() {
	if it.isFinalized {
		return
	}

	// set finalize error
	for _, wrapper := range it.items {
		wrapper.Finalize()
	}

	it.isFinalized = true
}

func (it *Wrappers) GetFinalizePaths() *errstr.ResultsWithErrorCollection {
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

	for i, wrapper := range it.items {
		finalizedResult := wrapper.GetFinalizePath()
		errCollection.AddWrapperPtr(finalizedResult.ErrorWrapper)
		list[i] = finalizedResult.Value
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        list,
		ErrorWrappers: errCollection,
	}
}

func (it *Wrappers) Items() []*unipath.Wrapper {
	return it.items
}

func (it *Wrappers) ListPtr() []*unipath.Wrapper {
	list := make([]*unipath.Wrapper, it.Length())

	i := 0
	for _, wrapper := range it.items {
		list[i] = wrapper
		i++
	}

	return list
}

func (it *Wrappers) ToStringsPtr(
	separator string,
	isNormalize bool,
) *[]string {
	list := make([]string, it.Length())

	i := 0
	for _, wrapper := range it.items {
		list[i] = wrapper.ToString(
			separator,
			isNormalize)
		i++
	}

	return &list
}

func (it *Wrappers) Strings() []string {
	list := make([]string, it.Length())

	i := constants.Zero
	for _, wrapper := range it.items {
		list[i] = wrapper.String()

		i++
	}

	return list
}

func (it *Wrappers) StringsCollectionPtr() *corestr.Collection {
	return corestr.New.Collection.Strings(
		it.Strings(),
	)
}

func (it *Wrappers) GetAt(
	index int,
) *unipath.Wrapper {
	return it.items[index]
}

func (it *Wrappers) GetSafeAt(
	index int,
) *unipath.Wrapper {
	if index > constants.InvalidNotFoundCase && index <= it.Length()-1 {
		return it.items[index]
	}

	return nil
}
