package pathjoin

import "gitlab.com/evatix-go/pathhelper/normalize"

func JoinFixedThree(
	first,
	second,
	third string,
) string {
	joined := JoinSimpleThree(
		first,
		second,
		third)

	return normalize.Path(joined)
}
