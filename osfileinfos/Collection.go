package osfileinfos

import "os"

type Collection struct {
	Items *[]os.FileInfo
}

func New(infos *[]os.FileInfo) *Collection {
	return &Collection{
		Items: infos,
	}
}

func NewUsingCap(cap int) *Collection {
	infos := make([]os.FileInfo, 0, cap)

	return &Collection{
		Items: &infos,
	}
}

func (receiver *Collection) IsEmpty() bool {
	return receiver.Items == nil ||
		len(*receiver.Items) == 0
}

func (receiver *Collection) HasItems() bool {
	return receiver.Items != nil &&
		len(*receiver.Items) > 0
}

func (receiver *Collection) Length() int {
	if receiver.Items == nil {
		return 0
	}

	return len(*receiver.Items)
}

func (receiver *Collection) AddInfo(info os.FileInfo) *Collection {
	if info == nil {
		return receiver
	}

	*receiver.Items = append(
		*receiver.Items,
		info)

	return receiver
}

func (receiver *Collection) Add(info os.FileInfo, err error) error {
	if err != nil || info == nil {
		return err
	}

	*receiver.Items = append(
		*receiver.Items,
		info)

	return err
}
