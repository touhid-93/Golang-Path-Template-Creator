package oldtests

//
// type hasPathIssuesTestCaseWrapper struct {
// 	input, expectedMessage string
// 	expected               bool
// }
//
// var hasPathIssuesTestCaseWrappers = []hasPathIssuesTestCaseWrapper{
// 	{
// 		input:           "file:///C:/",
// 		expected:        true,
// 		expectedMessage: "ShouldBeTrue",
// 	},
// 	{
// 		input:           "//C:\\win",
// 		expected:        true,
// 		expectedMessage: "ShouldBeTrue",
// 	},
// 	{
// 		input:           "file:///C:\\win\\users",
// 		expected:        true,
// 		expectedMessage: "ShouldBeTrue",
// 	},
// 	{
// 		input:           "C:/windows/",
// 		expected:        false,
// 		expectedMessage: "ShouldBeFalse",
// 	},
// }
//
// func TestHasPathIssues(t *testing.T) {
// 	for _, testCase := range hasPathIssuesTestCaseWrappers {
// 		// Arrange
// 		testCaseMessage := fmt.Sprintf("[HasPathIssues] inputs (%s) expects (%s)", testCase.input, testCase.expectedMessage)
//
// 		Convey(testCaseMessage, t, func() {
// 			// Act
// 			actual := pathhelper.HasPathIssues(testCase.input)
//
// 			// Arrange
// 			// HasPathIssues returns true if argument has prefix or double forward and backward slashes  or both
// 			So(actual, ShouldNotBeNil)
// 			So(actual, ShouldNotBeEmpty)
// 			So(actual, ShouldEqual, testCase.expected)
// 		})
// 	}
// }
