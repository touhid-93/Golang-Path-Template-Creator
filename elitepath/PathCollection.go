package elitepath

import (
	"encoding/json"
	"regexp"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathfixer"
)

type PathCollection struct {
	Items []*Path `json:"Items,omitempty"`
}

func EmptyPathCollection() *PathCollection {
	return &PathCollection{Items: []*Path{}}
}

func NewPathCollectionUsingCapacity(capacity int) *PathCollection {
	paths := make([]*Path, 0, capacity)

	return &PathCollection{
		Items: paths,
	}
}

func NewPathCollection(
	isSkipEmpty,
	isNormalize,
	isExpandEnvVars bool,
	overridingPathOptions *pathfixer.PathOptions,
	collection ...string,
) *PathCollection {
	if len(collection) == 0 {
		return EmptyPathCollection()
	}

	if overridingPathOptions == nil {
		overridingPathOptions = &pathfixer.PathOptions{
			IsContinueOnError: false,
			IsNormalize:       isNormalize,
			IsExpandEnvVar:    isExpandEnvVars,
			IsRecursive:       false,
			IsSkipOnInvalid:   false,
			IsSkipOnExist:     false,
			IsSkipOnEmpty:     false,
			IsRelative:        false,
		}
	}

	paths := make([]*Path, len(collection))

	for i, path := range collection {
		if isSkipEmpty && path == "" {
			continue
		}

		paths[i] = &Path{
			Location: *pathfixer.NewLocationUsingOptions(
				path,
				*overridingPathOptions),
		}
	}

	return &PathCollection{Items: paths}
}

func NewPathCollectionDirect(
	options *pathfixer.PathOptions,
	collection ...string,
) *PathCollection {
	if len(collection) == 0 {
		return EmptyPathCollection()
	}

	paths := make([]*Path, len(collection))

	if options == nil {
		options = &pathfixer.PathOptions{
			IsContinueOnError: false,
			IsNormalize:       false,
			IsExpandEnvVar:    false,
			IsRecursive:       false,
			IsSkipOnInvalid:   false,
			IsSkipOnExist:     false,
			IsSkipOnEmpty:     false,
			IsRelative:        false,
		}
	}

	for i, path := range collection {
		paths[i] = &Path{
			Location: *pathfixer.NewLocationUsingOptions(
				path,
				*options),
		}
	}

	return &PathCollection{Items: paths}
}

func (it *PathCollection) FirstDynamic() interface{} {
	return it.Items[0]
}

func (it *PathCollection) First() *Path {
	return it.Items[0]
}

func (it *PathCollection) LastDynamic() interface{} {
	return it.Items[it.LastIndex()]
}

func (it *PathCollection) Last() *Path {
	return it.Items[it.LastIndex()]
}

func (it *PathCollection) FirstOrDefault() *Path {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[0]
}

func (it *PathCollection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *PathCollection) LastOrDefault() interface{} {
	if it.IsEmpty() {
		return nil
	}

	return it.Items[it.LastIndex()]
}

func (it *PathCollection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *PathCollection) Skip(skippingItemsCount int) *PathCollection {
	return &PathCollection{
		Items: it.Items[skippingItemsCount:],
	}
}

func (it *PathCollection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Skip(skippingItemsCount)
}

func (it *PathCollection) Take(takeDynamicItems int) *PathCollection {
	return &PathCollection{
		Items: it.Items[:takeDynamicItems],
	}
}

func (it *PathCollection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Take(takeDynamicItems)
}

func (it *PathCollection) Limit(limit int) *PathCollection {
	return it.Take(limit)
}

