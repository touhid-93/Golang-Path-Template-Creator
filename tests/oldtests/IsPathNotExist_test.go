package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

type pathNotExistTestCaseWrapper struct {
	input, expectedMessage, operatingSystemMessage string
	expected                                       bool
	operatingSystem                                ostype.Variation
}

var pathNotExistTestCaseWrappers = []pathNotExistTestCaseWrapper{
	{
		input:                  "",
		expected:               true,
		expectedMessage:        "true",
		operatingSystemMessage: "any OS",
	},
	{
		input:                  "C:\\Users",
		expected:               false,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Windows",
		operatingSystem:        ostype.Windows,
	},
	{
		input:                  "home/user",
		expected:               true,
		expectedMessage:        "false",
		operatingSystemMessage: "OS is Unix",
		operatingSystem:        ostype.Linux,
	},
}

func TestIsPathNotExist_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathNotExistTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathNotExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathNotExist(t, testCaseMessage, testCase, i)
	}
}

func TestIsPathNotExist_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathNotExistTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathNotExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestCaseForIsPathNotExist(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForIsPathNotExist(
	t *testing.T, testCaseMessage string, testCase pathNotExistTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := ispath.NotExists(testCase.input)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
