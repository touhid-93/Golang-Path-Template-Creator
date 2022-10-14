package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

var getConfPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConf",
	expected: "/etc/nginx/conf.d",
}

func TestGetConf(t *testing.T) {
	getPathTestCommonMethodLinux(t, getConfPathTestCaseData, nginxlinuxpath.GetConf)
}
