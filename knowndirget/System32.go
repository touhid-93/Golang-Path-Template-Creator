package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to System32 path in windows system.
func System32() string {
	return knowndir.
		System32.
		CombineWith(WidowsDirectory())
}
