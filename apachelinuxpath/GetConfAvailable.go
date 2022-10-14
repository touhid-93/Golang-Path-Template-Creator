package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetConfAvailable returns /etc/apache/conf-available as a string
func GetConfAvailable() string {
	if osconsts.IsWindows {
		errtype.NotSupportInWindows.PanicNoRefs(constants.EmptyString)
	}

	return knowndir.ConfAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
