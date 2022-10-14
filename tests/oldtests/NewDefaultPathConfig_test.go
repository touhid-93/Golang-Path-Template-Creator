package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var expected = pathhelpercore.PathConfig{
	IsNormalize:                  false,
	IsIgnoreEmptyPath:            false,
	IsExpandEnvironmentVariables: true,
}

func TestNewDefaultPathConfig(t *testing.T) {
	// Arrange
	testCaseMessage := fmt.Sprintf("[NewDefaultPathConfig] inputs () expects (type *PathConfig)")

	Convey(testCaseMessage, t, func() {
		// Act
		actual := pathhelpercore.NewDefaultPathConfigOrExisting(nil)

		// Assert
		So(actual.IsNormalize, ShouldBeTrue)
		So(actual.IsIgnoreEmptyPath, ShouldBeFalse)
		So(actual.IsExpandEnvironmentVariables, ShouldBeTrue)
	})
}
