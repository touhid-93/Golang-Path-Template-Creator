package fscache

import (
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type newCacheFileCreator struct{}

func (it newCacheFileCreator) Default(
	isLock bool,
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsAcquireLock:           isLock,
		IsRemoveFileBeforeWrite: true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}

func (it newCacheFileCreator) DefaultWriteEmptyOnNull(
	isLock bool,
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsAcquireLock:           isLock,
		IsRemoveFileBeforeWrite: true,
		IsWriteEmptyOnNull:      true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}

func (it newCacheFileCreator) DefaultLock(
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsAcquireLock:           true,
		IsRemoveFileBeforeWrite: true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}

func (it newCacheFileCreator) DefaultNoLock(
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsRemoveFileBeforeWrite: true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}

func (it newCacheFileCreator) CollectReadWriteError(
	isLock bool,
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsAcquireLock:           isLock,
		IsRemoveFileBeforeWrite: true,
		IsCollectReadError:      true,
		IsCollectWriteError:     true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}

func (it newCacheFileCreator) CollectReadWriteErrorWriteEmptyOnNull(
	isLock bool,
	onInvalidGenerator errfunc.OnInvalidGenerator,
	location string,
) *CacheFile {
	return &CacheFile{
		ChmodWrapper:            pathchmod.DefaultWrapper(),
		AbsFilePath:             location,
		IsAcquireLock:           isLock,
		IsRemoveFileBeforeWrite: true,
		IsCollectReadError:      true,
		IsCollectWriteError:     true,
		IsWriteEmptyOnNull:      true,
		OnInvalidGenerator:      &onInvalidGenerator,
	}
}
