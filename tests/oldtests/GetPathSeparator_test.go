package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type pathSeparatorTestCaseWrapper struct {
	operatingSystem                                   ostype.Variation
	operatingSystemMessage, expected, expectedMessage string
}

var pathSeparatorTestCaseWrappers = []pathSeparatorTestCaseWrapper{
	{
		operatingSystem:        ostype.Windows,
		operatingSystemMessage: "Os is windows",
		expected:               constants.BackSlash,
		expectedMessage:        constants.BackSlash,
	},
	{
		operatingSystem:        ostype.Linux,
		operatingSystemMessage: "Unix os",
		expected:               constants.ForwardSlash,
		expectedMessage:        constants.ForwardSlash,
	},
}

func TestGetPathSeparator_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathSeparatorTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystemMessage, testCase.expectedMessage)

		executeTestCaseForGetPathSeparator(t, testCaseMessage, testCase, i)
	}
}

func TestGetPathSeparator_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathSeparatorTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathSeparator] inputs () expects (%s)", testCase.operatingSystemMessage, testCase.expectedMessage)

		executeTestCaseForGetPathSeparator(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetPathSeparator(
	t *testing.T, testCaseMessage string, testCase pathSeparatorTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := constants.PathSeparator

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
