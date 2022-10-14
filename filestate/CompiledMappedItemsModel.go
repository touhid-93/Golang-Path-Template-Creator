package filestate

import "gitlab.com/evatix-go/core/coredata/corejson"

type CompiledMappedItemsModel struct {
	MappedInfoItems  MappedInfoItems
	CompiledChecksum string
}

func NewCompiledMappedItemsModel(
	isSortChecksum bool,
	mappedItems map[string]*Info,
) CompiledMappedItemsModel {
	if len(mappedItems) == 0 {
		return CompiledMappedItemsModel{
			MappedInfoItems:  MappedInfoItems{},
			CompiledChecksum: "",
		}
	}

	mappedInfoItems := MappedInfoItems{Items: mappedItems}

	return CompiledMappedItemsModel{
		MappedInfoItems:  mappedInfoItems,
		CompiledChecksum: mappedInfoItems.CompiledChecksumString(isSortChecksum),
	}
}

func NewCompiledMappedInfoItemsModel(
	isSortChecksum bool,
	mappedItems *MappedInfoItems,
) CompiledMappedItemsModel {
	if mappedItems.IsEmpty() {
		return CompiledMappedItemsModel{
			MappedInfoItems:  MappedInfoItems{},
			CompiledChecksum: "",
		}
	}

	return CompiledMappedItemsModel{
		MappedInfoItems:  *mappedItems,
		CompiledChecksum: mappedItems.CompiledChecksumString(isSortChecksum),
	}
}

func (it CompiledMappedItemsModel) Json() corejson.Result {
	return corejson.New(it)
}

func (it CompiledMappedItemsModel) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it CompiledMappedItemsModel) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it CompiledMappedItemsModel) JsonModelAny() interface{} {
	return it
}

func (it *CompiledMappedItemsModel) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *CompiledMappedItemsModel) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *CompiledMappedItemsModel) AsJsoner() corejson.Jsoner {
	return it
}

func (it *CompiledMappedItemsModel) Length() int {
	if it == nil {
		return 0
	}

	return it.MappedInfoItems.Length()
}

func (it *CompiledMappedItemsModel) IsEmpty() bool {
	if it == nil {
		return true
	}

	return it.CompiledChecksum == "" || it.MappedInfoItems.IsEmpty()
}

func (it *CompiledMappedItemsModel) HasCompiledChecksum() bool {
	if it == nil {
		return false
	}

	return it.CompiledChecksum != ""
}

func (it *CompiledMappedItemsModel) IsEqual(
	isCompareCollection bool,
	right *CompiledMappedItemsModel,
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

	if isCompareCollection && !it.MappedInfoItems.IsEqualDefault(&right.MappedInfoItems) {
		return false
	}

	return true
}
