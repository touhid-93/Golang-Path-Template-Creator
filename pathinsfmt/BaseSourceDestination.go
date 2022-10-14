package pathinsfmt

type BaseSourceDestination struct {
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
}

func (it *BaseSourceDestination) Clone() *BaseSourceDestination {
	if it == nil {
		return nil
	}

	return &BaseSourceDestination{
		Source:      it.Source,
		Destination: it.Destination,
	}
}
