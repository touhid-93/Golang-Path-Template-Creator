package pathjoin

import "path"

// JoinsNormalizedExpandIf normalized and expand applied if condition meets
func JoinsNormalizedExpandIf(
	isExpandNormalize bool,
	baseDir string,
	relatives ...string,
) string {
	joined := path.Join(relatives...)

	return JoinNormalizedExpandIf(
		isExpandNormalize,
		baseDir,
		joined)
}
