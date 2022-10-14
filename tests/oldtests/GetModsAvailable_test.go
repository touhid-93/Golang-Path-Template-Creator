package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getModsAvailablePathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModsAvailable",
	expected: "/etc/apache/mods-available",
}

func TestGetModsAvailable(t *testing.T) {
	getPathTestCommonMethodLinux(t, getModsAvailablePathTestCaseData, apachelinuxpath.GetModsAvailable)
}
