package pathhelpercore

import "gitlab.com/evatix-go/errorwrapper"

type BaseErrorWrapper struct {
	ErrorWrapper *errorwrapper.Wrapper `json:"ErrorWrapper,omitempty"`
}
