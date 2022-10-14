package pathinsfmtexectestwrappers

import (
	"gitlab.com/evatix-go/errorwrapper/errverify"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

type PathWithModifyAndVerifyWrapper struct {
	Header      string
	Modifier    *pathinsfmt.PathWithModifier
	Verifier    *pathinsfmt.PathWithVerifier
	ErrorVerify *errverify.CollectionVerifier
}
