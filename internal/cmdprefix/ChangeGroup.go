package cmdprefix

// ChangeGroup Format: chgrp -R $group /dir
//  - Recursive (changeGroupRecursiveFormat) : chgrp -R $group
//  - Non-Recursive (changeGroupNonRecursiveFormat) : chgrp $group
func ChangeGroup(isRecursive bool, groupName string) string {
	return recursiveFormat(
		isRecursive,
		changeGroupRecursiveFormat,
		changeGroupNonRecursiveFormat,
		groupName)
}
