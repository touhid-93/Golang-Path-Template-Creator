package pathinsfmt

type Attributes struct {
	IsRecursive    bool           `json:"IsRecursive"`
	IsCache        bool           `json:"IsCache"`
	IsRedis        bool           `json:"IsRedis"`
	IsWriteToFiles bool           `json:"IsWriteToFiles"`
	CacheFilePath  string         `json:"CacheFilePath"`
	CachesRefresh  *CachesRefresh `json:"CachesRefresh,omitempty"`
}
