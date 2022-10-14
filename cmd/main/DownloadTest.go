package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func DownloadTest() {
	ins := &pathinsfmt.Download{
		Url:              "https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh",
		Destination:      "/home/a/dtestxxxx",
		FileName:         "installx.sh",
		IsCreateDir:      true,
		IsClearDir:       true,
		ParallelRequests: 4,
		MaxRetries:       5,
		FileModeDir:      filemode.X666,
	}

	errWrap := downloadinsexec.Apply(ins)
	fmt.Println(errWrap)
}
