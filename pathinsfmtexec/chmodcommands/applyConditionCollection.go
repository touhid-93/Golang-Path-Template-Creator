package chmodcommands

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func applyConditionCollection(
	condition *chmodins.Condition,
	collection *errcmd.CmdOnceCollection,
) *errorwrapper.Wrapper {
	if condition.IsContinueOnError {
		return collection.ExecuteUntilErr()
	}

	return collection.ExecuteAll()
}
