package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

var pathToCheck = "something/whatever"

func TestIsEmptyPath(t *testing.T) {
	// Arrange
	testCaseMessage := fmt.Sprintf("[Empty] inputs (something/whatever) expects (false)")

	Convey(testCaseMessage, t, func() {
		// Act
		actual := ispath.Empty(pathToCheck)

		// Assert
		So(actual, ShouldBeFalse)
	})
}
