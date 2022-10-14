package pathfixer

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

type PathOptions struct {
	IsContinueOnError bool `json:"IsContinueOnError,omitempty"`
	IsNormalize       bool `json:"IsNormalize,omitempty"`
	IsExpandEnvVar    bool `json:"IsExpandEnvVar,omitempty"` // can expand environment variables with %{Name} $Name %{Java_home} ${Java_HOME}
	IsRecursive       bool `json:"IsRecursive,omitempty"`
	IsSkipOnInvalid   bool `json:"IsSkipOnInvalid,omitempty"`
	IsSkipOnExist     bool `json:"IsSkipOnExist,omitempty"`
	IsSkipOnEmpty     bool `json:"IsSkipOnEmpty,omitempty"`
	IsRelative        bool `json:"IsRelative,omitempty"`
}

func (it *PathOptions) GetFixedPath(location string) string {
	return pathjoin.FixPath(
		it.IsNormalize,
		it.IsExpandEnvVar,
		location)
}

func (it *PathOptions) GetFixedPathJoined(location1, location2 string) string {
	return pathjoin.JoinConditionalNormalizedExpandIf(
		it.IsNormalize,
		it.IsExpandEnvVar,
		location1,
		location2)
}

func (it *PathOptions) Join(location1, location2 string) string {
	return pathjoin.JoinSimple(
		location1,
		location2)
}

func (it *PathOptions) GetFixedPathJoinedMany(
	locations ...string,
) string {
	return pathjoin.JoinWithSep(
		it.IsSkipOnEmpty,
		it.IsExpandEnvVar,
		it.IsNormalize,
		osconsts.PathSeparator,
		locations...)
}

func (it *PathOptions) ClonePathOptions() *PathOptions {
	if it == nil {
		return nil
	}

	return &PathOptions{
		IsContinueOnError: it.IsContinueOnError,
		IsNormalize:       it.IsNormalize,
		IsExpandEnvVar:    it.IsExpandEnvVar,
		IsRecursive:       it.IsRecursive,
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsSkipOnExist:     it.IsSkipOnExist,
		IsSkipOnEmpty:     it.IsSkipOnEmpty,
		IsRelative:        it.IsRelative,
	}
}

func (it *PathOptions) IsEqual(another *PathOptions) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.IsContinueOnError == another.IsContinueOnError &&
		it.IsNormalize == another.IsNormalize &&
		it.IsExpandEnvVar == another.IsExpandEnvVar &&
		it.IsRecursive == another.IsRecursive &&
		it.IsSkipOnInvalid == another.IsSkipOnInvalid &&
		it.IsSkipOnExist == another.IsSkipOnExist &&
		it.IsSkipOnEmpty == another.IsSkipOnEmpty &&
		it.IsRelative == another.IsRelative
}
