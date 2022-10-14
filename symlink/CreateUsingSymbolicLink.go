package symlink

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func CreateUsingSymbolicLink(symLink *pathinsfmt.SymbolicLink) *errorwrapper.Wrapper {
	if symLink == nil {
		return nil
	}

	isFileExist := fsinternal.IsPathExists(symLink.Src)
	isFileMissing := !isFileExist

	if symLink.IsSkipOnSrcMissing && isFileMissing {
		return nil
	}

	if !symLink.IsSkipOnSrcMissing && isFileMissing {
		return errnew.
			Path.
			Messages(
				errtype.PathNotFound,
				symLink.Src,
				"Cannot apply or create symbolic link when path is missing. Please select IsSkipOnSrcMissing to ignore.")
	}

	if symLink.IsClearBefore {
		errWrap := fsinternal.SafeRemove(symLink.Dst)

		if errWrap.HasError() {
			return errWrap
		}
	}

	if symLink.IsSkipOnExist && isFileExist {
		return nil
	}

	if symLink.IsMkDirAll {
		errWrap := fsinternal.CreateDirectoryAllUptoParentDefault(
			symLink.Dst)

		if errWrap.HasError() {
			return errWrap
		}
	}

	dst := symLink.Dst

	if fsinternal.IsDirectory(dst) {
		sourceFileName := fsinternal.GetFileName(symLink.Src)
		dst += osconsts.PathSeparator + sourceFileName
		dst = normalizeinternal.Fix(dst)
	}

	err := os.Symlink(
		symLink.Src,
		dst)

	if err != nil {
		return errnew.Ref.ManyWithError(
			errtype.SymbolicLink,
			err,
			ref.Value{
				Variable: "Source Symbolic Link",
				Value:    symLink.Src,
			},
			ref.Value{
				Variable: "Destination Symbolic Link",
				Value:    dst,
			},
			ref.Value{
				Variable: "Full Symbolic Link Request",
				Value:    symLink,
			})
	}

	return nil
}
