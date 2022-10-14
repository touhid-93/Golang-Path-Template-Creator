package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getMimeTypesPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetMimeTypes",
	expected: "/etc/nginx/mime.types",
}

func TestGetMimeTypes(t *testing.T) {
	getPathTestCommonMethodLinux(t, getMimeTypesPathTestCaseData, nginxlinuxpath.GetMimeTypes)
}
