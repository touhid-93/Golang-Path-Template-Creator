package verifypath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/ispath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func PrefixError(
	isVerifyExistence bool,
	homeDirPrefix, currentFullPath string,
) (fixedPath string, verificationErrorWrapper *errorwrapper.Wrapper) {
	prefixFix := normalize.Path(homeDirPrefix)
	currentPathFix := normalize.Path(currentFullPath)

	if isVerifyExistence && ispath.NotExists(currentPathFix) {
		// error
		return constants.EmptyString, errnew.
			Path.
			Messages(
				errtype.FileInvalidOrMissing,
				currentPathFix,
				"given path is not valid in the file system.",
				"path must contain user home prefix and don't contain any relative path!",
				"user home prefix:",
				prefixFix,
			)
	}

	if normalize.HasPrefix(prefixFix, currentPathFix) {
		return currentPathFix, nil
	}

	// homeDirPrefix doesn't match
	return constants.EmptyString, errnew.Ref.Messages(
		errtype.PathSyntaxIssue,
		"Prefix",
		prefixFix,
		"current path homeDirPrefix missing (\""+homeDirPrefix+"\")",
		"path needs to have user home dir!",
		"current path:",
		currentPathFix,
		"homeDirPrefix",
		prefixFix,
	)
}
