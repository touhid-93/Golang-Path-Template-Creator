package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/fileinfopath"
)

func fileInfoWithPathTest01() {
	files := pkgRootFiles()
	collection := fileinfopath.NewInstanceCollectionUsingFilePathsAsync(files...)

	fmt.Println(collection.First().String())
	fmt.Println(collection)
}
