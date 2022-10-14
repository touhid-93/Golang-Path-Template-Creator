package pathchmod

import "gitlab.com/evatix-go/core/chmodhelper"

var (
	FriendlyChmod = friendlyChmod{}
	rwxCreator    = chmodhelper.New.RwxWrapper.UsingFileMode
)
