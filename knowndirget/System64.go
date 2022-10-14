package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to SysWOW64 path in windows system.
func GetSystem64() string {
	return knowndir.System64.CombineWith(WidowsDirectory())
}
