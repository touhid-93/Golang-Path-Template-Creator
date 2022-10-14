package pathjoin

import "path"

// JoinsNormalizedIf normalized and expand applied if condition meets
func JoinsNormalizedIf(
	isNormalize bool,
	baseDir string,
	relatives ...string,
) string {
	joined := path.Join(relatives...)

	return JoinNormalizedIf(
		isNormalize,
		baseDir,
		joined)
}
