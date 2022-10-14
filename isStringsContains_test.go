package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/tests/oldtests"
)

type stringsContainsTestCaseWrapper struct {
	arrayWhereToSearch []string
	findingItem        string
	expected           bool
	expectedMessage    string
}

var nilArray []string
var emptyArray []string
var array = []string{"hello", "world"}

var stringsContainTestCaseWrappers = []stringsContainsTestCaseWrapper{
	{
		arrayWhereToSearch: emptyArray,
		findingItem:        "",
		expected:           false,
		expectedMessage:    "false",
	},
	{
		arrayWhereToSearch: nil,
		findingItem:        "",
		expected:           false,
		expectedMessage:    "false",
	},
	{
		arrayWhereToSearch: nilArray,
		findingItem:        "",
		expected:           false,
		expectedMessage:    "false",
	},
	{
		arrayWhereToSearch: array,
		findingItem:        "hello",
		expected:           true,
		expectedMessage:    "true",
	},
	{
		arrayWhereToSearch: array,
		findingItem:        "ello",
		expected:           false,
		expectedMessage:    "false",
	},
}

func TestIsStringsContains(t *testing.T) {
	for i, testCase := range stringsContainTestCaseWrappers {
		// Arrange
		testcaseMessage := fmt.Sprintf("[IsStringsContains] inputs (%s, %s) expects (%s)\"", testCase.arrayWhereToSearch, testCase.findingItem, testCase.expectedMessage)

		Convey(testcaseMessage, t, func() {
			// Act
			actual := isStringsContains(testCase.arrayWhereToSearch, testCase.findingItem)

			// Assert
			Convey(oldtests.GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeEmpty)
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
