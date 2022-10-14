package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/fscache"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func hashsetReadWriteCache01() {
	cacheFile := fscache.HashsetCacheFile{
		CacheFile: fscache.CacheFile{
			ChmodWrapper: pathchmod.Wrapper{},
			AbsFilePath: pathjoin.WithTempPlusDefaults(
				"cache-file-testing",
				"hashset.conf"),
			IsAcquireLock:           false,
			IsCollectReadError:      false,
			IsCollectWriteError:     false,
			IsWriteEmptyOnNull:      false,
			IsRemoveFileBeforeWrite: true,
			OnReadActionInvoker: func(fileInfo *pathhelpercore.FileInfo) *errorwrapper.Wrapper {
				fmt.Println("read")
				fmt.Println(fileInfo.JsonPtr().PrettyJsonString())

				return nil
			},
			OnWriteActionInvoker: func(fileInfo *pathhelpercore.FileInfo) *errorwrapper.Wrapper {
				fmt.Println("write")
				fmt.Println(fileInfo.JsonPtr().PrettyJsonString())

				return nil
			},
			OnCacheExpireActionInvoker: func(fileInfo *pathhelpercore.FileInfo) *errorwrapper.Wrapper {
				fmt.Println("cache expire")
				fmt.Println(fileInfo.JsonPtr().PrettyJsonString())

				return nil
			},
			OnInvalidGenerateActionInvoker: func(fileInfo *pathhelpercore.FileInfo) *errorwrapper.Wrapper {
				fmt.Println("on invalidate generate")
				fmt.Println(fileInfo.JsonPtr().PrettyJsonString())

				return nil
			},
		},
	}

	fmt.Println("IsFileExist", cacheFile.IsFileExist())
	fmt.Println("IsFileIntegrityAlright", cacheFile.IsHashsetFileIntegrityAlright())

	coreHashset, readErrWrap := cacheFile.CoreHashset()
	readErrWrap.HandleError()

	fmt.Println(coreHashset.String())

	isAnyAdded, errWrap := cacheFile.AddOrUpdateManySave(
		"some key1",
		"some key 2",
		"some key 3")

	fmt.Println("is any added", isAnyAdded)
	errWrap.MustBeSafe()

	nextSlice := []string{
		"key1",
		"key2",
		"key3",
	}

	isAnyAdded2, errWrap := cacheFile.AddOrUpdateManySave(
		nextSlice...)

	fmt.Println("is any added", isAnyAdded2, nextSlice)
	errWrap.MustBeSafe()

	isAdded, coreHashset, errWrap := cacheFile.AddOrUpdateCoreHashsetSave(corestr.New.Hashset.StringsSpreadItems(
		"some new key"))
	errWrap.MustBeSafe()

	fmt.Println("is any added", isAdded, "some new key")
	fmt.Println("Final", coreHashset.String())
}
