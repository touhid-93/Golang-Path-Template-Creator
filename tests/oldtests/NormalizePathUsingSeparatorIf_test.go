package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

type normalizePathUsingSeparatorIfTestCaseWrapper struct {
	inputString, inputSeparator, expected  string
	isForced, isLongPathFixed, isNormalize bool
}

var normalizePathUsingSeparatorIfTestCaseWrappers = []normalizePathUsingSeparatorIfTestCaseWrapper{
	{
		inputString:     "file://c:\\windows//system32\\//etc",
		inputSeparator:  "/",
		isForced:        false,
		isLongPathFixed: true,
		isNormalize:     true,
		expected:        "c:/windows/system32/etc",
	},
	{
		inputString:    "c:\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne",
		inputSeparator: "/",
		isForced:       false,

		isLongPathFixed: true,
		isNormalize:     false,
		expected:        "\\\\?\\c:\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne",
	},
	{
		inputString:    "",
		inputSeparator: "/",
		isForced:       false,

		isLongPathFixed: true,
		isNormalize:     true,
		expected:        "",
	},
}

func TestNormalizePathUsingSeparatorIf(t *testing.T) {
	for i, testCase := range normalizePathUsingSeparatorIfTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[NormalizePathUsingSeparatorIf] inputs (%s, %s, %v, %v) expects (%s)", testCase.inputString, testCase.inputSeparator, testCase.isLongPathFixed, testCase.isNormalize, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := normalize.PathUsingSeparatorIf(
				testCase.isLongPathFixed,

				testCase.isLongPathFixed,
				testCase.isNormalize,
				testCase.inputSeparator,
				testCase.inputString)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
