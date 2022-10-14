package downloadinsexectest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func Test_Checksum(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tempFile, _ := createTempFile(t)

	// spin up the server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tempFile)
	}))

	defer ts.Close()

	download := &pathinsfmt.Download{
		Url:                  ts.URL,
		Destination:          fileDir,
		FileName:             file,
		IsSkipOnExist:        true,
		FileModeDir:          consts.DefaultDirectoryFileMode,
		ChecksumVerifyMethod: hashas.Md5,
		ChecksumVerify:       "469e01d115cb913ad709c749df1c5666",
	}

	// Act
	errWrap := downloadinsexec.Apply(download)

	// Assert
	convey.Convey("Download ErrorWrapper Should Return False", t, func() {
		convey.So(errWrap.HasError(), convey.ShouldBeFalse)
	})

	convey.Convey("Downloaded Content's Checksum Should Match", t, func() {
		fileCheckSum := checksummer.NewSync(true, filePath, download.ChecksumVerifyMethod)

		convey.So(fileCheckSum.SingleHash(), convey.ShouldEqual, download.ChecksumVerify)
	})
}
