package filestate

import (
	"fmt"
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
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

type InfoCollection struct {
	Items []*Info `json:"Items,omitempty"`
}

func NewInfoCollection(capacity int) *InfoCollection {
	slice := make([]*Info, 0, capacity)

	return &InfoCollection{
		slice,
	}
}

func EmptyInfoCollection() *InfoCollection {
	return NewInfoCollection(constants.Zero)
}

func NewInfoCollectionUsingInfoItems(
	infoItems ...*Info,
) *InfoCollection {
	if infoItems == nil {
		return EmptyInfoCollection()
	}

	return &InfoCollection{
		infoItems,
	}
}

func NewInfoCollectionUsingFilePathsAsync(
	hashMethod hashas.Variant,
	isNormalize bool,
	filePaths ...string,
) (*InfoCollection, *errwrappers.Collection) {
	if len(filePaths) == 0 {
		return EmptyInfoCollection(), nil
	}

	slice := make([]*Info, len(filePaths))
	errCollection := errwrappers.Empty()
	locker := sync.Mutex{}
	wg := sync.WaitGroup{}

	adderFunc := func(index int, filePath string) {
		defer wg.Done()

		info, errWrap := NewInfo(
			hashMethod,
			isNormalize,
			filePath)

		slice[index] = info

		if errWrap.HasError() {
			locker.Lock()

			errCollection.AddWrapperPtr(errWrap)
			errCollection.AddRefOne(
				errtype.MissingPathsOrInvalidPaths,
				"Remaining info couldn't process",
				filePaths[index+1:])

			locker.Unlock()
		}
	}

	wg.Add(len(filePaths))
	for i, filePath := range filePaths {
		go adderFunc(i, filePath)
	}

	wg.Wait()

	return &InfoCollection{
		Items: slice,
	}, errCollection
}

func (it *InfoCollection) CleanupNonNull() *InfoCollection {
	slice := make([]*Info, 0, it.Length())

	for _, item := range it.Items {
		if item == nil {
			continue
		}

		slice = append(slice, item)
	}

	it.Items = slice

	return it
}

func (it *InfoCollection) Add(
	stateInfo *Info,
) *InfoCollection {
	it.Items = append(it.Items, stateInfo)

	return it
}

func (it *InfoCollection) Adds(
	infoItems ...*Info,
) *InfoCollection {
	if len(infoItems) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		infoItems...)

	return it
}

func (it *InfoCollection) AddsPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *InfoCollection {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid {
			continue
		}

		it.Items = append(
			it.Items,
			stateInfo)
	}

	return it
}

func (it *InfoCollection) AddFilesPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *InfoCollection {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid || stateInfo.IsDir() {
			continue
		}

		it.Items = append(
			it.Items,
			stateInfo)
	}

	return it
}

func (it *InfoCollection) AddDirsPtr(
	isSkipOnIssues bool,
	infoItems ...*Info,
) *InfoCollection {
	if len(infoItems) == 0 {
		return it
	}

	for _, stateInfo := range infoItems {
		if isSkipOnIssues && stateInfo.IsInvalid || stateInfo.IsFile {
			continue
		}

		it.Items = append(
			it.Items,
			stateInfo)
	}

	return it
}

func (it *InfoCollection) ConcatNew(
	additionalInfoItems ...*Info,
) *InfoCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalInfoItems...)
}

func (it *InfoCollection) ConcatNewPtr(
	additionalInfoItems ...*Info,
) *InfoCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		true,
		additionalInfoItems...)
}

func (it *InfoCollection) AddsIf(
	isAdd bool,
	infoItems ...*Info,
) *InfoCollection {
	if !isAdd {
		return it
	}

	return it.Adds(infoItems...)
}

func (it *InfoCollection) HasChecksum(hexChecksum string) bool {
	return it.HasFilterFuncAsync(func(index int, info *Info) (isSuccess bool) {
		return info.HexContentChecksum == hexChecksum
	})
}

func (it *InfoCollection) HasAnyChecksum(hexChecksums ...string) bool {
	findingChecksumsHashset := corestr.New.Hashset.StringsPtr(
		&hexChecksums)

	return it.HasFilterFuncAsync(func(index int, info *Info) (isSuccess bool) {
		return findingChecksumsHashset.Has(info.HexContentChecksum)
	})
}

