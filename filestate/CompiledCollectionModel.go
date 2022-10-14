package filestate

import "gitlab.com/evatix-go/core/coredata/corejson"

type CompiledCollectionModel struct {
	Collection       InfoCollection
	CompiledChecksum string
}

func NewCompiledCollectionModel(
	isSortChecksum bool,
	collection *InfoCollection,
) CompiledCollectionModel {
	if collection.IsEmpty() {
		return CompiledCollectionModel{
			Collection:       *EmptyInfoCollection(),
			CompiledChecksum: "",
		}
	}

	return CompiledCollectionModel{
		Collection:       *collection,
		CompiledChecksum: collection.CompiledChecksumString(isSortChecksum),
	}
}

func (it CompiledCollectionModel) Json() corejson.Result {
	return corejson.New(it)
}

func (it CompiledCollectionModel) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it CompiledCollectionModel) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it CompiledCollectionModel) JsonModelAny() interface{} {
	return it
}

func (it *CompiledCollectionModel) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *CompiledCollectionModel) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *CompiledCollectionModel) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CompiledCollectionModel) Length() int {
	if it == nil {
		return 0
	}

	return it.Collection.Length()
}

func (it *CompiledCollectionModel) IsEmpty() bool {
	if it == nil {
		return true
	}

	return it.CompiledChecksum == "" || it.Collection.IsEmpty()
}

func (it *CompiledCollectionModel) HasCompiledChecksum() bool {
	if it == nil {
		return false
	}

	return it.CompiledChecksum != ""
}

func (it *CompiledCollectionModel) IsEqual(
	isCompareCollection bool,
	right *CompiledCollectionModel,
) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.CompiledChecksum != right.CompiledChecksum {
		return false
	}

	if isCompareCollection && !it.Collection.IsEqualDefault(&right.Collection) {
		return false
	}

	return true
}
