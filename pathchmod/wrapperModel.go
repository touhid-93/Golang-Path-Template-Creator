package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corejson"
)

type wrapperModel struct {
	DirChmod, FileChmod chmodhelper.RwxWrapper
	IsRecursive         bool
	IsSkipOnInvalid     bool
	IsContinueOnError   bool
	IsKeepExistingChmod bool
}

func (it *wrapperModel) Json() corejson.Result {
	return corejson.New(it)
}

func (it *wrapperModel) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *wrapperModel) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it wrapperModel) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
