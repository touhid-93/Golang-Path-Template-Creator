package fscache

import (
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type newMapStringAnyCacheFileCreator struct{}

func (it newMapStringAnyCacheFileCreator) Default(
	isLock bool,
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newMapStringAnyCacheFileCreator) DefaultWriteEmptyOnNull(
	isLock bool,
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           isLock,
			IsRemoveFileBeforeWrite: true,
			IsWriteEmptyOnNull:      true,
		},
	}
}

func (it newMapStringAnyCacheFileCreator) DefaultLock(
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsAcquireLock:           true,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newMapStringAnyCacheFileCreator) DefaultNoLock(
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
		CacheFile: CacheFile{
			ChmodWrapper:            pathchmod.DefaultWrapper(),
			AbsFilePath:             location,
			IsRemoveFileBeforeWrite: true,
		},
	}
}

func (it newMapStringAnyCacheFileCreator) CollectReadWriteError(
	isLock bool,
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
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

func (it newMapStringAnyCacheFileCreator) CollectReadWriteErrorWriteEmptyOnNull(
	isLock bool,
	location string,
) *MapStringAnyCacheFile {
	return &MapStringAnyCacheFile{
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
