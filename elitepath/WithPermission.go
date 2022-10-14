package elitepath

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathfixer"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

type WithPermission struct {
	Path
	BasePathPermission
	PathVerifiers *pathinsfmt.PathVerifiers `json:"PathVerifiers,omitempty"`
	PathVerifier  *pathinsfmt.PathVerifier  `json:"PathVerifier,omitempty"`
}

func NewWithPermissionRwxChown(
	isNormalize,
	isExpandEnvVars bool,
	rwx *chmodins.RwxInstruction,
	chown *pathinsfmt.Chown,
	location string,
) WithPermission {
	return WithPermission{
		Path: Path{
			Location: pathfixer.Location{
				PathOptions: pathfixer.PathOptions{
					IsContinueOnError: false,
					IsNormalize:       isNormalize,
					IsExpandEnvVar:    isExpandEnvVars,
					IsRecursive:       false,
					IsSkipOnInvalid:   false,
					IsSkipOnExist:     false,
					IsSkipOnEmpty:     false,
					IsRelative:        false,
				},
				Path: location,
			},
		},
		BasePathPermission: BasePathPermission{
			RwxInstruction: rwx,
			Chown:          chown,
		},
	}
}

func NewWithPermissionPtr(
	isNormalize,
	isExpandEnvVars bool,
	location string,
) *WithPermission {
	return NewWithPermissionRwxChownPtr(
		isNormalize,
		isExpandEnvVars,
		nil,
		nil,
		location)
}

func NewWithPermissionRwxChownPtr(
	isNormalize,
	isExpandEnvVars bool,
	rwx *chmodins.RwxInstruction,
	chown *pathinsfmt.Chown,
	location string,
) *WithPermission {
	return &WithPermission{
		Path: Path{
			Location: pathfixer.Location{
				PathOptions: pathfixer.PathOptions{
					IsContinueOnError: false,
					IsNormalize:       isNormalize,
					IsExpandEnvVar:    isExpandEnvVars,
					IsRecursive:       false,
					IsSkipOnInvalid:   false,
					IsSkipOnExist:     false,
					IsSkipOnEmpty:     false,
					IsRelative:        false,
				},
				Path: location,
			},
		},
		BasePathPermission: BasePathPermission{
			RwxInstruction: rwx,
			Chown:          chown,
		},
	}
}

func (it *WithPermission) Apply(
	isSkipOnInvalid bool,
) *errwrappers.Collection {
	errorsCollection := errwrappers.Empty()

	if it.IsEmptyPath() {
		return errorsCollection
	}

	it.ApplyUsingErrorCollection(
		isSkipOnInvalid,
		errorsCollection)

	return errorsCollection
}

func (it *WithPermission) ApplyUsingErrorCollection(
	isSkipOnInvalid bool,
	errorCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmptyPath() {
		return true
	}

	stateTracker := errorCollection.StateTracker()
	errorCollection.AddAllFunctions(
		it.ApplyRwx,
		it.ApplyChown,
	)

	it.ApplyPathVerifiers(errorCollection)
	it.ApplyPathVerifier(
		isSkipOnInvalid,
		errorCollection)

	return stateTracker.IsSuccess()
}

func (it *WithPermission) ApplyChown() *errorwrapper.Wrapper {
	if it.IsEmptyChown() ||
		it.IsEmptyPath() {
		return nil
	}

	return it.Path.ApplyChown(it.Chown)
}

func (it *WithPermission) ApplyRwx() *errorwrapper.Wrapper {
	if it.IsEmptyRwx() ||
		it.IsEmptyPath() {
		return nil
	}

	return it.ApplyRwxInstruction(it.RwxInstruction)
}

func (it *WithPermission) ApplyPathVerifiers(
	errorCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmptyPath() || it.PathVerifiers == nil {
		return true
	}

	return it.Path.ApplyPathVerifiers(
		errorCollection,
		it.PathVerifiers)
}

func (it *WithPermission) ApplyPathVerifier(
	isSkipOnInvalid bool,
	errorCollection *errwrappers.Collection,
) (isSuccess bool) {
	if it.IsEmptyPath() || it.PathVerifier == nil {
		return true
	}

	return it.Path.ApplyPathVerifier(
		isSkipOnInvalid,
		errorCollection,
		it.PathVerifier)
}
