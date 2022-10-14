package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/tests/oldtests"

	"gitlab.com/evatix-go/pathhelper/urischemes"
)

type whichPrefixTestCaseWrapper struct {
	input, expectedMessage string
	expected               urischemes.Type
}

var whichPrefixTestCaseWrappers = []whichPrefixTestCaseWrapper{
	{
		input:           "",
		expected:        urischemes.UriUnknown,
		expectedMessage: "UriUnknown",
	},
	{
		input:           "file:///",
		expected:        urischemes.UriSchemePrefixStandard,
		expectedMessage: "UriSchemePrefixStandard",
	},
	{
		input:           "file://",
		expected:        urischemes.UriSchemePrefixTwoSlashes,
		expectedMessage: "UriSchemePrefixTwoSlashes",
	},
}

func TestWhichPrefix(t *testing.T) {
	for i, testCase := range whichPrefixTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[whichPrefix] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := whichPrefix(testCase.input)

			// Assert
			Convey(oldtests.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeEmpty)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
