package filestate

import (
	"math"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

// MappedInfoItems
//
// Key path
type MappedInfoItems struct {
	Items map[string]*Info `json:"MapItems,omitempty"` // Key path
}

func NewMappedInfoItems(capacity int) *MappedInfoItems {
	slice := make(map[string]*Info, capacity)

	return &MappedInfoItems{
		slice,
	}
}

func EmptyMappedInfoItems() *MappedInfoItems {
	return NewMappedInfoItems(constants.Zero)
}

func NewMappedInfoItemsUsingInfoItems(
	infoItems ...*Info,
) *MappedInfoItems {
	if infoItems == nil {
		return EmptyMappedInfoItems()
	}

	newMapCollection := NewMappedInfoItems(len(infoItems))

	return newMapCollection.Adds(infoItems...)
}

func NewMappedInfoItemsUsingFilePaths(
	hashMethod hashas.Variant,
	isNormalize bool,
	filePaths ...string,
) (*MappedInfoItems, *errwrappers.Collection) {
	if len(filePaths) == 0 {
		return EmptyMappedInfoItems(), nil
	}

	infoCollection, errCollection := NewInfoCollectionUsingFilePathsAsync(
		hashMethod,
		isNormalize,
		filePaths...)

	newMapCollection := NewMappedInfoItems(infoCollection.Length())
	newMapCollection.Adds(infoCollection.Items...)

	return newMapCollection, errCollection
}

func (it *MappedInfoItems) CleanupNonNull() *MappedInfoItems {
	mapItems := make(map[string]*Info, it.Length())

	for key, item := range it.Items {
		if item == nil {
			continue
		}

		mapItems[key] = item
	}

	it.Items = mapItems

	return it
}

func (it *MappedInfoItems) Add(
	stateInfo *Info,
) *MappedInfoItems {
	it.Items[stateInfo.FullPath] = stateInfo

	return it
}

func (it *MappedInfoItems) Adds(
	infoItems ...*Info,
) *MappedInfoItems {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if stateInfo == nil {
			continue
		}

		it.Items[stateInfo.FullPath] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) AddMapItems(
	infoMapItems map[string]*Info,
) *MappedInfoItems {
	if len(infoMapItems) == 0 {
		return it
	}

	for key, stateInfo := range infoMapItems {
		it.Items[key] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) AddMapCollection(
	mapCollection *MappedInfoItems,
) *MappedInfoItems {
	if mapCollection.IsEmpty() {
		return it
	}

	for key, stateInfo := range mapCollection.Items {
		it.Items[key] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) AddsPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *MappedInfoItems {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid {
			continue
		}

		it.Items[stateInfo.FullPath] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) AddFilesPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *MappedInfoItems {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid || stateInfo.IsDir() {
			continue
		}

		it.Items[stateInfo.FullPath] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) AddDirsPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *MappedInfoItems {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid || stateInfo.IsFile {
			continue
		}

		it.Items[stateInfo.FullPath] = stateInfo
	}

	return it
}

func (it *MappedInfoItems) ConcatNew(
	additionalInfoItems ...*Info,
) *MappedInfoItems {
	cloned := it.Clone()

	return cloned.Adds(additionalInfoItems...)
}

func (it *MappedInfoItems) ConcatNewPtr(
	additionalInfoItems ...*Info,
) *MappedInfoItems {
	cloned := it.Clone()

	return cloned.AddsPtr(
		true,
		additionalInfoItems...)
}

func (it *MappedInfoItems) AddsIf(
	isAdd bool,
	infoItems ...*Info,
) *MappedInfoItems {
	if !isAdd {
		return it
	}

	return it.Adds(infoItems...)
}

func (it *MappedInfoItems) HasChecksum(hexChecksum string) bool {
	return it.HasFilterFuncAsync(func(key string, info *Info) (isSuccess bool) {
		return info.HexContentChecksum == hexChecksum
	})
}

func (it *MappedInfoItems) HasAnyChecksum(hexChecksums ...string) bool {
	findingChecksumsHashset := corestr.New.Hashset.StringsPtr(
		&hexChecksums)

	return it.HasFilterFuncAsync(func(key string, info *Info) (isSuccess bool) {
		return findingChecksumsHashset.Has(info.HexContentChecksum)
	})
}

func (it *MappedInfoItems) HasAnyFilePath(filePaths ...string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, filePath := range filePaths {
		_, has := it.Items[filePath]

		if has {
			return true
		}
	}

	return false
}

func (it *MappedInfoItems) HasAllFilePath(filePaths ...string) bool {
	if it.IsEmpty() {
		return false
	}

	for _, filePath := range filePaths {
		_, has := it.Items[filePath]

		if !has {
			return false
		}
	}

	return true
}

func (it *MappedInfoItems) InfoUsingFilePaths(filePaths ...string) *InfoCollection {
	if it.IsEmpty() {
		return EmptyInfoCollection()
	}

	slice := make([]*Info, 0, len(filePaths))

	for _, filePath := range filePaths {
		info, has := it.Items[filePath]

		if has {
			slice = append(slice, info)
		}
	}

	return &InfoCollection{Items: slice}
}

func (it *MappedInfoItems) InfoUsingChecksums(checksums ...string) *InfoCollection {
	if it.IsEmpty() {
		return EmptyInfoCollection()
	}

	slice := make([]*Info, 0, len(checksums))

	hexChecksumToInfoMap := it.AllHexChecksumToInfoMap()
	for _, currentChecksum := range checksums {
		info, has := hexChecksumToInfoMap[currentChecksum]

		if has {
			slice = append(slice, info)
		}
	}

	return &InfoCollection{Items: slice}
}

func (it *MappedInfoItems) GetInfoByFilePath(
	filePath string,
) *Info {
	info, has := it.Items[filePath]

	if has {
		return info
	}

	return nil
}

func (it *MappedInfoItems) AllFilteredInfoItemsByFilePaths(
	filePaths ...string,
) *InfoCollection {
	if len(filePaths) == 0 || it.IsEmpty() {
		return EmptyInfoCollection()
	}

	slice := make([]*Info, 0, len(filePaths))

	for _, filePath := range filePaths {
		item, has := it.Items[filePath]

		if has {
			slice = append(slice, item)
		}
	}

	return &InfoCollection{
		Items: slice,
	}
}

func (it *MappedInfoItems) AllKeys() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	index := 0
	for key := range it.Items {
		slice[index] = key
		index++
	}

	return slice
}

func (it *MappedInfoItems) AllKeysSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.AllKeys()
	sort.Strings(slice)

	return slice
}

