package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

type changeSeparatorTestCaseWrapper struct {
	inputPath, inputCurrentSeparator, inputChangeSeparator, expected string
}

var changeSeparatorTestCaseWrappers = []changeSeparatorTestCaseWrapper{
	{
		inputPath:             "c:\\windows/example",
		inputCurrentSeparator: "\\",
		inputChangeSeparator:  "/",
		expected:              "c:/windows/example",
	},
	{
		inputPath:             "home/user/example",
		inputCurrentSeparator: "\\",
		inputChangeSeparator:  "/",
		expected:              "home/user/example",
	},
	{
		inputPath:             "home/user/example",
		inputCurrentSeparator: "/",
		inputChangeSeparator:  "\\",
		expected:              "home\\user\\example",
	},
}

func TestChangeSeparator(t *testing.T) {
	for i, testCase := range changeSeparatorTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[ChangeSeparator] inputs (%s, %s, %s) expects (%s)", testCase.inputPath, testCase.inputCurrentSeparator, testCase.inputChangeSeparator, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := normalize.ChangeSeparator(testCase.inputPath, testCase.inputCurrentSeparator, testCase.inputChangeSeparator)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
