package pathinsfmtexectests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers"
)

func Test_ApplyPathWithVerifiersErrorsUnix(t *testing.T) {
	coretests.SkipOnWindows(t)
	// 0. Setup
	locations := testwrappers.SetupDefaultPathsUnix()

	// 1. Arrange
	pathWithVerifiers := make([]pathinsfmt.PathWithVerifier, len(locations))
	for i := range locations {
		ins := pathinsfmt.PathWithVerifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          locations[i],
				IsNormalize:   true,
				IsRecursive:   true,
				IsSkipInvalid: false,
			},
			Verifier: &pathinsfmt.PathVerifier{
				UserGroupName: *testwrappers.DefaultUserNameGroupName,
				BaseRwxInstructions: chmodins.BaseRwxInstructions{
					RwxInstructions: []chmodins.RwxInstruction{
						{
							RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
						},
					},
				},
			},
		}

		pathWithVerifiers[i] = ins
	}

	pathsWithVerifiers := &pathinsfmt.PathsWithVerifiers{
		IsContinueOnError: true,
		Verifiers:         pathWithVerifiers,
	}

	errColl := errwrappers.Empty()

	// Act
	isSuccess := pathmodifierverify.ApplyPathWithVerifiers(
		true,
		errColl,
		pathsWithVerifiers,
	)

	// Assert
	Convey("All paths should be verified", t, func() {
		So(errColl.String(), ShouldBeEmpty)
		So(errColl.IsEmpty(), ShouldBeTrue)
		So(isSuccess, ShouldBeTrue)
	})
}

func Test_ApplyPathWithVerifierErrorsUnix(t *testing.T) {
	coretests.SkipOnWindows(t)
	// 0. Setup
	locations := testwrappers.SetupDefaultPathsUnix()
	for i := range locations {
		// 1. Arrange
		ins := &pathinsfmt.PathWithVerifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          locations[i],
				IsNormalize:   true,
				IsRecursive:   true,
				IsSkipInvalid: false,
			},
			Verifier: &pathinsfmt.PathVerifier{
				UserGroupName: *testwrappers.DefaultUserNameGroupName,
				BaseRwxInstructions: chmodins.BaseRwxInstructions{
					RwxInstructions: []chmodins.RwxInstruction{
						{
							RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
							Condition: chmodins.Condition{
								IsSkipOnInvalid:   false,
								IsContinueOnError: false,
								IsRecursive:       false,
							},
						},
					},
				},
			},
		}

		errColl := errwrappers.Empty()
		// Act
		isSuccess := pathmodifierverify.ApplyPathWithVerifier(
			true,
			errColl,
			ins,
		)

		// Assert
		Convey("Location : "+locations[i]+" should be verified", t, func() {
			So(errColl.String(), ShouldBeEmpty)
			So(errColl.IsEmpty(), ShouldBeTrue)
			So(isSuccess, ShouldBeTrue)
		})
	}
}

func Test_ApplyPathVerifierErrorsUnix(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	locations := testwrappers.SetupDefaultPathsUnix()
	verifier := &pathinsfmt.PathVerifier{
		UserGroupName: *testwrappers.DefaultUserNameGroupName,
		BaseRwxInstructions: chmodins.BaseRwxInstructions{
			RwxInstructions: []chmodins.RwxInstruction{
				{
					RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
					Condition: chmodins.Condition{
						IsSkipOnInvalid:   false,
						IsContinueOnError: false,
						IsRecursive:       false,
					},
				},
			},
		},
	}
	errCollection := errwrappers.Empty()

	// Act
	isSuccess := pathmodifierverify.ApplyVerifier(true,
		true,
		false,
		true,
		verifier,
		errCollection,
		locations)

	// Assert
	Convey("Default Paths Create", t, func() {
		So(errCollection.String(), ShouldBeEmpty)
		So(errCollection.IsEmpty(), ShouldBeTrue)
		So(isSuccess, ShouldBeTrue)
	})
}

func Test_ApplyPathVerifierErrorsWindows(t *testing.T) {
	// Arrange
	locations := testwrappers.SetupDefaultPathsUnix()
	verifier := &pathinsfmt.PathVerifier{
		UserGroupName: pathinsfmt.UserGroupName{},
		BaseRwxInstructions: chmodins.BaseRwxInstructions{
			RwxInstructions: []chmodins.RwxInstruction{
				{
					RwxOwnerGroupOther: *testwrappers.DefaultRwxOwnerGroupOther,
					Condition: chmodins.Condition{
						IsSkipOnInvalid:   false,
						IsContinueOnError: false,
						IsRecursive:       false,
					},
				},
			},
		},
	}

	errCollection := errwrappers.Empty()

	// Act
	isSuccess := pathmodifierverify.ApplyVerifier(
		true,
		true,
		false,
		true,
		verifier,
		errCollection,
		locations)

	// Assert
	Convey("Default Paths Create", t, func() {
		So(errCollection.String(), ShouldBeEmpty)
		So(errCollection.IsEmpty(), ShouldBeTrue)
		So(isSuccess, ShouldBeTrue)
	})
}
