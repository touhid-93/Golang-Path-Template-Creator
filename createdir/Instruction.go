package createdir

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

type Instruction struct {
	Location string
	FileMode os.FileMode // 0 means no apply
	IsLock,
	IsCreate,
	IsSkipOnExist,
	IsRemoveAll bool
}

func (it *Instruction) ParentDir() string {
	return fsinternal.ParentDir(it.Location)
}

func (it *Instruction) HasFileMode() bool {
	return it.FileMode > 0
}

func (it *Instruction) FileModeOrDefault() os.FileMode {
	if it.HasFileMode() {
		return it.FileMode
	}

	return DefaultDirectoryFileMode
}

func (it *Instruction) CreateIf(isCreate bool, mode os.FileMode) *errorwrapper.Wrapper {
	return RemoveCreateAll(
		it.IsLock,
		it.IsSkipOnExist,
		it.IsRemoveAll,
		isCreate,
		it.Location,
		mode)
}

func (it *Instruction) Create(mode os.FileMode) *errorwrapper.Wrapper {
	return RemoveCreateAll(
		it.IsLock,
		it.IsSkipOnExist,
		it.IsRemoveAll,
		it.IsCreate,
		it.Location,
		mode)
}

func (it *Instruction) CreateParent(mode os.FileMode) *errorwrapper.Wrapper {
	parent := it.ParentDir()

	return RemoveCreateAll(
		it.IsLock,
		it.IsSkipOnExist,
		it.IsRemoveAll,
		it.IsCreate,
		parent,
		mode)
}

func (it *Instruction) CreateParentIf(isCreate bool, mode os.FileMode) *errorwrapper.Wrapper {
	parent := it.ParentDir()

	return RemoveCreateAll(
		it.IsLock,
		it.IsSkipOnExist,
		it.IsRemoveAll,
		isCreate,
		parent,
		mode)
}

func (it *Instruction) CreateDefault() *errorwrapper.Wrapper {
	return RemoveCreateAll(
		it.IsLock,
		it.IsSkipOnExist,
		it.IsRemoveAll,
		it.IsCreate,
		it.Location,
		it.FileModeOrDefault())
}

func (it *Instruction) IsExist() bool {
	return fsinternal.IsPathExists(it.Location)
}

func (it *Instruction) IsDirExist() bool {
	return fsinternal.IsDirectory(it.Location)
}
