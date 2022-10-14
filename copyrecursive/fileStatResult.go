package copyrecursive

import "os"

type fileStatResult struct {
	filename string
	isExist  bool
	mode     os.FileMode
	err      error
}

func newFileStatResult(filePath string) *fileStatResult {
	fi, err := os.Lstat(filePath)

	if err != nil {
		return &fileStatResult{
			filename: filePath,
			isExist:  false,
			mode:     0,
			err:      err,
		}
	}

	return &fileStatResult{
		filename: filePath,
		isExist:  !os.IsNotExist(err),
		mode:     fi.Mode(),
		err:      err,
	}
}

func (it *fileStatResult) IsExist() bool {
	return it.isExist
}

func (it *fileStatResult) IsDir() bool {
	return it.IsExist() && it.mode.IsDir()
}

func (it *fileStatResult) IsRegular() bool {
	return it.IsExist() && it.mode.IsRegular()
}

func (it *fileStatResult) IsSymlink() bool {
	isSymlink := it.mode&os.ModeType == os.ModeSymlink

	return it.IsExist() && isSymlink
}
