package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

type doubleBackSlashTestCaseWrapper struct {
	inputPath, inputSeparator, expected string
}

var doubleBackSlashTestCaseWrappers = []doubleBackSlashTestCaseWrapper{
	{
		inputPath:      "c:\\\\",
		inputSeparator: "\\",
		expected:       "c:\\",
	},
	{
		inputPath:      "c:\\\\",
		inputSeparator: "/",
		expected:       "c:/",
	},
	{
		inputPath:      "c:/",
		inputSeparator: "",
		expected:       "c:/",
	},
}

func TestChangeDoubleBackSlash(t *testing.T) {
	for i, testCase := range doubleBackSlashTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[ChangeDoubleBackSlash] inputs (%s, %s) expects (%s)", testCase.inputPath, testCase.inputSeparator, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := normalize.ChangeDoubleBackSlash(testCase.inputPath, testCase.inputSeparator)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
