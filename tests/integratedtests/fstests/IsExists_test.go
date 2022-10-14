package fstests

import (
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/simplewrap"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/ispath"
	"gitlab.com/evatix-go/pathhelper/ispaths"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers/fstestwrapper"
)

func Test_IsExists(t *testing.T) {
	// Setup
	setupErr := fstestwrapper.SetupFiles.SetupDefault()

	setupErr.HandleError()

	for _, testCase := range fstestwrapper.ExistPathTestCases {
		location := testCase.Options.JoinWithBaseDirPaths(
			testCase.BaseDir,
			testCase.RelPath)

		locationQuotation := simplewrap.WithDoubleQuote(location)
		finalHeader := testCase.Header +
			constants.Hyphen +
			locationQuotation +
			constants.Hyphen

		fileInfo, err := os.Stat(location)

		isActualExist := fs.IsPathExists(location)

		isActualExist2 := fs.IsPathExistsUsingLock(location)
		isActualExist3 := ispath.Exists(location)
		isActualExist4 := ispaths.Exist(location)[0]
		isActualExist5 := chmodhelper.IsPathExists(location)

		isActualNotExist := fs.IsNotPathExists(location)
		isActualNotExist2 := fs.IsNotPathExistsUsing(fileInfo, err)
		isActualNotExist3 := ispath.NotExists(location)

		isActualExistFile := ispaths.AllFiles(location)
		isActualExistDir := fs.IsExistButDirectory(location)
		isActualExistDir2 := ispaths.AllDirectories(location)

		convey.Convey(finalHeader+"isActualExist as expected (fs.IsPathExists)",
			t, func() {
				convey.So(
					isActualExist,
					convey.ShouldEqual,
					testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExist2 as expected (fs.IsPathExistsUsingLock)",
			t, func() {
				convey.So(
					isActualExist2,
					convey.ShouldEqual,
					testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExist3 as expected (ispath.Exists)",
			t, func() {
				convey.So(
					isActualExist3,
					convey.ShouldEqual,
					testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExist4 as expected (ispaths.Exist(location)[0])",
			t, func() {
				convey.So(
					isActualExist4,
					convey.ShouldEqual,
					testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExist5 as expected (chmodhelper.IsPathExists)",
			t, func() {
				convey.So(
					isActualExist5,
					convey.ShouldEqual,
					testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualNotExist as expected (fs.IsNotPathExists)",
			t, func() {
				convey.So(
					isActualNotExist,
					convey.ShouldEqual,
					!testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualNotExist2 as expected (fs.IsNotPathExistsUsing(fileInfo, err))",
			t, func() {
				convey.So(
					isActualNotExist2,
					convey.ShouldEqual,
					!testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualNotExist3 as expected (ispath.NotExists)",
			t, func() {
				convey.So(
					isActualNotExist3,
					convey.ShouldEqual,
					!testCase.IsExistExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExistFile as expected (ispaths.AllFiles)",
			t, func() {
				convey.So(
					isActualExistFile,
					convey.ShouldEqual,
					testCase.IsFileExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExistDir as expected (fs.IsExistButDirectory)",
			t, func() {
				convey.So(
					isActualExistDir,
					convey.ShouldEqual,
					testCase.IsDirExpectation,
				)
			})

		convey.Convey(finalHeader+"isActualExistDir2 as expected (ispaths.AllDirectories)",
			t, func() {
				convey.So(
					isActualExistDir2,
					convey.ShouldEqual,
					testCase.IsDirExpectation,
				)
			})
	}
}
