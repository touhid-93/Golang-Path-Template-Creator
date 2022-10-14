package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var gitGlobalPathTestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "GitGlobal",
		expected:               "XDG_CONFIG_HOME/git/config",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "GitGlobal",
		expected:               "C:\\Users\\Administrator\\.gitconfig",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetGitGlobal_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range gitGlobalPathTestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GitGlobal, i)
	}
}

func TestGetGitGlobal_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range gitGlobalPathTestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.GitGlobal, i)
	}
}