func (it *MappedInfoItems) AllChecksums() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	index := 0
	for _, item := range it.Items {
		slice[index] = item.HexContentChecksum
		index++
	}

	return slice
}

func (it *MappedInfoItems) AllChecksumsSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.AllChecksums()
	sort.Strings(slice)

	return slice
}

func (it *MappedInfoItems) KeyValueStringMapUsingFmtFunc(
	fmtFunc MapKeyValFmtFunc,
) map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	itemsMap := make(map[string]string, it.Length())

	for _, item := range it.Items {
		k, v := fmtFunc(item)

		itemsMap[k] = v
	}

	return itemsMap
}

func (it *MappedInfoItems) PathStatSlice(isTakeOnlyValid bool) []*chmodhelper.PathExistStat {
	if it.IsEmpty() {
		return []*chmodhelper.PathExistStat{}
	}

	slice := make(
		[]*chmodhelper.PathExistStat,
		0,
		it.Length())

	if isTakeOnlyValid {
		for _, item := range it.Items {
			stat := item.Stat()

			if stat.IsInvalid() {
				continue
			}

			slice = append(slice, stat)
		}

		return slice
	}

	for _, item := range it.Items {
		stat := item.Stat()
		slice = append(slice, stat)
	}

	return slice
}

func (it *MappedInfoItems) KeyStringValueInfoMapUsingFmtFunc(
	fmtFunc MapKeyValInfoFmtFunc,
) map[string]*Info {
	if it.IsEmpty() {
		return map[string]*Info{}
	}

	itemsMap := make(map[string]*Info, it.Length())

	for _, item := range it.Items {
		k := fmtFunc(item)

		itemsMap[k] = item
	}

	return itemsMap
}

