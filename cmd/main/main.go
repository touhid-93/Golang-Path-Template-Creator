package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathcompiler"
)

func main() {
	var sampleFilePath string = "abc.txt"
	var pathToNormalize string = pathcompiler.MyOsTest.JoinWithTempRoot(sampleFilePath)
	fmt.Println(normalize.Path(pathToNormalize))

	// fmt.Println("Hello World")

	// downloadChecksumTest()
	// options := normalize.Options{
	// 	IsNormalize:        true,
	// 	IsLongPathFix:      true,
	// 	IsForceLongPathFix: true,
	// }
	//
	// samplePath := options.JoinWithBaseDirPaths(pathjoin.WithTemp(), "basedir", "something")
	// samplePath2 := options.JoinWithBaseDirPaths(samplePath, "basedir", "something")
	//
	// sha1 := hashas.Sha1
	//
	// rs := sha1.HexOfJsonResult(corejson.New(samplePath2))
	//
	// fmt.Println(rs.Value)
	//
	// fmt.Println(samplePath)
	// fmt.Println(samplePath2)

	// sha1 := hashas.Sha1
	//
	// slice1 := []string{
	// 	"alim1",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// slice2 := []string{
	// 	"alim1",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// slice3 := []string{
	// 	"alim5",
	// 	"alim2",
	// 	"alim3",
	// }
	//
	// rs2 := sha1.HexSumOfAnyItems(slice1, slice2, slice3)
	//
	// fmt.Println(rs2.String())
	//
	// rs3 := sha1.HexSumOfAnyItemsToCombinedSingleString(slice1, slice2, slice3)
	//
	// fmt.Println(rs3.String())
	//
	// files := recursivepaths.Files(pathsconst.RootDir)
	// fmt.Println("Files", files.String())
	//
	// rs5 := hexchecksum.OfFiles(&hexchecksum.FilesRequest{
	// 	Method:                     sha1,
	// 	IsGenerateContentsChecksum: true,
	// 	Files:                      files.Values,
	// })
	//
	// fmt.Println(converters.AnyToFullNameValueString(rs5))
	// fmt.Println(converters.AnyToFullNameValueString(rs5.HexChecksumOfResult()))

	// checkSumCheck()

	// pathTest()
	// fileStateTest01()
	// fileStateTest02()
	// filterTest02()
	// filterTest03()
	// filterTest04()

	// checksumTest01()
	// checksumTest02()
	// checksumTest03()
	// fileInfoWithPathTest01()
	// fileInfoWithPathTest02()
	// checksumTest04()
	// nginxPathTest01()

	// testPathWithVerifier()
	// hashsetReadWriteCache01()
	// mapStringAnyReadWriteCache01()
	// hashmapReadWriteTest02()
}
