package pathhelpercore

import (
	"gitlab.com/evatix-go/core/constants"
)

// PathConfig By default:
// Separator = os.PathSeparator
// IsNormalize = true
// IsIgnoreEmptyPath = false
// IsExpandEnvironmentVariables = true
type PathConfig struct {
	Separator                    string
	IsNormalize                  bool
	IsIgnoreEmptyPath            bool
	IsLongPathFix                bool
	IsExpandEnvironmentVariables bool
}

// By default:
// Separator = os.PathSeparator
// IsNormalize = true
// IsIgnoreEmptyPath = false
// IsLongPathFix = true
// IsExpandEnvironmentVariables: true
func NewDefaultPathConfigOrExisting(existingConfig *PathConfig) *PathConfig {
	if existingConfig != nil {
		return existingConfig
	}

	return &PathConfig{
		Separator:                    constants.PathSeparator,
		IsNormalize:                  true,
		IsIgnoreEmptyPath:            false,
		IsLongPathFix:                true,
		IsExpandEnvironmentVariables: true,
	}
}