// AllFilePathToHexChecksumMap
//
// Key = filePath,
//
// Value = hex-checksum
func (it *MappedInfoItems) AllFilePathToHexChecksumMap() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	itemsMap := make(map[string]string, it.Length())

	for _, item := range it.Items {
		itemsMap[item.FullPath] = item.HexContentChecksum
	}

	return itemsMap
}

// AllFilePathToInfoMap
//
// Key = filePath,
//
// Value = Info
func (it *MappedInfoItems) AllFilePathToInfoMap() map[string]*Info {
	if it.IsEmpty() {
		return map[string]*Info{}
	}

	itemsMap := make(map[string]*Info, it.Length())

	for _, item := range it.Items {
		itemsMap[item.FullPath] = item
	}

	return itemsMap
}

// AllHexChecksumToInfoMap
//
// Key = HexChecksum,
//
// Value = Info
func (it *MappedInfoItems) AllHexChecksumToInfoMap() map[string]*Info {
	if it.IsEmpty() {
		return map[string]*Info{}
	}

	itemsMap := make(map[string]*Info, it.Length())

	for _, item := range it.Items {
		itemsMap[item.HexContentChecksum] = item
	}

	return itemsMap
}

// AllHexChecksumToFilePathMap
//
// Key = HexChecksum,
//
// Value = FilePath
func (it *MappedInfoItems) AllHexChecksumToFilePathMap() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	itemsMap := make(map[string]string, it.Length())

	for _, item := range it.Items {
		itemsMap[item.HexContentChecksum] = item.FullPath
	}

	return itemsMap
}

func (it *MappedInfoItems) CompiledChecksum(isSortChecksum bool) *errstr.Result {
	if it.IsEmpty() {
		return errstr.Empty.Result()
	}

	return hexchecksum.OfChecksums(
		true,
		isSortChecksum,
		DefaultHashMethod,
		it.AllChecksums()...)
}

func (it *MappedInfoItems) CompiledChecksumString(isSortChecksum bool) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	result := hexchecksum.OfChecksums(
		true,
		isSortChecksum,
		DefaultHashMethod,
		it.AllChecksums()...)

	go result.ErrorWrapper.Dispose()

	return result.Value
}

func (it *MappedInfoItems) MappedInfoItemsByKeys(keys ...string) *MappedInfoItems {
	if it.IsEmpty() {
		return EmptyMappedInfoItems()
	}

	newMappedInfoItems := NewMappedInfoItems(len(keys))
	for _, key := range keys {
		item, has := it.Items[key]

		if has {
			newMappedInfoItems.Adds(item)
		}
	}

	return newMappedInfoItems
}

func (it *MappedInfoItems) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *MappedInfoItems) GetPagedCollection(
	eachPageSize int,
) []*MappedInfoItems {
	length := it.Length()

	if length < eachPageSize {
		return []*MappedInfoItems{}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*MappedInfoItems, pagesPossibleCeiling)
	allKeys := it.AllKeysSorted()
	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
			allKeys)

		collectionOfCollection[oneBasedPageIndex-1] = pagedCollection

		wg.Done()
	}

	wg.Add(pagesPossibleCeiling)
	for i := 1; i <= pagesPossibleCeiling; i++ {
		go addPagedItemsFunc(i)
	}

	wg.Wait()

	return collectionOfCollection
}

