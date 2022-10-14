package ispath

func NotEqualString(
	isApplyNormalize bool,
	left, right string,
) bool {
	return !EqualString(
		isApplyNormalize,
		left, right)
}
