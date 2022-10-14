package pathcompilertests

import "gitlab.com/evatix-go/core/coredata/corejson"

func DeserializedFieldsToMap(jsonResult *corejson.Result) (
	fieldsMap map[string]interface{},
	parsingErr error,
) {
	if jsonResult == nil || len(jsonResult.Bytes) == 0 {
		return map[string]interface{}{}, nil
	}

	fieldsMap = map[string]interface{}{}
	parsingErr = jsonResult.Deserialize(&fieldsMap)

	return fieldsMap, parsingErr
}
