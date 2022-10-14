package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var userPicturesPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "UserPicturesPath",
		expected:               homePath + "/Pictures",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "UserPicturesPath",
		expected:               "C:\\Users\\Administrator\\Pictures",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetUserPicturesPath_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserPicturesPath, i)
	}
}

func TestGetUserPicturesPath_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range userPicturesPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.UserPicturesPath, i)
	}
}
