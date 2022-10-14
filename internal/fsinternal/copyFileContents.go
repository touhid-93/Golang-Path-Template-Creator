package fsinternal

import (
	"io"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdefer"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyFileContents(srcPath, dstPath string) (errWrap *errorwrapper.Wrapper) {
	if srcPath == dstPath {
		return nil
	}

	inFile, errOpen := os.Open(srcPath)

	defer errdefer.CloseFile(
		srcPath,
		errWrap,
		inFile)

	if errOpen != nil {
		return errnew.
			Path.
			Error(
				errtype.FileRead,
				errOpen,
				srcPath)
	}

	outFile, errCreate := os.Create(dstPath)

	defer errdefer.CloseFile(
		dstPath,
		errWrap,
		outFile)

	if errCreate != nil {
		return errnew.
			Path.
			Error(
				errtype.FileRead,
				errCreate,
				dstPath)
	}

	if _, err := io.Copy(outFile, inFile); err != nil {
		return errnew.SrcDst.Error(
			errtype.Copy,
			err,
			srcPath,
			dstPath,
		)
	}

	err2 := outFile.Sync()

	if err2 != nil {
		return errnew.SrcDst.Error(
			errtype.Copy,
			err2,
			srcPath,
			dstPath,
		)
	}

	return nil
}
