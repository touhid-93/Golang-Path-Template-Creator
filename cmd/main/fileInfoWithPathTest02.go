package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func fileInfoWithPathTest02() {
	files := pkgRootFiles()
	collection := fileinfopath.NewInstanceCollectionUsingFilePathsAsync(files...)

	fmt.Println(collection.JsonString())
}
