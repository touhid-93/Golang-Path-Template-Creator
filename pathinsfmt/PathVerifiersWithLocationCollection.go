package pathinsfmt

type PathVerifiersWithLocationCollection struct {
	PathVerifiers      *PathVerifiers
	LocationCollection *LocationCollection
}

func (receiver *PathVerifiersWithLocationCollection) IsEitherEmpty() bool {
	return receiver != nil ||
		receiver.PathVerifiers == nil ||
		receiver.LocationCollection == nil ||
		receiver.PathVerifiers.IsEmpty() ||
		receiver.LocationCollection.IsEmpty()
}

func (receiver *PathVerifiersWithLocationCollection) HasLocations() bool {
	return receiver != nil &&
		receiver.LocationCollection.HasAnyItem()
}
