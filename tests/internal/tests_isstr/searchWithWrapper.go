package tests_isstr

import "gitlab.com/evatix-go/core/coretests"

type searchWithWrapper struct {
	baseDir, search string
	isIgnoreCase    bool
	Expectation     bool
	actual          bool
	funcName        coretests.TestFuncName
}

func (receiver *searchWithWrapper) FuncName() string {
	return receiver.funcName.Value()
}

func (receiver *searchWithWrapper) Value() interface{} {
	return receiver
}

func (receiver *searchWithWrapper) Expected() interface{} {
	return receiver.Expectation
}

func (receiver *searchWithWrapper) Actual() interface{} {
	return receiver.actual
}

func (receiver *searchWithWrapper) SetActual(actual bool) {
	receiver.actual = actual
}

func (receiver *searchWithWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	return coretests.TestCaseMessenger(receiver)
}
