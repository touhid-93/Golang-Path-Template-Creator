package main

import (
	"fmt"
	"path/filepath"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func pathTest() {
	a := "/tmp/dbapi/backup-storage//path-backup///alim-key1-dbapi/1/\\"
	b := "/dbmodel/\\webserverstoremodel/\\\\ServerWithSSL.go"
	joined3 := pathjoin.JoinConditionalNormalizedThreeExpandIf(
		true,
		false,
		a,
		b,
		"")
	fmt.Println(filepath.Clean(joined3))
	fmt.Println(filepath.Join(a, b))

	// joined := filepath.Join(a, b)
	fmt.Println(filepath.Clean(a))

	fmt.Println(normalize.Path("\\\\?\\tmp\\dbapi\\backup-storage\\path-backup\\alim-key1-dbapi\\1\\dbmodel\\webserverstoremodel\\ServerWithSSL.go"))
	fmt.Println(normalize.PathUsingSeparatorIf(true, true, true, osconsts.PathSeparator, "tmp\\dbapi\\backup-storage\\path-backup\\alim-key1-dbapi\\1\\dbmodel\\webserverstoremodel\\ServerWithSSL.go"))

	fmt.Println(normalize.Path("/home/a/../../git-repos"))
	fmt.Println(normalize.Path("\\home\\a\\../../git-repos"))
}
