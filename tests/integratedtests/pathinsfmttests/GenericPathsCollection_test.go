package pathinsfmttests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Test_GenericPathsCollection(t *testing.T) {
	// Arrange
	genPaths := pathinsfmt.GenericPathsCollection{
		Specification: nil,
		SimilarPaths: []pathinsfmt.SimilarPaths{
			{
				RootPath: "Similar",
				RelativePaths: []string{
					"rel path",
					"rel/ dw ////wdwdws path2",
				},
				IsNormalizeApply: true,
			},
		},
		AllDiffPaths: []pathinsfmt.AllDiffPaths{
			{
				Paths: []string{
					"all Diff / path",
					"all Diff path2",
				},
				IsNormalizeApply: true,
			},
		},
		DynamicPaths: &pathinsfmt.DynamicPaths{
			Vars: nil,
			AllDiffPaths: []pathinsfmt.AllDiffPaths{
				{
					Paths: []string{
						"dynamic path 1",
						"dynamic path 2",
						"dynamic path 3",
						"dynamic path 4",
						"dynamic path 5",
						"dynamic / path 6",
						"dynamic path 7",
					},
					IsNormalizeApply: true,
				},
			},
		},
	}
	expectation := "Similar\\rel path|" +
		"Similar\\rel\\ dw \\wdwdws path2|" +
		"all Diff \\ path|all Diff path2|" +
		"dynamic \\ path 6|dynamic path 1|" +
		"dynamic path 2|" +
		"dynamic path 3|" +
		"dynamic path 4|" +
		"dynamic path 5|" +
		"dynamic path 7"

	// Act
	allPaths := genPaths.LazyFlatPathsSorted()
	actual := strings.Join(allPaths, "|")

	// Assert
	Convey("Testing Generic Paths MappedInfoItems FlatPaths Method", t, func(c C) {
		So(actual, ShouldEqual, expectation)
	})
}

func Test_GenericPathsCollectionWhereOthersAreNil(t *testing.T) {
	// Arrange
	genPaths := pathinsfmt.GenericPathsCollection{
		Specification: nil,
		AllDiffPaths: []pathinsfmt.AllDiffPaths{
			{
				Paths: []string{
					"all Diff / path",
					"all Diff path2",
				},
				IsNormalizeApply: true,
			},
		},
	}
	expectation := "all Diff \\ path|" +
		"all Diff path2"

	// Act
	allPaths := genPaths.LazyFlatPathsSorted()
	actual := strings.Join(allPaths, "|")

	// Assert
	Convey("Testing Generic Paths MappedInfoItems FlatPaths Method while others are empty", t, func(c C) {
		So(actual, ShouldEqual, expectation)
	})
}