func (it *InfoCollection) HasAllChecksum(hexChecksums ...string) bool {
	if len(hexChecksums) == 0 {
		return true
	}

	if it.IsEmpty() {
		return false
	}

	if len(hexChecksums) <= 2 {
		for _, checksum := range hexChecksums {
			if !it.HasChecksum(checksum) {
				return false
			}
		}
	}

	mappedHexChecksum := it.AllHexChecksumToFilePathMap()

	for _, checksum := range hexChecksums {
		_, has := mappedHexChecksum[checksum]

		if !has {
			return false
		}
	}

	return true
}

func (it *InfoCollection) InfoUsingFilePaths(filePaths ...string) *InfoCollection {
	if it.IsEmpty() {
		return EmptyInfoCollection()
	}

	slice := make([]*Info, 0, len(filePaths))
	filePathsToInfoMap := it.AllFilePathToInfoMap()

	for _, filePath := range filePaths {
		info, has := filePathsToInfoMap[filePath]

		if has {
			slice = append(slice, info)
		}
	}

	return &InfoCollection{Items: slice}
}

func (it *InfoCollection) InfoUsingChecksums(checksums ...string) *InfoCollection {
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

func (it *InfoCollection) PathStatSlice(isTakeOnlyValid bool) []*chmodhelper.PathExistStat {
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

func (it *InfoCollection) GetAllUsingChecksumCollection(
	findingHexChecksums ...string,
) *InfoCollection {
	if len(findingHexChecksums) == 0 || it.IsEmpty() {
		return EmptyInfoCollection()
	}

	findingChecksumHashset := corestr.New.Hashset.StringsPtr(
		&findingHexChecksums)

	return it.FilterInfoCollection(func(info *Info) (isTake, isBreak bool) {
		return findingChecksumHashset.Has(info.HexContentChecksum), false
	})
}

func (it *InfoCollection) HasAnyFilePath(filePath ...string) bool {
	findingItemsHashset := corestr.New.Hashset.StringsPtr(
		&filePath)

	return it.HasFilterFuncAsync(func(index int, info *Info) (isSuccess bool) {
		return findingItemsHashset.Has(info.FullPath)
	})
}

func (it *InfoCollection) AllChecksums() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, item := range it.Items {
		slice[i] = item.HexContentChecksum
	}

	return slice
}

func (it *InfoCollection) AllChecksumsSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.AllChecksums()
	sort.Strings(slice)

	return slice
}

func (it *InfoCollection) AllFilePaths() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, item := range it.Items {
		slice[i] = item.FullPath
	}

	return slice
}

func (it *InfoCollection) AllFilePathsSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.AllFilePaths()
	sort.Strings(slice)

	return slice
}

func (it *InfoCollection) KeyValueStringMapUsingFmtFunc(
	fmtFunc MapKeyValFmtFunc,
) map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	itemsMap := make(
		map[string]string,
		it.Length())

	for _, item := range it.Items {
		k, v := fmtFunc(item)

		itemsMap[k] = v
	}

	return itemsMap
}

