package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type pathTestCaseDataWrapper struct {
	OSName, funcName, expected string
}

func getPathTestCommonMethodLinux(
	t *testing.T,
	testData pathTestCaseDataWrapper,
	callingFunctionToBeTested func() string,
) {
	// Arrange
	SkipOnWindows(t)

	testMessage := fmt.Sprintf("(%s) [%s] inputs() expects string output", testData.OSName, testData.funcName)

	Convey(testMessage, t, func() {
		// Act
		actual := callingFunctionToBeTested()

		// Assert
		So(actual, ShouldEqual, testData.expected)
	})
}

func pathTestCaseInternalFromWrappers(
	t *testing.T,
	testData []pathTestCaseDataWrapper,
	actualFuncCall func() string,
) {
	for _, testCase := range testData {
		getPathTestCommonMethodLinux(
			t,
			testCase,
			actualFuncCall)
	}
}
