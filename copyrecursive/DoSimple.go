package copyrecursive

// DoSimple recursively copies a src directory to a destination.
func DoSimple(src, dst string) error {
	return NewCopier(
		src,
		dst,
		Options{
			IsSkipOnExist:      false,
			IsRecursive:        true,
			IsClearDestination: false,
			IsUseShellOrCmd:    false,
			IsNormalize:        false,
		},
	).Copy().Error()
}
