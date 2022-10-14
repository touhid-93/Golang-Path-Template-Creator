package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func AddPathExtensionOnRequired(location, dotExt string) string {
	return splitinternal.AddPathExtensionOnRequired(location, dotExt)
}
