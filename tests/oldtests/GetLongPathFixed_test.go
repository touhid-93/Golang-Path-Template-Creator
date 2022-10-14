package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type longPathFixedTestCaseDataWrapper struct {
	input, expected string
}

var longPathFixedTestCaseDataWrappers = []longPathFixedTestCaseDataWrapper{
	{
		input:    "",
		expected: "",
	},
	{
		input:    "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
	{
		input:    "\\\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
		expected: "\\\\?\\UNC\\sample\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\somethingElseOne\\",
	},
}

func TestGetLongPathFixed(t *testing.T) {
	options := normalize.Options{
		IsNormalize:        false,
		IsLongPathFix:      true,
		IsForceLongPathFix: true,
	}

	for i, testCase := range longPathFixedTestCaseDataWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetLongPathFixed] inputs (%s) expects (%s)", testCase.input, testCase.expected)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := options.FixPath(testCase.input)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
