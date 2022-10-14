package pathchmod

import (
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type SimpleStatMap struct {
	Items map[string]*SimpleStat
}

func NewSimpleStatMap(capacity int) *SimpleStatMap {
	itemsMap := make(
		map[string]*SimpleStat,
		capacity)

	return &SimpleStatMap{
		itemsMap,
	}
}

func (it *SimpleStatMap) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *SimpleStatMap) Count() int {
	return it.Length()
}

func (it *SimpleStatMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it *SimpleStatMap) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *SimpleStatMap) LastIndex() int {
	return it.Length() - 1
}

func (it *SimpleStatMap) Add(loc string) *SimpleStatMap {
	stat := GetSimpleStat(loc)

	it.Items[loc] = stat

	return it
}

func (it *SimpleStatMap) Adds(locations ...string) *SimpleStatMap {
	if len(locations) == 0 {
		return it
	}

	for _, location := range locations {
		it.Add(location)
	}

	return it
}

func (it *SimpleStatMap) HasLocation(loc string) bool {
	_, has := it.Items[loc]

	return has
}

func (it *SimpleStatMap) Locations() []string {
	slice := make([]string, it.Length())

	index := 0
	for loc := range it.Items {
		slice[index] = loc
		index++
	}

	return slice
}

func (it *SimpleStatMap) Names() []string {
	slice := make([]string, it.Length())

	index := 0
	for _, loc := range it.Items {
		slice[index] = loc.Name
		index++
	}

	return slice
}

func (it *SimpleStatMap) SimpleStats() []*SimpleStat {
	slice := make([]*SimpleStat, it.Length())

	index := 0
	for _, simpleStat := range it.Items {
		slice[index] = simpleStat
		index++
	}

	return slice
}

func (it *SimpleStatMap) HasIndex(index int) bool {
	return it.LastIndex() <= index
}

func (it SimpleStatMap) String() string {
	return stringutil.AnyToStringNameField(it.Items)
}

func (it *SimpleStatMap) AsBasicMapper() coreinterface.BasicMapper {
	return it
}
