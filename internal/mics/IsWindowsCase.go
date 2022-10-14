package mics

import "gitlab.com/evatix-go/core/ostype"

func IsWindowsCase(os ostype.Variation) bool {
	return os == ostype.Windows
}
