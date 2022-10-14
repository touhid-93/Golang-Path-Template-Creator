package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/ostype"
)

type generalizedPathWithoutInputTestCaseDataWrapper struct {
	expected, operatingSystemMessage, funcName string
	operatingSystem                            ostype.Variation
}

func executeTestForGeneralizedPathWithoutInput(
	t *testing.T,
	testCase generalizedPathWithoutInputTestCaseDataWrapper,
	funcCall func() string,
	i int,
) {
	testCaseMessage := fmt.Sprintf("(%s) [%s] expects (%s)", testCase.operatingSystemMessage, testCase.funcName, testCase.expected)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := funcCall()

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
