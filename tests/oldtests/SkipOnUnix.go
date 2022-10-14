package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/core/coretests"
)

// SkipOnUnix Skip on Unix
func SkipOnUnix(t *testing.T) {
	coretests.SkipOnUnix(t)
}
