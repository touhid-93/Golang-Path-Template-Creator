package pathchmod

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

type SimpleStats struct {
	Items []*SimpleStat
}

func NewSimpleStats(capacity int) *SimpleStats {
	items := make(
		[]*SimpleStat,
		0,
		capacity)

	return &SimpleStats{
		items,
	}
}

func NewSimpleStatsUsingItems(locations ...string) *SimpleStats {
	capacity := len(locations) + constants.Capacity5

	return NewSimpleStats(capacity).
		Adds(locations...)
}

func (it *SimpleStats) Length() int {
	if it == nil || it.Items == nil {
		return 0
	}

	return len(it.Items)
}

func (it *SimpleStats) Count() int {
	return it.Length()
}

func (it *SimpleStats) IsEmpty() bool {
	return it.Length() == 0
}

func (it *SimpleStats) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *SimpleStats) LastIndex() int {
	return it.Length() - 1
}

func (it *SimpleStats) Add(loc string) *SimpleStats {
	stat := GetSimpleStat(loc)

	it.Items = append(it.Items, stat)

	return it
}

func (it *SimpleStats) Adds(locations ...string) *SimpleStats {
	if len(locations) == 0 {
		return it
	}

	for _, location := range locations {
		it.Add(location)
	}

	return it
}

func (it *SimpleStats) HasLocation(loc string) bool {
	for _, item := range it.Items {
		if item.Location == loc {
			return true
		}
	}

	return false
}

func (it *SimpleStats) Locations() []string {
	slice := make([]string, it.Length())

	index := 0
	for _, loc := range it.Items {
		slice[index] = loc.Location
		index++
	}

	return slice
}

func (it *SimpleStats) Names() []string {
	slice := make([]string, it.Length())

	index := 0
	for _, loc := range it.Items {
		slice[index] = loc.Name
		index++
	}

	return slice
}

func (it *SimpleStats) SimpleStats() []*SimpleStat {
	slice := make([]*SimpleStat, it.Length())

	index := 0
	for _, simpleStat := range it.Items {
		slice[index] = simpleStat
		index++
	}

	return slice
}

func (it *SimpleStats) HasIndex(index int) bool {
	return it.LastIndex() <= index
}

func (it SimpleStats) String() string {
	return stringutil.AnyToStringNameField(it.Items)
}

func (it *SimpleStats) AsBasicMapper() coreinterface.BasicMapper {
	return it
}
