package downloadinsexec

import (
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func aria2cBashCommandArg(download *pathinsfmt.Download) string {
	cmdArray := make([]string, constants.Two, constants.Eight)
	cmdArray = append(cmdArray, Aria2C, download.Url)

	if download.Destination != constants.EmptyString {
		cmdArray = append(cmdArray, HyphenD, download.Destination)
	}

	if download.FileName != constants.EmptyString {
		cmdArray = append(cmdArray, HyphenO, download.FileName)
	}

	if download.ParallelRequests > constants.Zero {
		parallelRequest := HyphenX + strconv.Itoa(int(download.ParallelRequests))
		cmdArray = append(cmdArray, parallelRequest)
	}

	if download.MaxRetries > constants.Zero {
		maxRetries := MaxRetries + strconv.Itoa(int(download.MaxRetries))
		cmdArray = append(cmdArray, maxRetries)
	}

	return strings.Join(cmdArray, constants.Space)
}
