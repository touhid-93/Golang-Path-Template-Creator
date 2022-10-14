package fileinfopath

import (
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

type InstanceCollection struct {
	Items []*Instance `json:"Items,omitempty"`
}

func NewInstanceCollection(capacity int) *InstanceCollection {
	slice := make([]*Instance, 0, capacity)

	return &InstanceCollection{
		slice,
	}
}

func EmptyInstanceCollection() *InstanceCollection {
	return NewInstanceCollection(constants.Zero)
}

func NewInstanceCollectionUsingInstanceItems(
	infoItems ...*Instance,
) *InstanceCollection {
	if infoItems == nil {
		return EmptyInstanceCollection()
	}

	return &InstanceCollection{
		infoItems,
	}
}

func NewInstanceCollectionUsingFilePathsAsync(
	filePaths ...string,
) *InstanceCollection {
	if len(filePaths) == 0 {
		return EmptyInstanceCollection()
	}

	slice := make([]*Instance, len(filePaths))
	isErrorFound := false
	wg := sync.WaitGroup{}

	adderFunc := func(index int, filePath string) {
		defer wg.Done()

		if isErrorFound {
			return
		}

		instance := New(filePath)

		slice[index] = instance
	}

	wg.Add(len(filePaths))
	for i, filePath := range filePaths {
		go adderFunc(i, filePath)
	}

	wg.Wait()

	return &InstanceCollection{
		Items: slice,
	}
}

func (it *InstanceCollection) CleanupNonNull() *InstanceCollection {
	slice := make([]*Instance, 0, it.Length())

	for _, item := range it.Items {
		if item == nil {
			continue
		}

		slice = append(slice, item)
	}

	it.Items = slice

	return it
}

func (it *InstanceCollection) Add(
	instanceItem *Instance,
) *InstanceCollection {
	it.Items = append(it.Items, instanceItem)

	return it
}

func (it *InstanceCollection) Adds(
	infoItems ...*Instance,
) *InstanceCollection {
	if len(infoItems) == 0 {
		return it
	}

	it.Items = append(
		it.Items,
		infoItems...)

	return it
}

func (it *InstanceCollection) AddsPtr(
	isSkipOnIssues bool,
	instanceItems ...*Instance,
) *InstanceCollection {
	if len(instanceItems) == 0 {
		return it
	}

	for _, instanceItem := range instanceItems {
		if isSkipOnIssues && instanceItem.IsInvalidPath() {
			continue
		}

		it.Items = append(
			it.Items,
			instanceItem)
	}

	return it
}

func (it *InstanceCollection) AddFilesPtr(
	isSkipOnIssues bool,
	infoItems ...*Instance,
) *InstanceCollection {
	if len(infoItems) == 0 {
		return it
	}

	for _, instanceItem := range infoItems {
		if isSkipOnIssues && instanceItem.IsInvalidFileInfo() || instanceItem.IsDir() {
			continue
		}

		it.Items = append(
			it.Items,
			instanceItem)
	}

	return it
}

func (it *InstanceCollection) AddDirsPtr(
	isSkipOnIssues bool,
	infoItems ...*Instance,
) *InstanceCollection {
	if len(infoItems) == 0 {
		return it
	}

	for _, instanceItem := range infoItems {
		if isSkipOnIssues && instanceItem.IsInvalidPath() || instanceItem.IsFile() {
			continue
		}

		it.Items = append(
			it.Items,
			instanceItem)
	}

	return it
}

func (it *InstanceCollection) ConcatNew(
	additionalInstanceItems ...*Instance,
) *InstanceCollection {
	cloned := it.Clone()

	return cloned.Adds(additionalInstanceItems...)
}

func (it *InstanceCollection) ConcatNewPtr(
	additionalInstanceItems ...*Instance,
) *InstanceCollection {
	cloned := it.Clone()

	return cloned.AddsPtr(
		true,
		additionalInstanceItems...)
}

func (it *InstanceCollection) AddsIf(
	isAdd bool,
	infoItems ...*Instance,
) *InstanceCollection {
	if !isAdd {
		return it
	}

	return it.Adds(infoItems...)
}

func (it *InstanceCollection) KeyValueStringMapUsingFmtFunc(
	fmtFunc MapKeyValueStringFmtFunc,
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

func (it *InstanceCollection) KeyStringMapUsingFmtFunc(
	fmtFunc MapKeyStringValueInstanceFmtFunc,
) map[string]*Instance {
	if it.IsEmpty() {
		return map[string]*Instance{}
	}

	itemsMap := make(
		map[string]*Instance,
		it.Length())

	for _, item := range it.Items {
		k := fmtFunc(item)

		itemsMap[k] = item
	}

	return itemsMap
}

// AllFilePathToInstanceMap
//
// Key = filePath,
//
// Value = Instance
func (it *InstanceCollection) AllFilePathToInstanceMap() map[string]*Instance {
	if it.IsEmpty() {
		return map[string]*Instance{}
	}

	itemsMap := make(map[string]*Instance, it.Length())

	for _, item := range it.Items {
		itemsMap[item.FullPath] = item
	}

	return itemsMap
}

func (it *InstanceCollection) AllFileInfos() []*Instance {
	return it.Items
}

func (it *InstanceCollection) AllFilePaths() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, item := range it.Items {
		slice[i] = item.FullPath
	}

	return slice
}

func (it *InstanceCollection) AllFilePathsSorted() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.AllFilePaths()
	sort.Strings(slice)

	return slice
}

