package hashas

import "encoding/hex"

func convertBytesResultsToEncodedHexString(nonEmptyBytes []byte) string {
	return hex.EncodeToString(nonEmptyBytes)
}
