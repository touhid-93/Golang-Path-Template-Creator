package fscache

import (
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type newHashsetCacheFileCreator struct{}

func (it newHashsetCacheFileCreator) Default(
	isLock bool,
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashsetCacheFileCreator) DefaultWriteEmptyOnNull(
	isLock bool,
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsWriteEmptyOnNull:      true,
		},
	}
}

func (it newHashsetCacheFileCreator) DefaultLock(
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           true,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashsetCacheFileCreator) DefaultNoLock(
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashsetCacheFileCreator) CollectReadWriteError(
	isLock bool,
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsCollectReadError:      true,
			IsCollectWriteError:     true,
		},
	}
}

func (it newHashsetCacheFileCreator) CollectReadWriteErrorWriteEmptyOnNull(
	isLock bool,
	location string,
) *HashsetCacheFile {
	return &HashsetCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsCollectReadError:      true,
			IsCollectWriteError:     true,
			IsWriteEmptyOnNull:      true,
		},
	}
}
