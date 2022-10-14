package pathinsfmt

type BasePathsWithModifiers struct {
	PathWithModifiers []PathWithModifier `json:"PathWithModifiers,omitempty"`
}

func (it *BasePathsWithModifiers) IsPathWithModifiersUndefined() bool {
	return it == nil || len(it.PathWithModifiers) == 0
}

func (it *BasePathsWithModifiers) IsPathWithModifiersDefined() bool {
	return it != nil || len(it.PathWithModifiers) > 0
}

func (it *BasePathsWithModifiers) BasePathsWithVerifiers(
	isContinueOnError bool,
) *BasePathsWithVerifiers {
	if it.IsPathWithModifiersUndefined() {
		return &BasePathsWithVerifiers{
			PathsWithVerifiers: &PathsWithVerifiers{
				IsContinueOnError: isContinueOnError,
				Verifiers:         []PathWithVerifier{},
			},
		}
	}

	slice := make([]PathWithVerifier, len(it.PathWithModifiers))
	for i, modifier := range it.PathWithModifiers {
		slice[i] = *modifier.PathWithVerifier()
	}

	return &BasePathsWithVerifiers{
		PathsWithVerifiers: &PathsWithVerifiers{
			IsContinueOnError: isContinueOnError,
			Verifiers:         slice,
		},
	}
}
