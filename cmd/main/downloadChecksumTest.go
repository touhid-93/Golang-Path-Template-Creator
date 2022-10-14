package main

import (
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func downloadChecksumTest() {
	download := &pathinsfmt.Download{
		// Url:                  "https://github.com/aria2/aria2/releases/download/release-1.35.0/aria2-1.35.0.tar.xz",
		Url:                  "https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh",
		Destination:          "/home/a/checksum2",
		FileName:             "install.sh",
		IsSkipOnExist:        false,
		IsCreateDir:          true,
		FileModeDir:          consts.DefaultDirectoryFileMode,
		ChecksumVerifyMethod: hashas.Sha256,
		// ChecksumVerify:       "1e2b7fd08d6af228856e51c07173cfcf987528f1ac97e04c5af4a47642617dfd",
		ChecksumVerify: "b6af836b2662f21081091e0bd851d92b2507abb94ece340b663db7e4019f8c7c",
	}

	errWrap := downloadinsexec.Apply(download)
	errWrap.HandleError()
}
