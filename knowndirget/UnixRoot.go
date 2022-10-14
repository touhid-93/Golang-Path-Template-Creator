package knowndirget

import "gitlab.com/evatix-go/pathhelper/knowndir"

// Returns unix system root as a string
func UnixRoot() string {
	return knowndir.UnixRoot.Value()
}
