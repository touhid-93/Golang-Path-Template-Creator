package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper"
)

type combinedPathTestCaseWrapper struct {
	inputSeparator, expected, expectedMessage     string
	inputPaths1, inputPaths2, inputPaths3         string
	isIgnoreEmptyPath, isLongPathFix, isNormalize bool
}

var combinedPathTestCaseWrappers = []combinedPathTestCaseWrapper{
	{
		inputSeparator:    constants.BackSlash,
		inputPaths1:       "something",
		inputPaths2:       "more",
		inputPaths3:       "etc",
		isIgnoreEmptyPath: true,
		isNormalize:       true,
		expected:          "something\\more\\etc",
		expectedMessage:   "something\\more\\etc",
		isLongPathFix:     true,
	},
	{
		inputSeparator:    constants.ForwardSlash,
		inputPaths1:       "something",
		inputPaths2:       "more",
		inputPaths3:       "etc",
		isIgnoreEmptyPath: true,
		isNormalize:       true,
		expected:          "something/more/etc",
		expectedMessage:   "something/more/etc",
		isLongPathFix:     true,
	},
}

func TestGetCombinedPath(t *testing.T) {
	for i, testCase := range combinedPathTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetCombinedPath] inputs (%s, %v, %v, inputPaths:  %s, %s, %s) expects (%s)", testCase.inputSeparator, testCase.isIgnoreEmptyPath, testCase.isNormalize, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetCombinedPath(
				testCase.inputSeparator,
				testCase.isIgnoreEmptyPath,
				testCase.isLongPathFix,
				testCase.isNormalize,
				testCase.inputPaths1,
				testCase.inputPaths2,
				testCase.inputPaths3)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
