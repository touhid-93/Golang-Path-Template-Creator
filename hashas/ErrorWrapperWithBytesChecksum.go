package hashas

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func ErrorWrapperWithBytesChecksum(method Variant, errBytes *errbyte.Results) *errbyte.Results {
	if errBytes == nil || errBytes.Values == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Null.Simple(errBytes))
	}

	if errBytes.HasError() {
		return errBytes
	}

	return BytesChecksum(method, errBytes.Values)
}
