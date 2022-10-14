package fstestwrapper

import (
	"gitlab.com/evatix-go/core/ostype"
	"gitlab.com/evatix-go/pathhelper/expandnormalize"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

type ExistPathWrapper struct {
	Header           string
	BaseDir, RelPath string
	OsType           ostype.Variation
	expandnormalize.Options
	IsExistExpectation bool
	IsDirExpectation   bool
	IsFileExpectation  bool
}

var ExistPathTestCases = []ExistPathWrapper{
	// {
	// 	Header:  "",
	// 	BaseDir: os.TempPermanentDir(),
	// 	RelPath: "",
	// 	OsType:  ostype.Windows,
	// 	Options: expandnormalize.Options{
	// 		IsNormalize:    true,
	// 		IsExpandEnvVar: false,
	// 	},
	// 	IsExistExpectation: true,
	// 	IsDirExpectation:   true,
	// 	IsFileExpectation:  false,
	// },
	// {
	// 	Header:  "",
	// 	OsType:  ostype.Windows,
	// 	BaseDir: os.TempPermanentDir(),
	// 	RelPath: "/random-something/random-date/16-20-20-20/2",
	// 	Options: expandnormalize.Options{
	// 		IsNormalize:    true,
	// 		IsExpandEnvVar: false,
	// 	},
	// 	IsExistExpectation: false,
	// },
	// {
	// 	Header:  "",
	// 	OsType:  ostype.Windows,
	// 	BaseDir: "$temp",
	// 	RelPath: "/random-something/random-date/16-20-20-20/3",
	// 	Options: expandnormalize.Options{
	// 		IsNormalize:    true,
	// 		IsExpandEnvVar: true,
	// 	},
	// 	IsExistExpectation: false,
	// },
	{
		Header:  "File should exist, dir should be false",
		BaseDir: pathsconst.DefaultTempTestDir,
		RelPath: setupFilePath,
		OsType:  ostype.Windows,
		Options: expandnormalize.Options{
			Options: normalize.Options{
				IsNormalize:        true,
				IsLongPathFix:      true,
				IsForceLongPathFix: true,
			},
			IsExpandEnvVar: true,
		},
		IsExistExpectation: true,
		IsDirExpectation:   false,
		IsFileExpectation:  true,
	},
	{
		Header:  "Dir should exist, file should be false",
		BaseDir: pathsconst.DefaultTempTestDir,
		RelPath: "",
		OsType:  ostype.Windows,
		Options: expandnormalize.Options{
			Options: normalize.Options{
				IsNormalize:        true,
				IsLongPathFix:      true,
				IsForceLongPathFix: true,
			},
			IsExpandEnvVar: true,
		},
		IsExistExpectation: true,
		IsDirExpectation:   true,
		IsFileExpectation:  false,
	},
}