func (it *InfoCollection) KeyStringValueInfoMapUsingFmtFunc(
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
func (it *InfoCollection) AllFilePathToHexChecksumMap() map[string]string {
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
func (it *InfoCollection) AllFilePathToInfoMap() map[string]*Info {
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
func (it *InfoCollection) AllHexChecksumToInfoMap() map[string]*Info {
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
func (it *InfoCollection) AllHexChecksumToFilePathMap() map[string]string {
	if it.IsEmpty() {
		return map[string]string{}
	}

	itemsMap := make(map[string]string, it.Length())

	for _, item := range it.Items {
		itemsMap[item.HexContentChecksum] = item.FullPath
	}

	return itemsMap
}

func (it *InfoCollection) CompiledChecksum(isSortChecksum bool) *errstr.Result {
	if it.IsEmpty() {
		return errstr.Empty.Result()
	}

	return hexchecksum.OfChecksums(
		true,
		isSortChecksum,
		DefaultHashMethod,
		it.AllChecksums()...)
}

func (it *InfoCollection) CompiledChecksumString(isSortChecksum bool) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	result := hexchecksum.OfChecksums(
		true,
		isSortChecksum,
		DefaultHashMethod,
		it.AllChecksums()...)

	result.ErrorWrapper.Dispose()

	return result.Value
}

func (it *InfoCollection) InsertAt(index int, item *Info) *InfoCollection {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = item

	return it
}

func (it *InfoCollection) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *InfoCollection) First() *Info {
	return it.Items[0]
}

func (it *InfoCollection) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *InfoCollection) Last() *Info {
	return it.Items[it.LastIndex()]
}

func (it *InfoCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *InfoCollection) FirstOrDefault() *Info {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

func (it *InfoCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *InfoCollection) LastOrDefault() *Info {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

func (it *InfoCollection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Items[skippingItemsCount:]
}

func (it *InfoCollection) Skip(skippingItemsCount int) []*Info {
	return it.Items[skippingItemsCount:]
}

func (it *InfoCollection) SkipCollection(skippingItemsCount int) *InfoCollection {
	return &InfoCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *InfoCollection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Items[:takeDynamicItems]
}

func (it *InfoCollection) Take(takeDynamicItems int) []*Info {
	return it.Items[:takeDynamicItems]
}

func (it *InfoCollection) TakeCollection(takeDynamicItems int) *InfoCollection {
	return &InfoCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *InfoCollection) LimitCollection(limit int) *InfoCollection {
	return &InfoCollection{
		Items: it.Items[:limit],
	}
}

func (it *InfoCollection) SafeLimitCollection(limit int) *InfoCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &InfoCollection{
		Items: it.Items[:limit],
	}
}

func (it *InfoCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *InfoCollection) Limit(limit int) []*Info {
	return it.Take(limit)
}

func (it *InfoCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *InfoCollection) GetPagedCollection(
	eachPageSize int,
) []*InfoCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*InfoCollection{}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*InfoCollection, pagesPossibleCeiling)

	wg := sync.WaitGroup{}
	addPagedItemsFunc := func(oneBasedPageIndex int) {
		pagedCollection := it.GetSinglePageCollection(
			eachPageSize,
			oneBasedPageIndex,
		)

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
func (it *InfoCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *InfoCollection {
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

	list := it.Items[skipItems:endingIndex]

	return NewInfoCollectionUsingInfoItems(list...)
}

func (it *InfoCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *InfoCollection) Count() int {
	return it.Length()
}

func (it *InfoCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *InfoCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *InfoCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *InfoCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *InfoCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *InfoCollection) StringsUsingStringerFmtFunc(
	fmtFunc StringerFmtFunc,
) []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = fmtFunc(i, item)
	}

	return list
}

func (it *InfoCollection) HasFilterFuncAsync(
	hasFilterFunc HasFilterFunc,
) (isSuccess bool) {
	length := it.Length()
	if length == 0 {
		return false
	}

	if length <= consts.NonAsyncSafeRange {
		for i, item := range it.Items {
			if hasFilterFunc(i, item) {
				return true
			}
		}
	}

	isFound := false
	wg := sync.WaitGroup{}

	hasCheckerFunc := func(index int) {
		if isFound {
			return
		}

		defer wg.Done()

		if hasFilterFunc(index, it.Items[index]) {
			isFound = true
		}
	}

	for i := 0; i < length; i++ {
		if isFound {
			break
		}

		wg.Add(1)
		go hasCheckerFunc(i)

		if isFound {
			break
		}
	}

	wg.Wait()

	return isFound
}

func (it *InfoCollection) Filter(
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

func (it *InfoCollection) TakeAllFilter(
	limit int,
	filterFunc TakeAllFilterFunc,
) []*Info {
	list := make([]*Info, it.Length())
	wg := sync.WaitGroup{}
	foundItems := 0
	hasLimit :=
		limit > constants.TakeAllMinusOne

	adderFunc := func(index int) {
		defer wg.Done()

		item := it.Items[index]
		isTake := filterFunc(item)

		if isTake {
			list[index] = item
			foundItems++
		}
	}

	maxItems := defaultcapacity.MaxLimit(
		it.Length(),
		limit)

	for i := 0; i < maxItems; i++ {
		if hasLimit && foundItems >= limit {
			break
		}

		wg.Add(1)
		go adderFunc(i)
	}

	wg.Wait()

	finalList := make([]*Info, 0, foundItems+constants.Capacity2)

	for _, item := range list {
		if item == nil {
			continue
		}

		if hasLimit && len(finalList) >= limit {
			break
		}

		finalList = append(
			finalList,
			item)
	}

	return finalList
}

func (it *InfoCollection) TakeAllFilterCollection(
	limit int,
	filterFunc TakeAllFilterFunc,
) *InfoCollection {
	if it.IsEmpty() {
		return EmptyInfoCollection()
	}

	items := it.TakeAllFilter(limit, filterFunc)

	return &InfoCollection{Items: items}
}

func (it *InfoCollection) FilterWithLimit(
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

func (it *InfoCollection) FilterInfoCollection(
	filterFunc FilterFunc,
) *InfoCollection {
	list := it.Filter(filterFunc)

	traceCollection := NewInfoCollectionUsingInfoItems(
		list...)

	return traceCollection
}

func (it *InfoCollection) FilterExtInfoCollection(
	dotExtension string,
) *InfoCollection {
	return it.FilterInfoCollection(func(stateInfo *Info) (isTake, isBreak bool) {
		return stateInfo.LocationInfo().DotExtension == dotExtension, false
	})
}

func (it *InfoCollection) Reverse() *InfoCollection {
	length := it.Length()

	if length <= 1 {
		return it
	}

	if length == 2 {
		it.Items[0], it.Items[1] = it.Items[1], it.Items[0]

		return it
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		it.Items[i], it.Items[lastIndex-i] =
			it.Items[lastIndex-i], it.Items[i]
	}

	return it
}

func (it *InfoCollection) New(
	isNormalize,
	isSortedFilePaths bool,
) (*InfoCollection, *errwrappers.Collection) {
	if it.IsEmpty() {
		return nil, nil
	}

	filePaths := it.AllFilePaths()

	if isSortedFilePaths {
		sort.Strings(filePaths)
	}

	return NewInfoCollectionUsingFilePathsAsync(
		it.First().HashMethod,
		isNormalize,
		filePaths...)
}

func (it *InfoCollection) ReadCurrentHexChecksumsStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.ReadCurrentHexChecksumString()
	}

	return list
}

// ReadCurrentHexChecksumsMapStrings
//
// Key => FilePath
//
// Value => HexChecksum
func (it *InfoCollection) ReadCurrentHexChecksumsMapStrings() map[string]string {
	mappedItems := make(map[string]string, it.Length())

	for _, item := range it.Items {
		mappedItems[item.FullPath] = item.ReadCurrentHexChecksumString()
	}

	return mappedItems
}

func (it *InfoCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *InfoCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *InfoCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *InfoCollection) JoinLine() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *InfoCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *InfoCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *InfoCollection) IsEqualDefault(
	another *InfoCollection,
) bool {
	return it.IsEqual(
		false,
		false,
		false,
		false,
		another)
}

