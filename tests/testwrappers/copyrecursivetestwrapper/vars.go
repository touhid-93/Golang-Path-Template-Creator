package copyrecursivetestwrapper

import (
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

var (
	SrcTestDataDir = pathjoin.JoinNormalized(
		currentDir(), SourceDataDirName)
	SourceRecursivePath = pathjoin.JoinNormalized(
		SrcTestDataDir,
		SourceRecursiveDirName)
	Destination = pathjoin.WithTempTest(PkgDirName)

	RelPathOfSrcFiles = []string{
		filepath.Join("dir1", "a.txt"),
		filepath.Join("dir2", "b.txt"),
		"file1.txt",
		"file2.txt",
	}

	IndexOfExistingFilesInSkipOnExistDir    = []int{0, 2}
	IndexOfNonExistingFilesInSkipOnExistDir = []int{1, 3}
)

func currentDir() string {
	_, f, _, _ := runtime.Caller(1)

	return filepath.Dir(f)
}
