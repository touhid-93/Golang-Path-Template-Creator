package consts

import (
	"gitlab.com/evatix-go/core/filemode"
)

const (
	FilePathEmpty                     = "File path was empty(\"\")."
	BrokenLongPathUncPrefix           = `\?\UNC\`
	BrokenLongPathQuestionMarkPrefix  = `\?\`
	DoubleStars                       = "**"
	Export                            = "export"
	SpaceHyphenRightAngelBracketSpace = " -> "
	DefaultFileMode                   = filemode.FileDefault // cannot execute by everyone OwnerCanReadWriteGroupOtherCanReadOnly
	DefaultDirMode                    = filemode.DirDefault  // can execute by everyone OwnerCanDoAllExecuteGroupOtherCanReadExecute
	DefaultDirectoryFileMode          = filemode.X666
	NonAsyncSafeRange                 = 50
	FileInfoEachLineJoiner            = ",\n - "
	IndentFileInfoEachLineJoiner      = "\n   - "
)
