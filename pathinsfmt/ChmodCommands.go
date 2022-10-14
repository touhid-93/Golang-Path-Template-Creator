package pathinsfmt

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

type ChmodCommands struct {
	Condition *chmodins.Condition `json:"Condition,omitempty"`
	Commands  []string            `json:"Commands,omitempty"`
}

func (c *ChmodCommands) Length() int {
	if c == nil {
		return constants.Zero
	}

	return len(c.Commands)
}

func (c *ChmodCommands) HasAnyItem() bool {
	return c.Length() > 0
}

func (c *ChmodCommands) IsEmpty() bool {
	return c.Length() == 0
}

func (c *ChmodCommands) HasConditions() bool {
	if c == nil {
		return false
	}

	return c.Condition != nil
}

func (c *ChmodCommands) CreateCmdOnceCollection() *errcmd.CmdOnceCollection {
	cmdOnceCollection := errcmd.NewCmdOnceCollection(c.Length() + constants.One)

	if c.IsEmpty() {
		return cmdOnceCollection
	}

	return cmdOnceCollection.
		AddBashEachScriptAsEachCmdOnce(c.Commands...)
}

func (c *ChmodCommands) Clone() *ChmodCommands {
	if c == nil {
		return nil
	}

	return &ChmodCommands{
		Condition: c.Condition.Clone(),
		Commands:  stringslice.Clone(c.Commands),
	}
}
