package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/tests/oldtests"
)

type letterTestCaseWrapper struct {
	input, expectedMessage string
	expected               bool
}

var letterTestCaseWrappers = []letterTestCaseWrapper{
	{
		input:           "something",
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           "something/somethingelse",
		expected:        false,
		expectedMessage: "false",
	},
	{
		input:           "-something/21",
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsLetter(t *testing.T) {
	for i, testCase := range letterTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[isLetter] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := isLetter(testCase.input)

			// Assert
			Convey(oldtests.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeEmpty)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
