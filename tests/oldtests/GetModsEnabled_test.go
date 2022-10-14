package oldtests

import (
	"testing"

	"gitlab.com/evatix-go/pathhelper/apachelinuxpath"
)

var getModsEnabledPathTestCaseData = pathTestCaseDataWrapper{
	OSName:   "Unix OS",
	funcName: "GetModsEnabled",
	expected: "/etc/apache/mods-enabled",
}

func TestGetModsEnabled(t *testing.T) {
	getPathTestCommonMethodLinux(t, getModsEnabledPathTestCaseData, apachelinuxpath.GetModsEnabled)
}