func (it *InfoCollection) IsEqual(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	another *InfoCollection,
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
		another.Items...)
}

func (it *InfoCollection) IsEqualItems(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	lines ...*Info,
) bool {
	if it == nil && lines == nil {
		return true
	}

	if it == nil || lines == nil {
		return false
	}

	if it.Length() != len(lines) {
		return false
	}

	for i, item := range it.Items {
		anotherItem := lines[i]
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

func (it *InfoCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it InfoCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JoinLine()
}

func (it *InfoCollection) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String())
	}

	return newSlice
}

func (it *InfoCollection) JsonModel() []*Info {
	return it.Items
}

func (it *InfoCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it InfoCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it InfoCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *InfoCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*InfoCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyInfoCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *InfoCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *InfoCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *InfoCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *InfoCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *InfoCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *InfoCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *InfoCollection) Clear() *InfoCollection {
	it.Items = it.Items[:0]

	return it
}

func (it *InfoCollection) Dispose() {
	it.Clear()
	it.Items = nil
}

func (it *InfoCollection) CompiledCollectionModel(isSortChecksum bool) CompiledCollectionModel {
	return NewCompiledCollectionModel(isSortChecksum, it)
}

func (it InfoCollection) Clone() InfoCollection {
	list := NewInfoCollection(it.Length())

	return *list.Adds(it.Items...)
}

func (it *InfoCollection) ClonePtr() *InfoCollection {
	if it == nil {
		return nil
	}

	list := NewInfoCollection(it.Length())

	return list.Adds(it.Items...)
}
