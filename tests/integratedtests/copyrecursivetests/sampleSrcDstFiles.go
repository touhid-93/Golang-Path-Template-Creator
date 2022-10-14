package copyrecursivetests

import (
	"path/filepath"

	"gitlab.com/evatix-go/pathhelper/tests/testwrappers/copyrecursivetestwrapper"
)

func sampleSrcDstFiles(
	srcRoot, dstRoot string,
) (srcFiles, dstFiles []string) {
	for _, file := range copyrecursivetestwrapper.RelPathOfSrcFiles {
		srcFiles = append(srcFiles, filepath.Join(srcRoot, file))
		dstFiles = append(dstFiles, filepath.Join(dstRoot, file))
	}

	return srcFiles, dstFiles
}