func (it *PathCollection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *PathCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *PathCollection) Add(path *Path) *PathCollection {
	if path == nil {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) AddValid(path *Path) *PathCollection {
	if path == nil || !path.IsPathExist() {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) ExecuteAll(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	exeFunc func(path *Path) *errorwrapper.Wrapper,
) (isSuccess bool) {
	if it.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()

	for _, item := range it.Items {
		err := exeFunc(item)
		errCollection.AddWrapperPtr(err)

		if !isContinueOnError && err.HasError() {
			return false
		}
	}

	return stateTracker.IsSuccess()
}

func (it *PathCollection) AddIf(
	isAdd bool, path *Path,
) *PathCollection {
	if !isAdd || path == nil {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) AddDir(path *Path) *PathCollection {
	if path == nil || !path.IsDir() {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) AddFile(
	path *Path,
) *PathCollection {
	if path == nil || !path.IsFile() {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) AddFilterMatch(
	filter *Filter, path *Path,
) *PathCollection {
	if path == nil || !filter.IsMatch(path) {
		return it
	}

	it.Items = append(it.Items, path)

	return it
}

func (it *PathCollection) AddsFilterMatch(
	filter *Filter, paths ...*Path,
) *PathCollection {
	if len(paths) == 0 {
		return it
	}

	for _, path := range paths {
		if !filter.IsMatch(path) {
			continue
		}

		it.Items = append(it.Items, path)
	}

	return it
}

func (it *PathCollection) AddsRawValid(paths ...string) *PathCollection {
	if len(paths) == 0 {
		return it
	}

	for _, rawPath := range paths {
		path := &Path{
			Location: *pathfixer.NewLocation(rawPath),
		}

		if !path.IsPathExist() {
			continue
		}

		it.Items = append(it.Items, path)
	}

	return it
}

func (it *PathCollection) AddsRaw(paths ...string) *PathCollection {
	if len(paths) == 0 {
		return it
	}

	for _, path := range paths {
		it.Items = append(it.Items, &Path{
			Location: *pathfixer.NewLocation(path),
		})
	}

	return it
}

func (it *PathCollection) AddsRawUsingOption(
	options pathfixer.PathOptions,
	paths ...string,
) *PathCollection {
	if len(paths) == 0 {
		return it
	}

	for _, path := range paths {
		it.Items = append(it.Items, &Path{
			Location: *pathfixer.NewLocationUsingOptions(path, options),
		})
	}

	return it
}

func (it *PathCollection) AddCollections(
	collections ...*PathCollection,
) *PathCollection {
	if len(collections) == 0 {
		return it
	}

	for _, collection := range collections {
		if collection.IsEmpty() {
			continue
		}

		for _, item := range collection.Items {
			it.Items = append(it.Items, item)
		}
	}

	return it
}

func (it *PathCollection) ConcatNew(
	collections ...*PathCollection,
) *PathCollection {
	length := it.Length()

	for _, collection := range collections {
		length += collection.Length()
	}

	newPaths := make([]*Path, 0, length+5)
	newPaths = append(newPaths, it.Items...)

	for _, collection := range collections {
		if collection.IsEmpty() {
			continue
		}

		newPaths = append(
			newPaths,
			collection.Items...)
	}

	return &PathCollection{Items: newPaths}
}

func (it *PathCollection) Count() int {
	return it.Length()
}

func (it *PathCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *PathCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *PathCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *PathCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *PathCollection) Filter(filter *Filter) []string {
	newPaths := make([]string, 0, it.Length())

	for _, item := range it.Items {
		if filter.IsMatch(item) {
			newPaths = append(newPaths, item.CompiledPath())
		}
	}

	return newPaths
}

func (it *PathCollection) FilterPathCollection(filter *Filter) *PathCollection {
	newPaths := make([]*Path, 0, it.Length())

	for _, item := range it.Items {
		if filter.IsMatch(item) {
			newPaths = append(newPaths, item)
		}
	}

	return &PathCollection{Items: newPaths}
}

func (it *PathCollection) PathMatchesWithRegex(regexp *regexp.Regexp) []string {
	newPaths := make([]string, 0, it.Length())

	for _, item := range it.Items {
		compiledPath := item.CompiledPath()
		if regexp.MatchString(compiledPath) {
			newPaths = append(newPaths, compiledPath)
		}
	}

	return newPaths
}

func (it *PathCollection) Strings() []string {
	newPaths := make([]string, it.Length())

	for i, item := range it.Items {
		newPaths[i] = item.CompiledPath()
	}

	return newPaths
}

func (it *PathCollection) AllExistPaths() []string {
	newPaths := make([]string, 0, it.Length())

	for _, item := range it.Items {
		if item.IsPathExist() {
			newPaths = append(newPaths, item.CompiledPath())
		}
	}

	return newPaths
}

func (it *PathCollection) AllExistPathsCollection() *PathCollection {
	newPaths := make([]*Path, 0, it.Length())

	for _, item := range it.Items {
		if item.IsPathExist() {
			newPaths = append(newPaths, item)
		}
	}

	return &PathCollection{Items: newPaths}
}

func (it *PathCollection) AllDirs() []string {
	newPaths := make([]string, 0, it.Length())

	for _, item := range it.Items {
		if item.IsDir() {
			newPaths = append(newPaths, item.CompiledPath())
		}
	}

	return newPaths
}

func (it *PathCollection) AllDirsPaths() *PathCollection {
	newPaths := make([]*Path, 0, it.Length())

	for _, item := range it.Items {
		if item.IsDir() {
			newPaths = append(newPaths, item)
		}
	}

	return &PathCollection{Items: newPaths}
}

func (it *PathCollection) AllFiles() []string {
	newPaths := make([]string, 0, it.Length())

	for _, item := range it.Items {
		if item.IsFile() {
			newPaths = append(newPaths, item.CompiledPath())
		}
	}

	return newPaths
}

func (it *PathCollection) AllFilesPaths() *PathCollection {
	newPaths := make([]*Path, 0, it.Length())

	for _, item := range it.Items {
		if item.IsFile() {
			newPaths = append(newPaths, item)
		}
	}

	return &PathCollection{Items: newPaths}
}

func (it *PathCollection) SaveToFile(
	filePath string,
) *errorwrapper.Wrapper {
	return fs.WriteJsonResultUsingLock(
		true,
		false,
		it.JsonPtr(),
		filePath,
	)
}

func (it *PathCollection) ReadFromFile(
	filePath string,
) *errorwrapper.Wrapper {
	return fs.ReadJsonParseSelfInjectorUsingLock(
		filePath,
		it)
}

func (it PathCollection) String() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *PathCollection) JsonModel() *PathCollection {
	return it
}

func (it *PathCollection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *PathCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *PathCollection) UnmarshalJSON(
	data []byte,
) error {
	var dataModel PathCollection
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.Items = dataModel.Items
	}

	return err
}

func (it PathCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it PathCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *PathCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PathCollection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyPathCollection(), defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(jsonResult.Bytes, &it)

	if err != nil {
		return EmptyPathCollection(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *PathCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PathCollection {
	parsedResult, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return parsedResult
}

func (it *PathCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *PathCollection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *PathCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *PathCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *PathCollection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *PathCollection) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
