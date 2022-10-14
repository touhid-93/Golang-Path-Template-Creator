package pathinsfmt

// Downloads Use aria2c
// Reference : https://aria2.github.io/manual/en/html/aria2c.html#options
type Downloads struct {
	IsAsyncAll        bool       // Allows parallel downloads
	IsContinueOnError bool       // Cannot exit on IsAsyncAll true
	Downloads         []Download `json:"Downloads,omitempty"`
}

func (it *Downloads) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Downloads)
}

func (it *Downloads) IsEmpty() bool {
	return it.Length() == 0
}
