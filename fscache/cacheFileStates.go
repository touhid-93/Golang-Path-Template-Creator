package fscache

import "gitlab.com/evatix-go/enum/strtype"

type cacheFileStates struct {
	ReadState               strtype.Variant
	WriteState              strtype.Variant
	CacheExpireState        strtype.Variant
	InvalidateGenerateState strtype.Variant
}

func (it cacheFileStates) IsRead(r strtype.Variant) bool {
	return r.IsEqualAnother(it.ReadState)
}

func (it cacheFileStates) IsWrite(w strtype.Variant) bool {
	return w.IsEqualAnother(it.WriteState)
}

func (it cacheFileStates) IsExpire(n strtype.Variant) bool {
	return n.IsEqualAnother(it.CacheExpireState)
}

func (it cacheFileStates) IsInvalidateGenerateRequired(n strtype.Variant) bool {
	return n.IsEqualAnother(it.InvalidateGenerateState)
}
