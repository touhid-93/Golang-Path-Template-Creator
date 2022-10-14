package fscache

import (
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type newHashmapCacheFileCreator struct{}

func (it newHashmapCacheFileCreator) Default(
	isLock bool,
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashmapCacheFileCreator) DefaultWriteEmptyOnNull(
	isLock bool,
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsWriteEmptyOnNull:      true,
		},
	}
}

func (it newHashmapCacheFileCreator) DefaultLock(
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           true,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashmapCacheFileCreator) DefaultNoLock(
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newHashmapCacheFileCreator) CollectReadWriteError(
	isLock bool,
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
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

func (it newHashmapCacheFileCreator) CollectReadWriteErrorWriteEmptyOnNull(
	isLock bool,
	location string,
) *HashmapCacheFile {
	return &HashmapCacheFile{
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
