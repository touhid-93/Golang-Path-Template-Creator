package copyrecursive

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type Copier struct {
	src  string
	dst  string
	opts Options
}

func NewCopier(
	src, dst string,
	opts Options,
) *Copier {
	src = opts.FixedPath(src)
	dst = opts.FixedPath(dst)

	return &Copier{
		src:  src,
		dst:  dst,
		opts: opts,
	}
}

func NewCopierUsingInstruction(
	instruction Instruction,
) *Copier {
	return NewCopier(
		instruction.Source,
		instruction.Destination,
		instruction.Options)
}

func (it *Copier) Copy() *errorwrapper.Wrapper {
	if it.opts.IsUseShellOrCmd && osconsts.IsLinux {
		return copyUsingLinuxCP(
			it.opts,
			it.src,
			it.dst)
	}

	if it.opts.IsClearDestination {
		err := os.RemoveAll(it.dst)

		if err != nil {
			return errnew.Ref.TwoWithError(
				errtype.DeleteFailed,
				err,
				"src",
				it.src,
				"dst",
				it.dst)
		}
	}

	srcFileStat := newFileStatResult(it.src)
	if srcFileStat.IsRegular() {
		copyErr := it.copyFile(it.src, it.dst)

		if copyErr != nil {
			return errnew.Ref.TwoWithError(
				errtype.Copy,
				copyErr,
				"src",
				it.src,
				"dst",
				it.dst)
		}

		// Only delete source iff the copy is successful
		if it.opts.IsMove {
			errRemove := os.RemoveAll(it.src)

			return errnew.Ref.TwoWithError(
				errtype.RemoveFailed,
				errRemove,
				"src",
				it.src,
				"dst",
				it.dst)
		}
	}

	copyDirErr := it.copyDir(
		it.opts.IsRecursive,
		it.src,
		it.dst)

	if copyDirErr != nil {
		return errnew.Ref.TwoWithError(
			errtype.Copy,
			copyDirErr,
			"src",
			it.src,
			"dst",
			it.dst)
	}

	// Remove src iff the copy is successful
	if it.opts.IsMove {
		errRemove := os.RemoveAll(it.src)

		return errnew.Ref.TwoWithError(
			errtype.RemoveFailed,
			errRemove,
			"src",
			it.src,
			"dst",
			it.dst)
	}

	return nil
}

// copySymLink copies a symbolic link from src to dst.
func (it *Copier) copySymLink(src, dst string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}

	return os.Symlink(link, dst)
}

func (it *Copier) createDir(
	dir string,
	perm os.FileMode,
) error {
	if isExists(dir) && it.opts.IsSkipOnExist {
		return nil
	}

	return os.MkdirAll(dir, perm)
}

func (it *Copier) copyFile(src, dst string) error {
	if it.opts.IsSkipOnExist &&
		isFileExists(it.dst) {
		return nil
	}

	return CopyFile(src, dst, defaultFileMode)
}

// copyDir copies a src directory to a destination.
func (it *Copier) copyDir(isRecursive bool, src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := pathjoin.JoinSimple(src, entry.Name())
		destPath := pathjoin.JoinSimple(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		// Flag check
		if it.opts.IsSkipOnExist {
			if isExists(destPath) {
				continue
			}
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if !isRecursive {
				continue
			}

			// Should non-recursive version also create the empty directories?
			// Probably not
			if err := it.createDir(destPath, defaultFileMode); err != nil {
				return err
			}

			if err := it.copyDir(true, sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := it.copySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := it.copyFile(sourcePath, destPath); err != nil {
				return err
			}
		}

		// `go test` fails on Windows even with this `if` supposedly
		// protecting the `syscall.Stat_t` and `os.Lchown` calls (not
		// available on windows). why?
		/*
			if runtime.GOOS != "windows" {
				stat, ok := fileInfo.Sys().(*syscall.Stat_t)
				if !ok {
					return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
				}
				if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
					return err
				}
			}
		*/
		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			// Changing the destination permission as same as the source permission
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}

	return nil
}
