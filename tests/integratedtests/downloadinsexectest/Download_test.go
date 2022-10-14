package downloadinsexectest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func Test_Download(t *testing.T) {
	coretests.SkipOnWindows(t)

	// Arrange
	tempFile, buff := createTempFile(t)

	// spin up the server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tempFile)
	}))

	defer ts.Close()

	// Act
	download := pathinsfmt.NewDownload(ts.URL, filePath)
	errWrap := downloadinsexec.Apply(download)
	errWrap.HandleError()

	// Assert
	convey.Convey("Download ErrorWrapper Should Return False", t, func() {
		convey.So(errWrap.HasError(), convey.ShouldBeFalse)
	})

	convey.Convey("Download Content Should Resemble Temp Content", t, func() {
		errBytesResults := fs.ReadFile(filePath)
		convey.So(errBytesResults.HasError(), convey.ShouldBeFalse)
		convey.So(errBytesResults.SafeValues(), convey.ShouldResemble, buff)
	})
}
