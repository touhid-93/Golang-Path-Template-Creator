![](https://gitlab.com/evatix-go/pathhelper/uploads/6d228f94832193e553ddbc24401f7a52/image.png)

# pathhelper

path helper utility tool

## Git Clone

`git clone https://gitlab.com/evatix-go/pathhelper.git`

### Prerequisites

- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/evatix-go/pathhelper`

## Why *pathhelper*?

Package pathhelper provides an easy and fast way to get your desired OS(operating system) functionality without the
hassle of considering the OS you are on, what GO packages you may need to decode a path or simply find out if the path
or file exists etc. We have brought features of different packages and injected some new features to make this package a
complete solution for obtaining information regarding filepath independent of platform.

## Examples

```go

func main() {
// Checking if path is empty
ispath.Empty("") // returns true

samplePath := "C:\\users\\"

// Checking if path exists
exists := ispath.Exists(samplePath) // .File / .Directory / .AnyExists ...
fmt.Println(exists) // returns true if directory or file exist on that path

// Getting path as URI
fmt.Println(pathhelper.GetPathAsUri(samplePath, true)) // file:///c:/users

// Normalize path
pathToNormalize := "file:///C:/something/otherthing"
fmt.Println(normalize.Path(pathToNormalize)) // C:\something\otherthing if OS is windows; C:/something/otherthing if OS is Unix

// Create Directory
createdir.New("SampleNewDir", filemodes.AllRead) // Creates a directory named "SampleNewDir" with AllRead access at the folder where the function is called

// Environmental variables
envpath.GetExecutableEnvironmentPathCollection() // Outputs a struct containing the environment variable paths

// Check directory validity
fmt.Println(dirinfo.New("c:\\windows\\py.exe").IsValidDir) // outputs a bool

// Gets all paths recursively and continues on error as per provided argument
fmt.Println(recursivepaths.All("D:\\SampleDir", true)) // &[\\?\D:\SampleDir \\?\D:\SampleDir\sampleFile.txt] # Error Wrappers - Collection - Length[0]
fmt.Println(recursivepaths.All(" ", false)) // panics with detailed message

// Getting path as wrapper
fmt.Println(unipath.New("\\").Add("c://windows").Add("\\sys32").GetAsPathWrapper().String()) // \\?\c:\windows\sys32 on windows based on separator given
// "\\?\c:\windows\sys32\sys64\somethingnew" because "\\" separator given
// "c:/windows/sys32/sys64/somethingnew" because "/" separator given
fmt.Println(unipath.
New("\\").
Add("c://windows").
Add("\\sys32//").
Add("\\sys64//").
Add("somethingnew").
String()) // .ToString(...) Can change separator for new path

// "c:/windows/sys32/sys64", 
// "somethingnew" not printed because we skipped it
fmt.Println(unipath.
New("/").
Add("c://windows").
Add("\\sys32//").
Add("//sys64\\").
Add("somethingnew").
ToStringUptoLastMinus(1, "/", true))
}
```

## Acknowledgement

For this package we have mainly used OS and filepath packages of GO. For testing the package we have used the very
convenient package *[Go Convey](http://goconvey.co/)*.

## Links

- [What are conventions for filenames in Go? - Stack Overflow](https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go)
- [go - Pass method argument to function - Stack Overflow](https://stackoverflow.com/questions/38897529/pass-method-argument-to-function)
- [exec.Command() in Go with environment variable - Stack Overflow](https://stackoverflow.com/questions/51015569/exec-command-in-go-with-environment-variable)
- [go - Pointers vs. values in parameters and return values - Stack Overflow](https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values?rq=1)
- [jmhodges/copyfighter: Statically analyzes Go code and reports functions that are passing large structs by value](https://github.com/jmhodges/copyfighter)
- [CodeReviewComments · golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments#pass-values)
- [Difference between := and = operators in Go - Stack Overflow](https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go?rq=1)
- [How to manage Long Paths in Bash? - Stack Overflow](https://stackoverflow.com/questions/670488/how-to-manage-long-paths-in-bash)
- [Go import cycle issue fix](https://stackoverflow.com/questions/16168601/any-good-advice-about-how-to-avoid-import-cycle-in-go)
- [Samples for path and filename extraction](https://play.golang.org/p/oT6eWNZAeEi)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)