package cmdprefix

// ChownUser Format: chown -R $user:$group /dir
//  - Recursive : chown -R $user:$group /dir
//  - Non Recursive : chown $user:$group /dir
func ChownUser(isRecursive bool, userName, groupName string) string {
	return recursiveFormat(
		isRecursive,
		chownRecursiveFormat,
		chownNonRecursiveFormat,
		userName,
		groupName)
}
