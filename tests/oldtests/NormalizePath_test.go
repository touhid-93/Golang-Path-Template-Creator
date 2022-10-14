package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type normalizePathTestCaseWrapper struct {
	input, expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                          ostype.Variation
}

var normalizePathTestCaseWrappers = []normalizePathTestCaseWrapper{
	{
		input:                  "",
		expected:               "",
		expectedMessage:        "empty return",
		operatingSystemMessage: "Any OS",
	},
	{
		input:                  "c:/windows/system32/etc",
		expected:               "c:/windows/system32/etc",
		expectedMessage:        "non-empty return (c:/windows/system32/etc)",
		operatingSystemMessage: "Any OS",
	},
	{
		input:                  "c:\\windows//system32\\//etc",
		expected:               "c:\\windows\\system32\\etc",
		expectedMessage:        "non-empty return (c:\\windows\\system32\\etc)",
		operatingSystemMessage: "OS is windows",
		operatingSystem:        ostype.Windows,
	},
	{
		input:                  "c:\\windows//system32\\//etc",
		expected:               "c:/windows/system32/etc",
		expectedMessage:        "non-empty return (c:/windows/system32/etc)",
		operatingSystemMessage: "OS other than windows",
		operatingSystem:        ostype.Linux,
	},
}

func TestNormalizePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range normalizePathTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[NormalizePath] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestNormalizePath(t, testCaseMessage, testCase, i)
	}
}

func TestNormalizePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range normalizePathTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s)[IsPathExist] inputs (%s) expects (%s)", testCase.operatingSystemMessage, testCase.input, testCase.expectedMessage)

		executeTestNormalizePath(t, testCaseMessage, testCase, i)
	}
}

func executeTestNormalizePath(t *testing.T, testCaseMessage string, testCase normalizePathTestCaseWrapper, i int) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := normalize.Path(testCase.input)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeEmpty)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
