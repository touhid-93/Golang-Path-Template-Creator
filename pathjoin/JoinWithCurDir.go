package pathjoin

import (
	"path"
	"path/filepath"
	"runtime"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func JoinWithCurDir(
	isNormalize,
	isExpandEnv bool,
	paths ...string,
) string {
	_, b, _, _ := runtime.Caller(constants.One)
	finalSlice := stringslice.PrependLineNew(filepath.Dir(b), paths)
	joined := path.Join(finalSlice...)

	expanded := expandpath.ExpandVariablesIf(
		isExpandEnv,
		joined)

	return normalize.PathUsingSeparatorIf(
		false,
		isNormalize,
		isNormalize,
		osconsts.PathSeparator,
		expanded)
}
