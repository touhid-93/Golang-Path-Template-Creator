package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns .ssh path as string.
func SSHGlobal() string {
	return knowndir.SSHGlobal.CombineWith(UserPath())
}
