package pathsysinfo

import (
	"os/user"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func LookupUser(userName string) (userResult *user.User, errorWrapper *errorwrapper.Wrapper) {
	userObj, errLookup := user.Lookup(userName)

	if errLookup != nil {
		return nil, errorwrapper.NewRef(
			codestack.SkipNone,
			errtype.SearchFailed,
			errLookup,
			"UserName",
			userName)
	}

	return userObj, nil
}
