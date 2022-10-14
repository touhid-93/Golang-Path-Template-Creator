package main

import (
	"gitlab.com/evatix-go/pathhelper/pathsconst"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

func pkgRootFiles() []string {
	return recursivepaths.Files(pathsconst.RootDir).Values
}