func (it *InstanceCollection) InsertAt(index int, item *Instance) *InstanceCollection {
	it.Items = append(it.Items[:index+1], it.Items[index:]...)
	it.Items[index] = item

	return it
}

func (it *InstanceCollection) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *InstanceCollection) First() *Instance {
	return it.Items[0]
}

func (it *InstanceCollection) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *InstanceCollection) Last() *Instance {
	return it.Items[it.LastIndex()]
}

func (it *InstanceCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *InstanceCollection) FirstOrDefault() *Instance {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

func (it *InstanceCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *InstanceCollection) LastOrDefault() *Instance {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

func (it *InstanceCollection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Items[skippingItemsCount:]
}

func (it *InstanceCollection) Skip(skippingItemsCount int) []*Instance {
	return it.Items[skippingItemsCount:]
}

func (it *InstanceCollection) SkipCollection(skippingItemsCount int) *InstanceCollection {
	return &InstanceCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *InstanceCollection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Items[:takeDynamicItems]
}

func (it *InstanceCollection) Take(takeDynamicItems int) []*Instance {
	return it.Items[:takeDynamicItems]
}

func (it *InstanceCollection) TakeCollection(takeDynamicItems int) *InstanceCollection {
	return &InstanceCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *InstanceCollection) LimitCollection(limit int) *InstanceCollection {
	return &InstanceCollection{
		Items: it.Items[:limit],
	}
}

func (it *InstanceCollection) SafeLimitCollection(limit int) *InstanceCollection {
	limit = defaultcapacity.
		MaxLimit(it.Length(), limit)

	return &InstanceCollection{
		Items: it.Items[:limit],
	}
}

func (it *InstanceCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *InstanceCollection) Limit(limit int) []*Instance {
	return it.Take(limit)
}

func (it *InstanceCollection) GetPagesSize(
	eachPageSize int,
) int {
	length := it.Length()

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))

	return pagesPossibleCeiling
}

func (it *InstanceCollection) GetPagedCollection(
	eachPageSize int,
) []*InstanceCollection {
	length := it.Length()

	if length < eachPageSize {
		return []*InstanceCollection{}
	}

	pagesPossibleFloat := float64(length) / float64(eachPageSize)
	pagesPossibleCeiling := int(math.Ceil(pagesPossibleFloat))
	collectionOfCollection := make([]*InstanceCollection, pagesPossibleCeiling)

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
func (it *InstanceCollection) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *InstanceCollection {
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

	return NewInstanceCollectionUsingInstanceItems(list...)
}

func (it *InstanceCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *InstanceCollection) Count() int {
	return it.Length()
}

func (it *InstanceCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *InstanceCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *InstanceCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *InstanceCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *InstanceCollection) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *InstanceCollection) StringsUsingStringerFmtFunc(
	fmtFunc StringerFmtFunc,
) []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = fmtFunc(i, item)
	}

	return list
}

