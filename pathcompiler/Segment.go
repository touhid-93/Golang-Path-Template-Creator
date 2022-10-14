package pathcompiler

import "gitlab.com/evatix-go/core/coredata/corestr"

type Segment struct {
	Name, Format   string
	compiledFormat corestr.SimpleStringOnce
}
