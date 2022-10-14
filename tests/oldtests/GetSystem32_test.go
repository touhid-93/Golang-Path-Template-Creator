package oldtests

import (
	"fmt"
	"testing"

	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/pathhelper/knowndirget"

	. "github.com/smartystreets/goconvey/convey"
)

var expectedSystem32 = "C:\\Windows\\System32"

func TestGetSystem32(t *testing.T) {
	coretests.SkipOnUnix(t)

	// Arrange
	testCaseMessage := fmt.Sprintf("[System32] expects (%s)", expectedSystem32)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := knowndirget.System32()

		// Assert
		So(actual, ShouldEqual, expectedSystem32)
	})
}
