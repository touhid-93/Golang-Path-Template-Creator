package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type absoluteFromExecutableDirectoryPathTestCaseWrapper struct {
	inputRelativePath, expected, expectedMessage, operatingSystemMessage string
	isLongPathFix, isNormalize                                           bool
	operatingSystem                                                      ostype.Variation
}

// todo
var absoluteFromExecutableDirectoryPathTestCaseWrappers = []absoluteFromExecutableDirectoryPathTestCaseWrapper{
	{
		inputRelativePath:      "\\Users",
		expected:               "C:\\Users",
		expectedMessage:        "C:\\Users",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		inputRelativePath:      "\\Users",
		expected:               "C:\\Users",
		expectedMessage:        "C:\\Users",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		inputRelativePath:      "//\\home//",
		expected:               "/home/",
		expectedMessage:        "/home/",
		operatingSystemMessage: "Linux OS",
		operatingSystem:        ostype.Linux,
		isNormalize:            true,
		isLongPathFix:          true,
	},
}

func TestGetAbsoluteFromExecutableDirectoryPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range absoluteFromExecutableDirectoryPathTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsoluteFromExecutableDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t, testCaseMessage, testCase, i)
	}
}

func TestGetAbsoluteFromExecutableDirectoryPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range absoluteFromExecutableDirectoryPathTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsoluteFromExecutableDirectoryPath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetAbsoluteFromExecutableDirectoryPath(
	t *testing.T, testCaseMessage string, testCase absoluteFromExecutableDirectoryPathTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsoluteFromExecutableDirectoryPath(
			testCase.inputRelativePath,
			testCase.isLongPathFix,
			testCase.isNormalize)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldContainSubstring, testCase.expected)
		})
	})
}
