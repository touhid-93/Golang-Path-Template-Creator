package oldtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/osconsts"
)

func TestIsWindows(t *testing.T) {
	// Arrange
	Convey("function should return true on windows OS", t, func() {
		SkipOnUnix(t)

		// Act
		actual := osconsts.IsWindows

		// Assert
		So(actual, ShouldBeTrue)
	})

	Convey("function should return false on unix OS", t, func() {
		SkipOnWindows(t)

		// Act
		actual := osconsts.IsWindows

		// Assert
		So(actual, ShouldBeFalse)
	})
}
