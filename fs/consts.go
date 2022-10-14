package fs

import (
	"os"
)

const (
	FlagAppend                = os.O_APPEND
	FlagWrite                 = os.O_WRONLY
	FlagCreate                = os.O_CREATE
	FlagReadOnly              = os.O_RDONLY
	FlagAppendOrWrite         = FlagAppend | FlagWrite
	FlagWriteOrCreate         = FlagWrite | FlagCreate
	FlagWriteOrCreateOrAppend = FlagWrite | FlagCreate | FlagAppend
)
