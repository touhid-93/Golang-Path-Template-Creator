package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
)

type pathFromUriTestCaseWrapper struct {
	givenPath, expected, expectedMessage, operatingSystemMessage string
	isNormalize                                                  bool
	operatingSystem                                              ostype.Variation
}

var pathFromUriTestCaseWrappers = []pathFromUriTestCaseWrapper{
	{
		givenPath:              "file://c:/windows/users/etc/more",
		isNormalize:            true,
		expected:               "c:\\windows\\users\\etc\\more",
		expectedMessage:        "c:\\windows\\users\\etc\\more",
		operatingSystemMessage: "Windows OS",
		operatingSystem:        ostype.Windows,
	},
	{
		givenPath:              "c:\\windows\\users\\etc\\more",
		isNormalize:            true,
		expected:               "c:/windows/users/etc/more",
		expectedMessage:        "c:/windows/users/etc/more",
		operatingSystemMessage: "Unix OS",
		operatingSystem:        ostype.Linux,
	},
}

func TestGetPathFromUri_Windows(t *testing.T) {
	SkipOnUnix(t)

	for i, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if mics.IsUnixCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.givenPath, testCase.isNormalize, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase, i)
	}
}

func TestGetPathFromUri_Unix(t *testing.T) {
	SkipOnWindows(t)

	for i, testCase := range pathFromUriTestCaseWrappers {
		// Arrange
		if mics.IsWindowsCase(testCase.operatingSystem) {
			continue
		}

		testCaseMessage := fmt.Sprintf("(%s) [GetPathFromURI] inputs (%s, %v) expects (%s)", testCase.operatingSystemMessage, testCase.givenPath, testCase.isNormalize, testCase.expectedMessage)

		executeTestCaseForGetPathFromUri(t, testCaseMessage, testCase, i)
	}
}

func executeTestCaseForGetPathFromUri(
	t *testing.T, testCaseMessage string, testCase pathFromUriTestCaseWrapper, i int,
) {
	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelper.GetPathFromUri(testCase.givenPath, testCase.isNormalize)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
