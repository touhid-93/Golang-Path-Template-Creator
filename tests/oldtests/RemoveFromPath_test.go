package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

type removeFromPathTestCaseWrapper struct {
	inputPath, expected, expectedMessage, operatingSystemMessage string
	inputBool                                                    bool
	operatingSystem                                              ostype.Variation
}

var removingArray = []string{"/"}

var removeFromPathTestCaseWrappers = []removeFromPathTestCaseWrapper{
	{
		inputPath:              "",
		expected:               "",
		expectedMessage:        "empty return",
		operatingSystemMessage: "Any OS",
		operatingSystem:        ostype.Windows,
	},
	{
		inputPath:              "c:\\win\\etc",
		inputBool:              true,
		expected:               "c:\\win\\etc",
		expectedMessage:        "c:\\win\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
	},
	{
		inputPath:              "c:\\\\win\\\\etc",
		inputBool:              true,
		expected:               "c:/win/etc",
		expectedMessage:        "c:/win/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
	},
}

func TestRemoveFromPath_windows(t *testing.T) {
	if !osconsts.IsWindows {
		t.Skip("Windows tests ignored in Unix.")
	}

	for i, testCase := range removeFromPathTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [RemoveFromPath] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForRemoveFromPath(t, testCaseMessage, testCase, i)
	}
}

func TestRemoveFromPath_unix(t *testing.T) {
	if osconsts.IsWindows {
		t.Skip("Windows tests ignored in Unix.")
	}

	for i, testCase := range removeFromPathTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [RemoveFromPath] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath, testCase.inputBool, testCase.expectedMessage)

		executeTestCaseForRemoveFromPath(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForRemoveFromPath(
	t *testing.T, testCaseMessage string, testCase removeFromPathTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.RemoveFromPath(testCase.inputPath, &removingArray, testCase.inputBool)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			if ispath.Empty(testCase.inputPath) {
				So(actual, ShouldBeEmpty)
			}

			if !ispath.Empty(testCase.inputPath) {
				So(actual, ShouldNotBeEmpty)
				So(actual, ShouldEqual, testCase.expected)
			}
		})
	})
}
