package namegroup

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
	"gitlab.com/evatix-go/pathhelper/internal/deferrwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyLinuxUserGroupBoth(
	isRecursive bool,
	isContinueOnError bool,
	userNameGroupName *pathinsfmt.UserGroupName,
	paths ...string,
) *errorwrapper.Wrapper {
	if userNameGroupName == nil {
		return nil
	}

	if userNameGroupName.IsGroupNameEmpty() {
		return deferrwrappers.
			CannotApplyChmodWithSingleParameter.
			ConcatNew().
			Messages("Group name empty or not defined.")
	}

	if userNameGroupName.IsUserNameEmpty() {
		return deferrwrappers.
			CannotApplyChmodWithSingleParameter.
			ConcatNew().
			Messages("User name empty or not defined.")
	}

	pathsLength := len(paths)

	if pathsLength == 0 {
		return nil
	}

	groupName := userNameGroupName.GroupName
	userName := userNameGroupName.UserName

	// chown -R $user:$group /dir
	cmdPrefix := cmdprefix.ChownUser(
		isRecursive,
		userName,
		groupName)

	return applyCmdOnPaths(
		cmdPrefix,
		paths,
		isContinueOnError)
}
