package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
)

type isEmptyArrayTestCaseWrapper struct {
	input           []string
	expected        bool
	expectedMessage string
}

var isEmptyArrayTestCaseWrappers = []isEmptyArrayTestCaseWrapper{
	{
		input:           []string{},
		expected:        true,
		expectedMessage: "true",
	},
	{
		input:           []string{"hello", "world"},
		expected:        false,
		expectedMessage: "false",
	},
}

func TestIsEmptyArray(t *testing.T) {
	for i, testCase := range isEmptyArrayTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[EmptyArray] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := ispathinternal.EmptyArray(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
