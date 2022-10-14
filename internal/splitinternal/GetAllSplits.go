package splitinternal

import "gitlab.com/evatix-go/core/constants"

// GetAllSplits Ref : https://play.golang.org/p/oT6eWNZAeEi
func GetAllSplits(currentPath string) (baseDirNames *[]string) {
	i := len(currentPath) - 1
	list := make(
		[]string,
		0,
		constants.ArbitraryCapacity10)

	lastFoundSeparator := i + 1
	name := constants.EmptyString
	isIndexZero := i == 0
	for i >= 0 {
		isSeparator :=
			currentPath[i] == constants.ForwardChar ||
				currentPath[i] == constants.BackwardChar
		isIndexZero = i == 0

		if isIndexZero {
			name = currentPath[i:lastFoundSeparator]
		} else if isSeparator {
			name = currentPath[i+1 : lastFoundSeparator]
		}

		if isSeparator || isIndexZero {
			list = append(
				list,
				name)

			lastFoundSeparator = i
		}

		i--
	}

	return &list
}
