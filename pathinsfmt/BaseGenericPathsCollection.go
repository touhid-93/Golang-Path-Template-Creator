package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type BaseGenericPathsCollection struct {
	GenericPathsCollection *GenericPathsCollection `json:"GenericPathsCollection,omitempty"`
}

func (b *BaseGenericPathsCollection) GenericPathsCollectionLength() int {
	if b.GenericPathsCollection == nil {
		return constants.Zero
	}

	return b.GenericPathsCollection.Length()
}

func (b *BaseGenericPathsCollection) IsEmptyGenericPathsCollection() bool {
	return b.GenericPathsCollection.Length() == 0
}

func (b *BaseGenericPathsCollection) HasAnyGenericPathsCollection() bool {
	return b.GenericPathsCollection.Length() > 0
}
