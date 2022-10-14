package downloadinsexec

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func downloadChecksumVerify(download *pathinsfmt.Download) *errorwrapper.Wrapper {
	if download.ChecksumVerify == constants.EmptyString {
		return nil
	}

	downloadPath := download.FullPath()
	if downloadPath == constants.EmptyString {
		return errnew.Messages.Many(
			errtype.CheckSum,
			"checksum cannot be verified, file :"+
				downloadPath+
				" do not exist",
		)
	}

	hashedStrResult := download.ChecksumVerifyMethod.HexSumOfFile(downloadPath)
	if hashedStrResult.ErrorWrapper.HasError() {
		return hashedStrResult.ErrorWrapper
	}

	if hashedStrResult.Value != download.ChecksumVerify {
		return errnew.Messages.Many(
			errtype.CheckSum,
			"checksum mismatch -> file : "+
				downloadPath+
				" should have hash value : "+
				download.ChecksumVerify,
		)
	}

	return nil
}
