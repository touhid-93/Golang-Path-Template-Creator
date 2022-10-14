package hexchecksum

import "gitlab.com/evatix-go/pathhelper/hashas"

type FilesCompareRequest struct {
	Method     hashas.Variant
	LeftFiles  []string
	RightFiles []string
}