// GetSinglePageCollection PageIndex is one based index. Should be above or equal 1
func (it *MappedInfoItems) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
	allKeys []string,
) *MappedInfoItems {
	length := it.Length()

	if length < eachPageSize {
		return it
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := eachPageSize * (pageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndexType.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				pageIndex)
	}

	endingIndex := skipItems + eachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	keys := allKeys[skipItems:endingIndex]

	return it.MappedInfoItemsByKeys(keys...)
}

func (it *MappedInfoItems) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *MappedInfoItems) Count() int {
	return it.Length()
}

func (it *MappedInfoItems) IsEmpty() bool {
	return it.Length() == 0
}

func (it *MappedInfoItems) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *MappedInfoItems) LastIndex() int {
	return it.Length() - 1
}

func (it *MappedInfoItems) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *MappedInfoItems) Strings() []string {
	list := make([]string, it.Length())

	index := 0
	for _, item := range it.Items {
		list[index] = item.String()
		index++
	}

	return list
}

func (it *MappedInfoItems) StringsUsingStringerFmtFunc(
	fmtFunc StringerFmtFunc,
) []string {
	list := make([]string, it.Length())

	index := 0
	for _, item := range it.Items {
		list[index] = fmtFunc(index, item)
		index++
	}

	return list
}

func (it *MappedInfoItems) HasFilterFuncAsync(
	hasFilterFunc HasKeyFilterFunc,
) (isSuccess bool) {
	length := it.Length()
	if length == 0 {
		return false
	}

	if length <= consts.NonAsyncSafeRange {
		for key, item := range it.Items {
			if hasFilterFunc(key, item) {
				return true
			}
		}
	}

	isFound := false
	wg := sync.WaitGroup{}

	hasCheckerFunc := func(key string) {
		if isFound {
			return
		}

		defer wg.Done()

		if hasFilterFunc(key, it.Items[key]) {
			isFound = true
		}
	}

	for key := range it.Items {
		if isFound {
			break
		}

		wg.Add(1)
		go hasCheckerFunc(key)

		if isFound {
			break
		}
	}

	wg.Wait()

	return isFound
}

func (it *MappedInfoItems) Filter(
	filterFunc FilterFunc,
) []*Info {
	list := make([]*Info, 0, it.Length())

	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
		}

		if isBreak {
			return list
		}
	}

	return list
}

func (it *MappedInfoItems) TakeAllFilter(
	limit int,
	filterFunc TakeAllFilterFunc,
) []*Info {
	list := make([]*Info, it.Length())
	wg := sync.WaitGroup{}
	foundItems := 0
	hasLimit :=
		limit > constants.TakeAllMinusOne

	adderFunc := func(index int, key string) {
		defer wg.Done()

		item := it.Items[key]
		isTake := filterFunc(item)

		if isTake {
			list[index] = item
			foundItems++
		}
	}

	index := -1
	for key := range it.Items {
		index++
		if hasLimit && foundItems >= limit {
			break
		}

		wg.Add(1)
		go adderFunc(index, key)
	}

	wg.Wait()

	finalList := make([]*Info, 0, foundItems+constants.Capacity2)

	for _, item := range list {
		if item == nil {
			continue
		}

		finalList = append(
			finalList,
			item)
	}

	return finalList
}

func (it *MappedInfoItems) TakeAllFilterCollection(
	limit int,
	filterFunc TakeAllFilterFunc,
) *InfoCollection {
	if it.IsEmpty() {
		return EmptyInfoCollection()
	}

	items := it.TakeAllFilter(limit, filterFunc)

	return &InfoCollection{Items: items}
}

func (it *MappedInfoItems) TakeAllFilterMapCollection(
	limit int,
	filterFunc TakeAllFilterFunc,
) *MappedInfoItems {
	emptyMap := EmptyMappedInfoItems()

	if it.IsEmpty() {
		return EmptyMappedInfoItems()
	}

	items := it.TakeAllFilter(limit, filterFunc)

	// todo improve this logic
	return emptyMap.Adds(items...)
}

