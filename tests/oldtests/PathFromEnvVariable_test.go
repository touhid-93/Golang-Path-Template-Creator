package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

type pathFromEnvVariableTestCaseWrapper struct {
	input, expected, expectedMessage string
}

var pathFromEnvVariableTestCaseWrappers = []pathFromEnvVariableTestCaseWrapper{
	{
		input:           "",
		expected:        "",
		expectedMessage: "",
	},
	{
		input:           "$home $sys hello world $what",
		expected:        "",
		expectedMessage: "",
	},
	{
		input:           "$ComSpec hello $no",
		expected:        "",
		expectedMessage: "",
	},
}

func TestPathFromEnvVariable(t *testing.T) {
	for i, testCase := range pathFromEnvVariableTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[ExpandVariables] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := expandpath.ExpandVariables(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() { // todo check equality
				if ispath.Empty(testCase.input) {
					So(actual, ShouldBeEmpty)
				}

				if !ispath.Empty(testCase.input) {
					So(actual, ShouldNotBeEmpty)
				}
			})
		})

	}
}
