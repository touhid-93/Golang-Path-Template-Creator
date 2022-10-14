package oldtests

import "reflect"

func GetFieldValue(field reflect.Value) interface{} {
	isPtr := field.Kind() == reflect.Ptr
	isNil := isPtr && field.IsNil()
	if isNil {
		return "nil"
	}

	if !field.IsValid() && !isPtr {
		return nil
	}

	if field.IsValid() && field.CanAddr() && !isPtr {
		return field
	}

	if field.CanInterface() {
		return field.Interface()
	}

	if isPtr && !isNil {
		return field.Elem()
	}

	return field
}
