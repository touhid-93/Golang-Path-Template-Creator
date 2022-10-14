package mics

import "gitlab.com/evatix-go/core/ostype"

// it returns if not windows
func IsUnixCase(os ostype.Variation) bool {
	return os != ostype.Windows
}
