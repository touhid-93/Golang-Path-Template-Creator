package pathhelper

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/core/ostype"
	"gitlab.com/evatix-go/pathhelper/tests/oldtests"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type combinedPathUsingConfigInternalTestCaseWrapper struct {
	inputPathConfig                                             *pathhelpercore.PathConfig
	inputPaths                                                  []string
	expected, expectedMessage, operatingSystemMessage, funcName string
	operatingSystem                                             ostype.Variation
}

var combinedPathUsingConfigInternalTestCaseWrappers = []combinedPathUsingConfigInternalTestCaseWrapper{
	{
		inputPathConfig:        &pathhelpercore.PathConfig{Separator: "\\"},
		inputPaths:             []string{"C:\\", "somethingelse\\", "etc"},
		expected:               "C:\\somethingelse\\etc",
		expectedMessage:        "C:\\somethingelse\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		funcName:               "GetCombinedPathUsingConfigInternal",
	},
	{
		inputPathConfig:        nil,
		inputPaths:             []string{"C:", "somethingelse", "etc"},
		expected:               "C:\\somethingelse\\etc",
		expectedMessage:        "C:\\somethingelse\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
		funcName:               "GetCombinedPathUsingConfigInternal",
	},
	{
		inputPathConfig:        &pathhelpercore.PathConfig{IsNormalize: true},
		inputPaths:             []string{"home\\", "\\somethingelse\\", "\\etc"},
		expected:               "home/somethingelse/etc",
		expectedMessage:        "home/somethingelse/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
		funcName:               "GetCombinedPathUsingConfigInternal",
	},
}

func TestGetCombinedPathUsingConfigInternal_Windows(t *testing.T) {
	if !osconsts.IsWindows {
		t.Skip("Windows tests ignored in Unix.")
	}

	for i, testCase := range combinedPathUsingConfigInternalTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfigInternal]  inputs (pathConfig: %v, paths: %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPathConfig, testCase.inputPaths, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfigInternal(t, testCaseMessage, testCase, i)
	}
}

func TestGetCombinedPathUsingConfigInternal_Unix(t *testing.T) {
	if osconsts.IsWindows {
		t.Skip("Unix tests ignored in Windows.")
	}

	for i, testCase := range combinedPathUsingConfigInternalTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfigInternal]  inputs (pathConfig: %v, paths: %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputPathConfig, testCase.inputPaths, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfigInternal(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetCombinedPathUsingConfigInternal(
	t *testing.T,
	testCaseMessage string,
	testCase combinedPathUsingConfigInternalTestCaseWrapper,
	i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := getCombinedPathUsingConfigInternal(testCase.inputPathConfig, testCase.inputPaths)

		// Assert
		Convey(oldtests.GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
