package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var hostFilePathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "HostFilePath",
		expected:               "/etc/hosts",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "HostFilePath",
		expected:               "C:\\Windows\\System32\\drivers\\etc\\hosts",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetHostFilePath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.HostFilePath, i)
	}
}

func TestGetHostFilePath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range hostFilePathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.HostFilePath, i)
	}
}
