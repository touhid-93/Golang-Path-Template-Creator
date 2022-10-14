package fscache

import (
	"gitlab.com/evatix-go/core/coretaskinfo"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

var (
	New         = newCreator{}
	globalMutex = consts.GlobalMutex
	CacheStates = cacheFileStates{
		ReadState:               "Read",
		WriteState:              "Write",
		CacheExpireState:        "Expired",
		InvalidateGenerateState: "InvalidateGenerate",
	}
	hashsetOnInvalidDefaultFunc = &errfunc.OnInvalidGenerator{
		NameInfo: &coretaskinfo.Info{
			RootName: "Hashset Cache file",
		},
		OnInvalidGenFunc: func(toPtr interface{}) *errorwrapper.Wrapper {
			return errnew.Reflect.SetFromTo(map[string]bool{}, toPtr)
		},
	}

	mapStringAnyOnInvalidDefaultFunc = &errfunc.OnInvalidGenerator{
		NameInfo: &coretaskinfo.Info{
			RootName: "Map String to Any Value Cache file",
		},
		OnInvalidGenFunc: func(toPtr interface{}) *errorwrapper.Wrapper {
			return errnew.Reflect.SetFromTo(map[string]interface{}{}, toPtr)
		},
	}

	hashmapOnInvalidDefaultFunc = &errfunc.OnInvalidGenerator{
		NameInfo: &coretaskinfo.Info{
			RootName: "Hashmap Cache file",
		},
		OnInvalidGenFunc: func(toPtr interface{}) *errorwrapper.Wrapper {
			return errnew.Reflect.SetFromTo(map[string]string{}, toPtr)
		},
	}

	stringsOnInvalidDefaultFunc = &errfunc.OnInvalidGenerator{
		NameInfo: &coretaskinfo.Info{
			RootName: "Strings Cache file",
		},
		OnInvalidGenFunc: func(toPtr interface{}) *errorwrapper.Wrapper {
			return errnew.Reflect.SetFromTo([]string{}, toPtr)
		},
	}
)
