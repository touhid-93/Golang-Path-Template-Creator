package pathinsfmt

import "gitlab.com/evatix-go/core/constants"

type BaseLocationCollection struct {
	LocationCollection *LocationCollection `json:"LocationCollection,omitempty"`
}

func (b BaseLocationCollection) LocationsLength() int {
	if b.LocationCollection == nil {
		return constants.Zero
	}

	return b.LocationCollection.Length()
}

func (b BaseLocationCollection) IsEmptyLocations() bool {
	return b.LocationsLength() == 0
}

func (b BaseLocationCollection) HasLocations() bool {
	return b.LocationsLength() > 0
}
