package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileInfoWrapper(t *testing.T) {
	// Arrange
	testMessage := fmt.Sprint("[FileInfoWrapper] methods expect boolean return")

	Convey(testMessage, t, func() {
		// Act
		// actualNew := pathhelpercore.NewFileInfoWrapper("") // todo
		// actualHasError := actualNew.HasError()
		// actualPathExists := actualNew.IsPathExists()

		// Assert
		// So(actualHasError, ShouldBeTrue)
		// So(actualPathExists, ShouldBeFalse)
	})
}
