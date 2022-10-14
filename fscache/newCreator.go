package fscache

type newCreator struct {
	CacheFile        newCacheFileCreator
	StringsCacheFile newStringsCacheFileCreator
}