func (it *InstanceCollection) HasFilterFuncAsync(
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

func (it *InstanceCollection) Filter(
	filterFunc FilterFunc,
) []*Instance {
	list := make([]*Instance, 0, it.Length())

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

func (it *InstanceCollection) TakeAllFilter(
	limit int,
	filterFunc TakeAllFilterFunc,
) []*Instance {
	list := make([]*Instance, it.Length())
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

	finalList := make([]*Instance, 0, foundItems+constants.Capacity2)

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

func (it *InstanceCollection) TakeAllFilterCollection(
	limit int,
	filterFunc TakeAllFilterFunc,
) *InstanceCollection {
	if it.IsEmpty() {
		return EmptyInstanceCollection()
	}

	items := it.TakeAllFilter(limit, filterFunc)

	return &InstanceCollection{Items: items}
}

func (it *InstanceCollection) FilterWithLimit(
	limit int,
	filterFunc FilterFunc,
) []*Instance {
	length := defaultcapacity.MaxLimit(
		it.Length(),
		limit)
	list := make(
		[]*Instance,
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

func (it *InstanceCollection) FilterInstanceCollection(
	filterFunc FilterFunc,
) *InstanceCollection {
	list := it.Filter(filterFunc)

	traceCollection := NewInstanceCollectionUsingInstanceItems(
		list...)

	return traceCollection
}

func (it *InstanceCollection) FilterExtInstanceCollection(
	dotExtension string,
) *InstanceCollection {
	return it.FilterInstanceCollection(func(instanceItem *Instance) (isTake, isBreak bool) {
		if instanceItem.IsInvalidPath() {
			return false, false
		}

		return filepath.Ext(instanceItem.FileInfo.Name()) == dotExtension, false
	})
}

func (it *InstanceCollection) Reverse() *InstanceCollection {
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

func (it *InstanceCollection) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *InstanceCollection) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *InstanceCollection) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *InstanceCollection) JoinLine() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *InstanceCollection) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *InstanceCollection) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

func (it *InstanceCollection) IsEqualDefault(
	another *InstanceCollection,
) bool {
	return it.IsEqual(
		false,
		false,
		false,
		another)
}

func (it *InstanceCollection) IsEqual(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	another *InstanceCollection,
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
		isQuickVerifyOnPathEqual,
		isPathMustMatchIfDir,
		isVerifyContent,
		another.Items...)
}

func (it *InstanceCollection) IsEqualItems(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	lines ...*Instance,
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
			isQuickVerifyOnPathEqual,
			isPathMustMatchIfDir,
			isVerifyContent,
			anotherItem)

		if isNotEqual {
			return false
		}
	}

	return true
}

func (it *InstanceCollection) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JsonPtr().JsonString()
}

func (it InstanceCollection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.JoinLine()
}

func (it *InstanceCollection) CsvStrings() []string {
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

func (it *InstanceCollection) JsonModel() []*Instance {
	return it.Items
}

func (it *InstanceCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it InstanceCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it InstanceCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *InstanceCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*InstanceCollection, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return EmptyInstanceCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *InstanceCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *InstanceCollection {
	hashSet, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return hashSet
}

func (it *InstanceCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *InstanceCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *InstanceCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *InstanceCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *InstanceCollection) Clear() *InstanceCollection {
	it.Items = it.Items[:0]

	return it
}

func (it *InstanceCollection) Dispose() {
	it.Clear()
	it.Items = nil
}

func (it InstanceCollection) Clone() InstanceCollection {
	list := NewInstanceCollection(it.Length())

	return *list.Adds(it.Items...)
}

func (it *InstanceCollection) ClonePtr() *InstanceCollection {
	if it == nil {
		return nil
	}

	list := NewInstanceCollection(it.Length())

	return list.Adds(it.Items...)
}
