package pathinsfmtexectestwrappers

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/errorwrapper/errverify"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers"
)

var PathWithModifyAndVerifyTestCases = []PathWithModifyAndVerifyWrapper{
	// Case 1 ===================>
	{
		Header: "Given Path Exists Modifying And Verifying Path With Exact Modification Parameters\n" +
			"Should Have Success And No Error",
		Modifier: &pathinsfmt.PathWithModifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          PathOneTextFile,
				IsNormalize:   true,
				IsRecursive:   true,
				IsSkipInvalid: false,
			},
			Modifier: &pathinsfmt.PathModifier{
				Chown: &pathinsfmt.Chown{
					BaseIsRecursive: pathinsfmt.BaseIsRecursive{
						IsRecursive: true,
					},
					UserGroupName: *testwrappers.DefaultUserNameGroupName,
				},
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
		},
		Verifier: &pathinsfmt.PathWithVerifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          PathOneTextFile,
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
		},
		ErrorVerify: &errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header: "Given Path Exists Modifying And Verifying Path With Exact Modification Parameters\n" +
					"Should Have Success And No Error",
				FunctionName:             "Test_PathModifierAndVerifierErrors",
				VerifyAs:                 stringcompareas.Equal,
				IsCompareEmpty:           false,
				IsVerifyErrorMessageOnly: false,
				IsPrintError:             true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				Items: []string{},
			},
			ErrorLength: 0,
		},
	},
	// Case 2 ===================>
	{
		Header: "Given Path Does Not Exist Modifying And Verifying Path With Exact Modification Parameters\n" +
			"Should Fail And Have Errors",
		Modifier: &pathinsfmt.PathWithModifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          NotExistPath,
				IsNormalize:   true,
				IsRecursive:   true,
				IsSkipInvalid: false,
			},
			Modifier: &pathinsfmt.PathModifier{
				Chown: &pathinsfmt.Chown{
					BaseIsRecursive: pathinsfmt.BaseIsRecursive{
						IsRecursive: true,
					},
					UserGroupName: *testwrappers.DefaultUserNameGroupName,
				},
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
		},
		Verifier: &pathinsfmt.PathWithVerifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          NotExistPath,
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
		},
		ErrorVerify: &errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header: "Given Path Does Not Exist Modifying And Verifying Path With Exact Modification Parameters\n" +
					"Should Fail And Have Errors",
				FunctionName:             "Test_PathModifierAndVerifierErrors",
				VerifyAs:                 stringcompareas.Equal,
				IsCompareEmpty:           false,
				IsVerifyErrorMessageOnly: false,
				IsPrintError:             true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				Items: []string{
					"[Error (MissingPathsOrInvalidPaths - #317): Missing path(s) or invalid path(s)! Additional : stat /tmp/pkg-testing//nopath/no.txt: no such file or directory. Ref(s) {[File Path (string): \"/tmp/pkg-testing//nopath/no.txt\"]}]",
					"[Error (FileExpand - #532): File expand failed! Additional : [Error (MissingPathsOrInvalidPaths - #317): Missing path(s) or invalid path(s)! Additional : stat /tmp/pkg-testing//nopath/no.txt: no such file or directory. Ref(s) {[File Path (string): \"/tmp/pkg-testing//nopath/no.txt\"]}]. Ref(s) {[File Path (string): \"/tmp/pkg-testing//nopath/no.txt\"]}]",
				},
			},
			ErrorLength: 2,
		},
	},
	// Case 3 ===================>
	{
		Header: "Given Path Exists Modifying First Then Verifying Path With Different Parameters\n" +
			"Should Fail And Have Errors",
		Modifier: &pathinsfmt.PathWithModifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          PathOneTextFile,
				IsNormalize:   true,
				IsRecursive:   true,
				IsSkipInvalid: false,
			},
			Modifier: &pathinsfmt.PathModifier{
				Chown: &pathinsfmt.Chown{
					BaseIsRecursive: pathinsfmt.BaseIsRecursive{
						IsRecursive: true,
					},
					UserGroupName: *testwrappers.DefaultUserNameGroupName,
				},
				BaseRwxInstructions: chmodins.BaseRwxInstructions{
					RwxInstructions: []chmodins.RwxInstruction{
						{
							RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
								Owner: "rwx",
								Group: "rwx",
								Other: "rwx",
							},
							Condition: chmodins.Condition{
								IsSkipOnInvalid:   false,
								IsContinueOnError: false,
								IsRecursive:       false,
							},
						},
					},
				},
			},
		},
		Verifier: &pathinsfmt.PathWithVerifier{
			PathWithOptions: pathinsfmt.PathWithOptions{
				Path:          PathOneTextFile,
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
		},
		ErrorVerify: &errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header: "Given Path Exists Modifying First Then Verifying Path With Different Parameters\n" +
					"Should Fail And Have Errors",
				FunctionName:             "Test_PathModifierAndVerifierErrors",
				VerifyAs:                 stringcompareas.Equal,
				IsCompareEmpty:           false,
				IsVerifyErrorMessageOnly: false,
				IsPrintError:             true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				Items: []string{
					"[Error (RwxMismatch - #304): Rwx mismatch! Path:/tmp/pkg-testing//pathone/abc.txt - Expect [\"rwxr-xr-x\"] != [\"rwxrwxrwx\"] Actual.]",
				},
			},
			ErrorLength: 1,
		},
	},
}
