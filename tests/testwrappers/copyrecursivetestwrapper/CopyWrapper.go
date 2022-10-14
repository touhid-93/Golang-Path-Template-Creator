package copyrecursivetestwrapper

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/errorwrapper/errverify"
	"gitlab.com/evatix-go/pathhelper/copyrecursive"
)

type CopyRecursiveTestWrapper struct {
	Header                                        string
	IsClearDestinationPre, IsClearDestinationPost bool
	CopyInstruction                               copyrecursive.Instruction
	ErrorCollectionVerifier                       errverify.CollectionVerifier
}

var CopyRecursiveTestCases = []CopyRecursiveTestWrapper{
	{
		Header:                 "",
		IsClearDestinationPre:  true,
		IsClearDestinationPost: true,
		CopyInstruction: copyrecursive.Instruction{
			SourceDestination: copyrecursive.SourceDestination{
				Source:      SourceRecursivePath,
				Destination: Destination,
			},
			Options: copyrecursive.Options{
				IsSkipOnExist:      false,
				IsRecursive:        true,
				IsClearDestination: false,
				IsUseShellOrCmd:    false,
				IsNormalize:        false,
				IsExpandVar:        false,
			},
		},
		ErrorCollectionVerifier: errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header:       "",
				FunctionName: "TestCopierRecursive_3",
				VerifyAs:     stringcompareas.Equal,
				IsPrintError: true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				Items: []string{},
			},
			ErrorLength: 0,
		},
	},
}
