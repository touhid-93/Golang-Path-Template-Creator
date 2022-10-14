package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getModulesEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModulesEnabled",
	expected: "/etc/nginx/modules-enabled",
}

func TestGetModulesEnabled(t *testing.T) {
	getPathTestCommonMethodLinux(t, getModulesEnabledPathTestCaseData, nginxlinuxpath.GetModulesEnabled)
}
