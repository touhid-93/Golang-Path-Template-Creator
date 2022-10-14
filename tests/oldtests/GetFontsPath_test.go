package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var fontsPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "FontsPath",
		expected:               "/usr/share/fonts",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "FontsPath",
		expected:               "C:\\Windows\\Fonts",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetFontsPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.FontsPath, i)
	}
}

func TestGetFontsPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range fontsPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.FontsPath, i)
	}
}
