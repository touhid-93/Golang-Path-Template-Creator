package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

var getApacheLinuxPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "ApacheLinuxPath",
	expected: "/etc/apache/",
}

func TestGetApacheLinuxPath(t *testing.T) {
	getPathTestCommonMethodLinux(t, getApacheLinuxPathTestCaseData, knowndirget.ApacheLinuxPath)
}
