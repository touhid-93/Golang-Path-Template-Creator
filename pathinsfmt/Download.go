package pathinsfmt

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/createdir"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

// Download Use aria2c
// Reference : https://aria2.github.io/manual/en/html/aria2c.html#options
type Download struct {
	Url                  string         `json:"Url,omitempty"`
	Destination          string         `json:"Destination,omitempty"`
	FileName             string         `json:"FileName,omitempty"`
	ParallelRequests     byte           `json:"ParallelRequests,omitempty"`
	MaxRetries           byte           `json:"MaxRetries,omitempty"`
	IsCreateDir          bool           `json:"IsCreateDir,omitempty"`
	IsClearDir           bool           `json:"IsClearDir,omitempty"`
	IsNormalizePath      bool           `json:"IsNormalizePath,omitempty"`
	IsSkipOnExist        bool           `json:"IsSkipOnExist,omitempty"`
	FileModeDir          os.FileMode    `json:"FileModeDir,omitempty"`
	ChecksumVerifyMethod hashas.Variant `json:"ChecksumVerifyMethod,omitempty"` // Md5, Sha1, Sha256
	ChecksumVerify       string         `json:"CheckSumVerify,omitempty"`       // only verify if given
	fixPath              *string
}

// NewDownload Example :
//  NewDownload("http://url.com/file.ext","path/to/file/file.ext")
func NewDownload(url, destination string) *Download {
	parentDir, fileName := fsinternal.GetDirFileName(destination)

	return &Download{
		Url:              url,
		Destination:      parentDir,
		FileName:         fileName,
		ParallelRequests: constants.Capacity2,
		MaxRetries:       constants.Capacity2,
		IsCreateDir:      true,
		IsClearDir:       true,
		IsNormalizePath:  true,
		IsSkipOnExist:    false,
		FileModeDir:      consts.DefaultDirectoryFileMode,
	}
}

func (it *Download) FullPath() string {
	if it == nil {
		return ""
	}

	if it.fixPath != nil {
		return *it.fixPath
	}

	fixedPath := pathjoin.JoinIf(
		it.IsNormalizePath,
		it.Destination,
		it.FileName,
	)

	it.fixPath = &fixedPath

	return *it.fixPath
}

func (it *Download) PathStat() *chmodhelper.PathExistStat {
	return chmodhelper.GetPathExistStat(it.FullPath())
}

func (it *Download) CreateDirInstruction() *createdir.Instruction {
	return &createdir.Instruction{
		Location:      it.Destination,
		FileMode:      it.FileModeDir,
		IsLock:        true,
		IsCreate:      it.IsCreateDir,
		IsSkipOnExist: it.IsSkipOnExist,
		IsRemoveAll:   it.IsClearDir,
	}
}
