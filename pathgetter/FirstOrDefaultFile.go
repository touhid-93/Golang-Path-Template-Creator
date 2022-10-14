package pathgetter

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func FirstOrDefaultFile(
	isNormalize bool,
	exploringPath string,
) *errstr.Result {
	results := FilesDefault(
		isNormalize,
		exploringPath)

	return results.FirstOrDefaultResult()
}
