package nginxlinuxpath

func NewNginxDirDefaultChmod(
	isNormalize bool,
	currentNginxRoot,
	username string,
) *NginxDir {
	return NewNginxDir(
		isNormalize,
		DefaultDirChmod,
		currentNginxRoot,
		username)
}
