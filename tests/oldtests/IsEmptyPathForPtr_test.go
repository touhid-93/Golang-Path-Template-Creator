package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

var path = "something/whatever"

func TestIsEmptyPathForPtr(t *testing.T) {
	// Arrange
	testCaseMessage := fmt.Sprintf("[EmptyPtr] inputs (something/whatever) expects (false)")

	Convey(testCaseMessage, t, func() {
		// Act
		actual := ispath.EmptyPtr(&path)

		// Assert
		So(actual, ShouldBeFalse)
	})
}
