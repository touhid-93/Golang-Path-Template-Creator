package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type PathWithModifier struct {
	PathWithOptions
	Modifier *PathModifier `json:"Modifier,omitempty"`
}

func (it *PathWithModifier) IsModifierUndefined() bool {
	return it == nil || it.Modifier == nil
}

func (it *PathWithModifier) IsModifierDefined() bool {
	return it != nil && it.Modifier != nil
}

func (it *PathWithModifier) IsPathWithModifierDefined() bool {
	return it != nil && it.Path != "" && it.Modifier != nil
}

func (it *PathWithModifier) Clone() *PathWithModifier {
	if it == nil {
		return nil
	}

	return &PathWithModifier{
		PathWithOptions: *it.PathWithOptions.ClonePath(),
		Modifier:        it.Modifier.Clone(),
	}
}

// PathWithVerifier creates the user group validation from Modifier.Chown
func (it *PathWithModifier) PathWithVerifier() *PathWithVerifier {
	if it == nil {
		return nil
	}

	userGroupName := it.Modifier.Chown.UserGroupName.Clone()
	baseRwxInstructions := it.Modifier.BaseRwxInstructions.Clone()
	var userGroupNameSetter UserGroupName
	var baseRwxInstructionsSetter chmodins.BaseRwxInstructions

	if userGroupName != nil {
		userGroupNameSetter = *userGroupName
	}

	if baseRwxInstructions != nil {
		baseRwxInstructionsSetter = *baseRwxInstructions
	}

	return &PathWithVerifier{
		PathWithOptions: *it.PathWithOptions.ClonePath(),
		Verifier: &PathVerifier{
			UserGroupName:       userGroupNameSetter,
			BaseRwxInstructions: baseRwxInstructionsSetter,
		},
	}
}
