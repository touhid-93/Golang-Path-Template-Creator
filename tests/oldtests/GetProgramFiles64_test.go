package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var programFiles64TestCaseDataWrappers = []generalizedPathWithoutInputTestCaseDataWrapper{
	{
		operatingSystemMessage: "Unix OS",
		funcName:               "ProgramFiles64",
		expected:               "",
		operatingSystem:        ostype.Linux,
	},
	{
		operatingSystemMessage: "Windows OS",
		funcName:               "ProgramFiles64",
		expected:               "C:\\\\Program Files (x86)",
		operatingSystem:        ostype.Windows,
	},
}

func TestGetProgramFiles64_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range programFiles64TestCaseDataWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.ProgramFiles64, i)
	}
}

func TestGetProgramFiles64_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range programFiles64TestCaseDataWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		executeTestForGeneralizedPathWithoutInput(t, testCase, knowndirget.ProgramFiles64, i)
	}
}
