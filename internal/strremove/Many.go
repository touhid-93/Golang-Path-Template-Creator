package strremove

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func SimpleMany(content string, removeRequests ...string) string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return content
}

// SimpleManySplitsBy Remove as per removes then splits by the given separator
func SimpleManySplitsBy(content string, splitsBy string, removeRequests ...string) []string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return strings.Split(content, splitsBy)
}

func SimpleManyLimits(content string, limits int, removeRequests ...string) string {
	for _, remove := range removeRequests {
		content = strings.Replace(content, remove, constants.EmptyString, limits)
	}

	return content
}
