package downloadinsexectest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func Test_SkipOnExist(t *testing.T) {
	// Arrange
	tempFile, buff := createTempFile(t)

	// spin up the server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tempFile)
	}))

	defer ts.Close()
	writeErr := fs.WriteStringToFileUsingLock(
		true,
		filePath,
		"hello,world")
	writeErr.HandleError()

	download := &pathinsfmt.Download{
		Url:           ts.URL,
		Destination:   fileDir,
		FileName:      file,
		IsSkipOnExist: true,
		FileModeDir:   consts.DefaultDirectoryFileMode,
	}

	// Act
	errWrap := downloadinsexec.Apply(download)

	// Assert
	convey.Convey("Download ErrorWrapper Should Return False", t, func() {
		convey.So(errWrap.HasError(), convey.ShouldBeFalse)
	})

	convey.Convey("Download Content Should Not Resemble Temp Content", t, func() {
		errBytesResults := fs.ReadFile(filePath)
		convey.So(errBytesResults.HasError(), convey.ShouldBeFalse)
		convey.So(errBytesResults.Values, convey.ShouldNotResemble, buff)
	})
}
