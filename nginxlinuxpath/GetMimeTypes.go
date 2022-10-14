package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetMimeTypes
//
// returns /etc/nginx/mime.types as a string
func GetMimeTypes() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	if defaultMimeTypesPath.IsInitialized() {
		return defaultMimeTypesPath.String()
	}

	mimePath := knowndir.MimeTypes.CombineWith(
		knowndirget.NginxLinuxPath())

	return defaultMimeTypesPath.GetPlusSetOnUninitialized(
		mimePath)
}
