package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getModulesAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesAvailable",
	expected: "/etc/nginx/modules-available",
}

func TestGetModulesAvailable(t *testing.T) {
	getPathTestCommonMethodLinux(t, getModulesAvailablePathTestCaseData, nginxlinuxpath.GetModulesAvailable)
}
