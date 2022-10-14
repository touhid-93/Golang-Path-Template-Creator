package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var userVideosPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserVideosPath",
		expected:               homePath + "/Videos",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserVideosPath",
		expected:               "C:\\Users\\Administrator\\Videos",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserVideosPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userVideosPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserVideosPath, i)
	}
}

func TestGetUserVideosPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userVideosPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserVideosPath, i)
	}
}
