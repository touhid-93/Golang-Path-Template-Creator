package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var localTempPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "LocalTempPath",
		expected:               "/tmp",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "LocalTempPath",
		expected:               "C:\\Users\\Administrator\\AppData\\Roaming\\local\\temp",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetLocalTempPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.LocalTempPath, i)
	}
}

func TestGetLocalTempPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range localTempPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.LocalTempPath, i)
	}
}
