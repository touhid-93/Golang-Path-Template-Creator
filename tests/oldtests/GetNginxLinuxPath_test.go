package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var getNginxLinuxTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "NginxLinuxPath",
	expected: "/etc/nginx/",
}

func TestGetNginxLinuxPath(t *testing.T) {
	getPathTestCommonMethodLinux(t, getNginxLinuxTestCaseData, knowndirget.NginxLinuxPath)
}
