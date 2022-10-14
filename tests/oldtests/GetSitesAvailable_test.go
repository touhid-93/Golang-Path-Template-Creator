package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getSitesAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetSitesAvailable",
	expected: "/etc/nginx/sites-available",
}

func TestGetSitesAvailable(t *testing.T) {
	getPathTestCommonMethodLinux(
		t,
		getSitesAvailablePathTestCaseData,
		nginxlinuxpath.GetSitesAvailable,
	)
}