func (it *MappedInfoItems) FilterWithLimit(
	limit int,
	filterFunc FilterFunc,
) []*Info {
	length := defaultcapacity.MaxLimit(
		it.Length(),
		limit)
	list := make(
		[]*Info,
		0,
		length)

	collectedItems := 0
	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
			collectedItems++
		}

		if isBreak {
			return list
		}

		if collectedItems >= length {
			return list
		}
	}

	return list
}

func (it *MappedInfoItems) FilterMappedInfoItems(
	filterFunc FilterFunc,
) *MappedInfoItems {
	list := it.Filter(filterFunc)

	traceCollection := NewMappedInfoItemsUsingInfoItems(
		list...)

	return traceCollection
}

func (it *MappedInfoItems) FilterExtMappedInfoItems(
	dotExtension string,
) *MappedInfoItems {
	return it.FilterMappedInfoItems(func(stateInfo *Info) (isTake, isBreak bool) {
		return stateInfo.LocationInfo().DotExtension == dotExtension, false
	})
}

func (it *MappedInfoItems) JsonStrings() []string {
	list := make([]string, it.Length())

	index := 0
	for _, item := range it.Items {
		list[index] = item.JsonString()
		index++
	}

	return list
}

func (it *MappedInfoItems) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *MappedInfoItems) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *MappedInfoItems) JoinLine() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *MappedInfoItems) IsEqualDefault(
	another *MappedInfoItems,
) bool {
	return it.IsEqual(
		false,
		false,
		false,
		false,
		another)
}

func (it *MappedInfoItems) IsEqual(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	another *MappedInfoItems,
) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	return it.IsEqualItems(
		isIgnoreModifiedTimeCompare,
		isIgnoreChmodCompare,
		isIgnoreChownCompare,
		isIgnoreCompareOnAnyEmpty,
		another.Items)
}

func (it *MappedInfoItems) IsEqualItems(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	itemsMap map[string]*Info,
) bool {
	if it == nil && itemsMap == nil {
		return true
	}

	if it == nil || itemsMap == nil {
		return false
	}

	if it.Length() != len(itemsMap) {
		return false
	}

	for key, item := range it.Items {
		anotherItem := itemsMap[key]
		isNotEqual := !item.IsEqual(
			isIgnoreModifiedTimeCompare,
			isIgnoreChmodCompare,
			isIgnoreChownCompare,
			isIgnoreCompareOnAnyEmpty, anotherItem)

		if isNotEqual {
			return false
		}
	}

	return true
}

func (it *MappedInfoItems) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it MappedInfoItems) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JoinLine()
}

func (it *MappedInfoItems) JsonModel() map[string]*Info {
	return it.Items
}

func (it *MappedInfoItems) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it MappedInfoItems) Json() corejson.Result {
	return corejson.New(it)
}

func (it MappedInfoItems) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *MappedInfoItems) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*MappedInfoItems, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyMappedInfoItems(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *MappedInfoItems) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *MappedInfoItems {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *MappedInfoItems) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *MappedInfoItems) AsJsoner() corejson.Jsoner {
	return it
}

func (it *MappedInfoItems) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *MappedInfoItems) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *MappedInfoItems) CompiledCollectionModel(isSortChecksum bool) CompiledMappedItemsModel {
	return NewCompiledMappedInfoItemsModel(isSortChecksum, it)
}

func (it *MappedInfoItems) Clear() *MappedInfoItems {
	it.Items = map[string]*Info{}

	return it
}

func (it *MappedInfoItems) Dispose() {
	it.Clear()
	it.Items = nil
}

func (it MappedInfoItems) Clone() MappedInfoItems {
	list := NewMappedInfoItems(it.Length())
	list.AddMapCollection(&it)

	return *list
}

func (it *MappedInfoItems) ClonePtr() *MappedInfoItems {
	if it == nil {
		return nil
	}

	list := NewMappedInfoItems(it.Length())

	return list.AddMapCollection(it)
}
