package pathsysinfo

func GetUserGroupId(filePath string) *UserGroupId {
	instance := GetPathUserGroupId(filePath)

	if instance != nil {
		return instance.UserGroupId()
	}

	return nil
}
