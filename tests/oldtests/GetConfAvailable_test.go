package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getConfAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetConfAvailable",
	expected: "/etc/apache/conf-available",
}

func TestGetConfAvailable(t *testing.T) {
	getPathTestCommonMethodLinux(t, getConfAvailablePathTestCaseData, apachelinuxpath.GetConfAvailable)
}
