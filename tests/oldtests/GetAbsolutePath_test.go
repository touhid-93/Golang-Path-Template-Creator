package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type absolutePathTestCaseWrapper struct {
	basePath, inputRelativePath, expected, expectedMessage, operatingSystemMessage string
	isLongPathFix, isNormalize                                                     bool
	operatingSystem                                                                ostype.Variation
}

var absolutePathTestCaseWrappers = []absolutePathTestCaseWrapper{
	// todo catch panic
	// {
	//	basePath:               "",
	//	inputRelativePath:      "",
	//	expected:               "",
	//	expectedMessage:        "empty return",
	//	operatingSystemMessage: "Any OS",
	//	operatingSystem:        knowndir.Any,
	//	isNormalize:            true,
	//	isLongPathFix:          true,
	// },
	{
		basePath:               "c:\\Windows\\//",
		inputRelativePath:      "\\whatever",
		expected:               `\\?\c:\Windows\whatever`,
		expectedMessage:        "non-empty return of (\\\\?\\c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "c:\\\\Windows///",
		inputRelativePath:      "whatever",
		expected:               `\\?\c:\Windows\whatever`,
		expectedMessage:        "non-empty return of (\\\\?\\c:\\Windows\\whatever)",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "/home/\\//your_user_name/my_script/",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script/whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
		isNormalize:            true,
		isLongPathFix:          true,
	},
	{
		basePath:               "/home/your_user_name/my_script",
		inputRelativePath:      "/whatever",
		expected:               "/home/your_user_name/my_script//whatever",
		expectedMessage:        "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystemMessage: "Unix OS",
		isNormalize:            false,
		isLongPathFix:          true,
		operatingSystem:        ostype.Linux,
	},
}

func TestGetAbsolutePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for testCaseIndex, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf(
			"(%q) [GetAbsolutePath] inputs (%q, %q) expects (%q)",
			testCase.operatingSystemMessage,
			testCase.basePath,
			testCase.inputRelativePath,
			testCase.expectedMessage)

		executeTestForGetAbsolutePath(
			t,
			testCaseIndex,
			testCaseMessage,
			testCase,
		)
	}
}

func TestGetAbsolutePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for caseIndex, testCase := range absolutePathTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.basePath, testCase.inputRelativePath, testCase.expectedMessage)

		executeTestForGetAbsolutePath(
			t,
			caseIndex,
			testCaseMessage,
			testCase)
	}
}

func executeTestForGetAbsolutePath(
	t *testing.T,
	caseIndex int,
	testCaseMessage string,
	testCase absolutePathTestCaseWrapper,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetAbsolutePath(
			testCase.basePath,
			testCase.inputRelativePath,
			testCase.isLongPathFix,
			testCase.isNormalize)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, caseIndex), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
