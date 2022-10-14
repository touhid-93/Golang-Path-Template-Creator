package pathinsfmtexectestwrappers

import "gitlab.com/evatix-go/pathhelper/pathjoin"

var (
	PathOneTextFile = pathjoin.WithTempTest("pathone/abc.txt")
	PathTwoTextFile = pathjoin.WithTempTest("pathtwo/abc.txt")
	NotExistPath    = pathjoin.WithTempTest("nopath/no.txt")
)
