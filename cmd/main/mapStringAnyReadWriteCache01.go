package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/fscache"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func mapStringAnyReadWriteCache01() {
	cacheFile := fscache.MapStringAnyCacheFile{
		CacheFile: fscache.CacheFile{
			ChmodWrapper: pathchmod.Wrapper{},
			AbsFilePath: pathjoin.WithTempPlusDefaults(
				"cache-file-testing",
				"MapStringAnyCacheFile.conf"),
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
	fmt.Println("IsOnInvalidateGenerateActionInvokerDefined", cacheFile.IsOnInvalidGeneratorDefined())
	fmt.Println("IsFileIntegrityAlright", cacheFile.IsFileIntegrityAlright())

	currentMap, readErrWrap := cacheFile.Read()
	readErrWrap.HandleError()

	fmt.Println(currentMap)

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

	errWrap2 := cacheFile.AddOrUpdateMapSave(map[string]interface{}{
		"some key":  "val 1",
		"some key2": 2,
		"some key3": "so",
		"some key4": "val 4",
	})
	errWrap2.MustBeSafe()

	fmt.Println("Final", cacheFile.Strings())
}
