package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
)

type isEmptyArrayPtrTestCaseWrapper struct {
	input           []*string
	expected        bool
	expectedMessage string
}

var (
	hello = "hello"
	world = "world"
)
var isEmptyArrayPtrTestCaseWrappers = []isEmptyArrayPtrTestCaseWrapper{
	{
		input:           []*string{},
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           []*string{&hello, &world},
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsEmptyArrayPtr(t *testing.T) {
	for i, testCase := range isEmptyArrayPtrTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[EmptyArrayPtr] inputs (%v) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := ispathinternal.EmptyArrayPtr(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
