package pathcompilertests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/pathhelper/pathcompiler"
)

func Test_UnixOsSpecificPathSelect_OnWindows(t *testing.T) {
	for _, testCase := range pathCompilerTestCases {
		if !osmixtype.IsCurrentOsTypesContains(testCase.RunsIn...) {
			continue
		}

		// Arrange
		expectedMap := testCase.ExpectedAsDynamicMap()
		selectedSpecificPaths := pathcompiler.
			DefaultApp.ByEnvFlagOs(
			testCase.IsTestEnv,
			testCase.SelectBy)

		// Act
		jsonResult := selectedSpecificPaths.JsonPtr()
		fieldsMap, parsingErr := DeserializedFieldsToMap(
			jsonResult)
		errcore.MustBeEmpty(parsingErr)

		var dynamicMap enumimpl.DynamicMap = fieldsMap
		diffMessage := dynamicMap.LogShouldDiffMessage(
			true,
			testCase.Title,
			expectedMap)

		isValid := diffMessage == ""

		// Assert
		convey.Convey(testCase.Title, t, func() {
			convey.So(selectedSpecificPaths, convey.ShouldNotBeNil)
			convey.So(diffMessage, convey.ShouldBeEmpty)
			convey.So(isValid, convey.ShouldBeTrue)
			convey.So(selectedSpecificPaths.Name, convey.ShouldEqual, testCase.NameAssert)
			convey.So(selectedSpecificPaths.Description, convey.ShouldEqual, testCase.DescriptionAssert)
		})
	}
}
