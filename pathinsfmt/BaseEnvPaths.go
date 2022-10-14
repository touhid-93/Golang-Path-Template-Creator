package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type BaseEnvPaths struct {
	EnvPaths []string `json:"EnvPaths,omitempty"`
}

func (receiver *BaseEnvPaths) EnvPathsLength() int {
	if receiver == nil {
		return constants.Zero
	}

	return len(receiver.EnvPaths)
}

func (receiver *BaseEnvPaths) IsEmptyEnvPaths() bool {
	return receiver.EnvPathsLength() == 0
}

func (receiver *BaseEnvPaths) HasAnyItemEnvPaths() bool {
	return receiver.EnvPathsLength() > 0
}
