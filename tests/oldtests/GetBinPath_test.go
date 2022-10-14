package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var binPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "BinPath",
		expected:               "/usr/bin",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "BinPath",
		expected:               "C:\\Users\\Administrator\\bin",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetBinPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.BinPath, i)
	}
}

func TestGetBinPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range binPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.BinPath, i)
	}
}
