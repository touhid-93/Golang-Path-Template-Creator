package oldtests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/pathhelper/knowndirget"

	. "github.com/smartystreets/goconvey/convey"
)

var expectedWindowsRoot = "C:\\"

func TestGetWindowsRoot(t *testing.T) {
	SkipOnUnix(t)

	// Arrange
	testCaseMessage := fmt.Sprintf("[WindowsRoot] expects (%s)", expectedWindowsRoot)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := knowndirget.WindowsRoot()

		// Assert
		So(actual, ShouldEqual, expectedWindowsRoot)
	})
}
