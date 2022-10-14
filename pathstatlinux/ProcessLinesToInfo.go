package pathstatlinux

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
)

/*	ProcessLinesToInfo processLinesToInfo
➜  ~ stat /etc/mysql
	File: /etc/mysql
	Size: 4096            Blocks: 8          IO Block: 4096   directory
	Device: 10302h/66306d   Inode: 6293381     Links: 4
	Access: (0755/drwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
	Access: 2021-06-03 09:18:30.014302041 +0000
	Modify: 2021-06-03 09:18:32.086923601 +0000
	Change: 2021-06-03 09:18:32.086923601 +0000
*/
func ProcessLinesToInfo(
	lines []string,
	filePath string,
	errorWrapper *errorwrapper.Wrapper,
) *Info {
	accessLineIndex := lines[coreindexes.Fourth]
	accessTime := getTimePartAsString(lines[coreindexes.Fifth])
	modifyTime := getTimePartAsString(lines[coreindexes.Sixth])
	changeTime := getTimePartAsString(lines[coreindexes.Seventh])

	splits := strings.Split(accessLineIndex, constants.Colon)
	rwxSimple := getRwxSimpleFromPathStatLines(splits, filePath)
	user := getUserFromPathStatLines(splits)
	group := getGroupFromPathStatLines(splits)
	isValid := rwxSimple.IsRwxValid &&
		user.HasValidId &&
		group.HasValidId

	isDir, fileInfo := ispathinternal.DirectoryPlusFileInfo(
		filePath)

	isFileExist := isDir || fileInfo != nil || ispathinternal.Exists(
		filePath)

	return &Info{
		User:      *user,
		Group:     *group,
		RwxSimple: *rwxSimple,
		RawLocationTimestamp: RawLocationTimestamp{
			Access: accessTime,
			Modify: modifyTime,
			Change: changeTime,
		},
		Location:       filePath,
		IsPathExist:    isFileExist,
		IsDirectory:    isDir,
		IsValidParsing: isValid,
		ErrorWrapper:   errorWrapper,
		FileInfo:       fileInfo,
	}
}
