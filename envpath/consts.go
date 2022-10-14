package envpath

import (
	"gitlab.com/evatix-go/core/constants"
)

const (
	unixEnvPathSplitter    = constants.Colon
	windowsEnvPathSplitter = constants.SemiColon
	pathEqual              = "PATH="
)
