package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

type directoryPathTestCaseWrapper struct {
	operatingSystem                                ostype.Variation
	input, operatingSystemMessage, expectedMessage string
	expected                                       bool
}

var directoryPathTestCaseWrappers = []directoryPathTestCaseWrapper{
	{
		input:                  "C:\\users",
		operatingSystem:        ostype.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               true,
		expectedMessage:        "true",
	},
	{
		input:                  "/home",
		operatingSystem:        ostype.Linux,
		operatingSystemMessage: "Unix os",
		expected:               true,
		expectedMessage:        "true",
	},
}

func TestIsDirectoryPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range directoryPathTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [IsDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsDirectoryPath(t, testCaseMessage, testCase, i)
	}
}

func TestIsDirectoryPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range directoryPathTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [IsDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsDirectoryPath(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForIsDirectoryPath(
	t *testing.T, testCaseMessage string, testCase directoryPathTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := ispath.Directory(testCase.input)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
