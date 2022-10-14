package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var systemPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GetSystemPath",
		expected:               "/etc/systemd/system",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GetSystemPath",
		expected:               "C:\\Windows",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetSystemPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetSystemPath, i)
	}
}

func TestGetSystemPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range systemPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GetSystemPath, i)
	}
}
