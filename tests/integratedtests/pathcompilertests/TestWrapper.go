package pathcompilertests

import (
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/enum/osmixtype"
)

type TestWrapper struct {
	coretests.BaseTestCase
	SelectBy                      osmixtype.Variant
	NameAssert, DescriptionAssert string
	IsTestEnv                     bool
	RunsIn                        []osmixtype.Variant
}

func (it *TestWrapper) ExpectedAsDynamicMap() enumimpl.DynamicMap {
	conv, isSuccess := it.Expected().(enumimpl.DynamicMap)

	if !isSuccess {
		panic("expected type doesn't meet")
	}

	return conv
}
