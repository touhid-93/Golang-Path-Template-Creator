package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type combinePathsWithTestCaseDataWrapper struct {
	inputPath1, inputPath2, inputPath3 string
	operatingSystem                    ostype.Variation
	expected, operatingSystemMessage   string
}

var combinePathsWithTestCaseWrappers = []combinePathsWithTestCaseDataWrapper{
	{
		inputPath1:             "something",
		inputPath2:             "somethingElse",
		inputPath3:             "otherThings",
		operatingSystem:        ostype.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               "something\\somethingElse\\otherThings",
	},
	{
		inputPath1:             "",
		inputPath2:             "",
		inputPath3:             "",
		operatingSystem:        ostype.Windows,
		operatingSystemMessage: "Windows OS",
		expected:               "\\",
	},
	{
		inputPath1:             "something",
		inputPath2:             "somethingElse",
		inputPath3:             "otherThings",
		operatingSystem:        ostype.Linux,
		operatingSystemMessage: "Windows OS",
		expected:               "something/somethingElse/otherThings",
	},
}

func TestGetCombinePathsWith_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range combinePathsWithTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinePathsWith] inputs (%s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath1, testCase.inputPath2, testCase.inputPath3, testCase.expected)

		executeTestForGetCombinePathsWith(t, testCaseMessage, testCase, i)
	}
}

func TestGetCombinePathsWith_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range combinePathsWithTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinePathsWith] inputs (%s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPath1, testCase.inputPath2, testCase.inputPath3, testCase.expected)

		executeTestForGetCombinePathsWith(t, testCaseMessage, testCase, i)
	}
}

func executeTestForGetCombinePathsWith(
	t *testing.T, testCaseMessage string, testCase combinePathsWithTestCaseDataWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetCombinePathsWith(
			testCase.inputPath1,
			testCase.inputPath2,
			testCase.inputPath3)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
