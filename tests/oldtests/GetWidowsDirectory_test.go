package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

type windowsDirectoryTestCaseWrapper struct {
	expected, operatingSystemMessage string
	operatingSystem                  ostype.Variation
}

var windowsDirectoryTestCaseWrappers = []windowsDirectoryTestCaseWrapper{
	{
		expected:               "C:\\Windows",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
	},
	{
		expected:               "",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
	},
}

func TestGetWidowsDirectory_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range windowsDirectoryTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[WidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase, i)
	}
}

func TestGetWidowsDirectory_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range windowsDirectoryTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[WidowsDirectory] expects (%s)", testCase.operatingSystemMessage, testCase.expected)

		executeTestCaseForGetWidowsDirectory(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetWidowsDirectory(
	t *testing.T,
	testCaseMessage string,
	testCase windowsDirectoryTestCaseWrapper,
	i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := knowndirget.WidowsDirectory()

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
