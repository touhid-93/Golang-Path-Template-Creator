package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type Writer struct {
	ChmodWrapper            pathchmod.Wrapper
	Location                string
	IsApplyChmodMust        bool
	IsApplyChmodOnMismatch  bool // or else apply all the time
	IsSkipNullObject        bool
	IsWriteEmptyOnNull      bool
	IsKeepExistingFileChmod bool
}

func (it *Writer) WriteLock(
	rawBytes []byte,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.Write(rawBytes)
}

func (it *Writer) Write(
	rawBytes []byte,
) *errorwrapper.Wrapper {
	if it.IsWriteEmptyOnNull && rawBytes == nil {
		return WriteAllParams(
			true,
			it.IsApplyChmodMust,
			it.IsApplyChmodOnMismatch,
			it.IsSkipNullObject,
			it.IsKeepExistingFileChmod,
			it.ChmodWrapper.DirChmod,
			it.ChmodWrapper.FileChmod,
			it.Location,
			[]byte(""))
	}

	if it.IsSkipNullObject && rawBytes == nil {
		return nil
	}

	return WriteAllParams(
		true,
		it.IsApplyChmodMust,
		it.IsApplyChmodOnMismatch,
		it.IsSkipNullObject,
		it.IsKeepExistingFileChmod,
		it.ChmodWrapper.DirChmod,
		it.ChmodWrapper.FileChmod,
		it.Location,
		rawBytes)
}

func (it *Writer) WriteString(
	content string,
) *errorwrapper.Wrapper {
	return it.Write([]byte(content))
}

func (it *Writer) WriteStringLock(
	content string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteString(content)
}

func (it *Writer) WriteStringEofLine(
	content string,
) *errorwrapper.Wrapper {
	eofFix := it.appendEofLine(
		true,
		content)

	return it.Write([]byte(eofFix))
}

func (it *Writer) WriteStringEofLineLock(
	content string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteStringEofLine(content)
}

func (it *Writer) WriteStrings(
	isAppendEof bool, // adds new line at the end
	lines ...string,
) *errorwrapper.Wrapper {
	line := strings.Join(
		lines,
		constants.DefaultLine)
	line = it.appendEofLine(
		isAppendEof,
		line)

	return it.WriteString(line)
}

func (it *Writer) WriteLines(
	isAppendEof bool, // adds new line at the end
	lines ...string,
) *errorwrapper.Wrapper {
	line := strings.Join(
		lines,
		constants.DefaultLine)
	line = it.appendEofLine(
		isAppendEof,
		line)

	return it.WriteString(line)
}

func (it *Writer) WriteErrBytes(
	errBytes *errbyte.Results,
) *errorwrapper.Wrapper {
	if errBytes == nil {
		return errnew.Null.Simple(errBytes)
	}

	errWrap := errBytes.ErrWrap()

	if errWrap.HasAnyError() {
		return errWrap
	}

	return it.Write(errBytes.SafeValues())
}

func (it *Writer) WriteErrBytesLock(
	errBytes *errbyte.Results,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteErrBytes(errBytes)
}

func (it *Writer) WriteJsonResult(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	errWrap := errnew.Json.Result(jsonResult)

	if errWrap.HasAnyError() {
		return errWrap
	}

	return it.Write(jsonResult.SafeValues())
}

func (it *Writer) WriteJsonResultLock(
	jsonResult *corejson.Result,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteJsonResult(jsonResult)
}

func (it *Writer) WriteLinesLock(
	isAppendEof bool, // adds new line at the end
	lines ...string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return it.WriteLines(isAppendEof, lines...)
}

func (it *Writer) WriteLinesEofLine(
	lines []string,
) *errorwrapper.Wrapper {
	return it.WriteLines(
		true,
		lines...)
}

func (it *Writer) appendEofLine(isAppendEof bool, line string) string {
	if !isAppendEof {
		return line
	}

	if strings.HasSuffix(line, constants.DefaultLine) {
		return line
	}

	return line + constants.DefaultLine
}

func (it *Writer) IsEmptyLocation() bool {
	return it == nil || it.Location == ""
}

func (it *Writer) IsPathInvalid() bool {
	return chmodhelper.IsPathInvalid(it.Location)
}

func (it *Writer) IsPathExists() bool {
	return chmodhelper.IsPathExists(it.Location)
}
