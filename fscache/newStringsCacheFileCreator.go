package fscache

import (
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type newStringsCacheFileCreator struct{}

func (it newStringsCacheFileCreator) Default(
	isLock bool,
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newStringsCacheFileCreator) DefaultWriteEmptyOnNull(
	isLock bool,
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsWriteEmptyOnNull:      true,
		},
	}
}

func (it newStringsCacheFileCreator) DefaultLock(
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           true,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newStringsCacheFileCreator) DefaultNoLock(
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newStringsCacheFileCreator) CollectReadWriteError(
	isLock bool,
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
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

func (it newStringsCacheFileCreator) CollectReadWriteErrorWriteEmptyOnNull(
	isLock bool,
	location string,
) *StringsCacheFile {
	return &StringsCacheFile{
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
