package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type BaseEnvironmentVariables struct {
	EnvVars []EnvironmentVariable `json:"EnvVars,omitempty"`
}

func (receiver *BaseEnvironmentVariables) EnvVarsLength() int {
	if receiver == nil {
		return constants.Zero
	}

	return len(receiver.EnvVars)
}

func (receiver *BaseEnvironmentVariables) IsEmptyEnvVars() bool {
	return receiver.EnvVarsLength() == 0
}

func (receiver *BaseEnvironmentVariables) HasAnyItemEnvVars() bool {
	return receiver.EnvVarsLength() > 0
}
