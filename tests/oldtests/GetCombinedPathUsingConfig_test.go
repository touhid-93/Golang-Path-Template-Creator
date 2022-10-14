package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type combinedPathUsingConfigTestCaseWrapper struct {
	inputConfig                                       *pathhelpercore.PathConfig
	inputPaths1, inputPaths2, inputPaths3             string
	expected, expectedMessage, operatingSystemMessage string
	operatingSystem                                   ostype.Variation
}

var combinedPathUsingConfigTestCaseWrappers = []combinedPathUsingConfigTestCaseWrapper{
	{
		inputConfig:            &pathhelpercore.PathConfig{Separator: constants.PathSeparator},
		inputPaths1:            "something",
		inputPaths2:            "more",
		inputPaths3:            "etc",
		expected:               "something\\more\\etc",
		expectedMessage:        "something\\more\\etc",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
	},
	{
		inputConfig:            &pathhelpercore.PathConfig{Separator: constants.PathSeparator},
		inputPaths1:            "something",
		inputPaths2:            "more",
		inputPaths3:            "etc",
		expected:               "something/more/etc",
		expectedMessage:        "something/more/etc",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
	},
}

func TestGetCombinedPathUsingConfig_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range combinedPathUsingConfigTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfig] inputs (%v, %s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfig(t, testCaseMessage, testCase, i)
	}
}

func TestGetCombinedPathUsingConfig_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range combinedPathUsingConfigTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetCombinedPathUsingConfig] inputs (%v, %s, %s, %s) expects (%s)", testCase.operatingSystemMessage, testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3, testCase.expectedMessage)

		executeTestCaseForGetCombinedPathUsingConfig(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetCombinedPathUsingConfig(
	t *testing.T,
	testCaseMessage string,
	testCase combinedPathUsingConfigTestCaseWrapper,
	i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetCombinedPathUsingConfig(testCase.inputConfig, testCase.inputPaths1, testCase.inputPaths2, testCase.inputPaths3)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
