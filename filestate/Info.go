package filestate

import (
	"fmt"
	"os"
	"time"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/namevalue"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/elitepath"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/internal/pathcompareinternal"
	"gitlab.com/evatix-go/pathhelper/pathsysinfo"
)

type Info struct {
	FullPath           string
	LastModified       time.Time
	Chmod              os.FileMode
	UserGroupId        *pathsysinfo.UserGroupId
	Size               int64
	HexContentChecksum string
	IsFile, IsInvalid  bool
	HashMethod         hashas.Variant
	toString           corestr.SimpleStringOnce
}

func (it *Info) IsNull() bool {
	return it == nil
}

func (it *Info) IsNotNull() bool {
	return it != nil
}

func (it *Info) IsDir() bool {
	return it != nil && !it.IsFile
}

func (it *Info) IsDirExist() bool {
	return it != nil && !it.IsFile && !it.IsInvalid
}

func (it *Info) IsFileExist() bool {
	return it != nil && it.IsFile && !it.IsInvalid
}

func (it *Info) IsExist() bool {
	return it != nil && !it.IsInvalid
}

func (it *Info) Stat() *chmodhelper.PathExistStat {
	if it == nil {
		return nil
	}

	return chmodhelper.GetPathExistStat(it.FullPath)
}

func (it *Info) LocationInfo() *pathhelper.LocationInfo {
	if it == nil {
		return nil
	}

	return pathhelper.GetLocationInfo(it.FullPath)
}

func (it *Info) ElitePath() *elitepath.Path {
	if it == nil {
		return nil
	}

	return elitepath.NewPathDefault(it.FullPath)
}

func (it *Info) SizePtr() *int64 {
	if it == nil || it.IsInvalid {
		return nil
	}

	return &it.Size
}

func (it *Info) LastModifiedPtr() *time.Time {
	if it == nil || it.IsInvalid {
		return nil
	}

	return &it.LastModified
}

func (it *Info) IsLastModifiedEqual(right *Info) bool {
	return pathcompareinternal.IsLastModifiedEqualPtr(
		it.LastModifiedPtr(),
		right.LastModifiedPtr())
}

func (it *Info) ReadCurrentRawChecksum() *errbyte.Results {
	if it == nil {
		return errbyte.New.Results.ErrorWrapper(errnew.Null.Simple(it))
	}

	return it.HashMethod.SumOfFile(it.FullPath)
}

func (it *Info) ReadCurrentHexChecksum() *errstr.Result {
	if it == nil {
		return errstr.
			New.
			Result.
			ErrorWrapper(
				errnew.Null.Simple(it))
	}

	return it.HashMethod.HexSumOfFile(it.FullPath)
}

func (it *Info) ReadCurrentHexChecksumString() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.HashMethod.HexSumOfFile(it.FullPath).String()
}

func (it Info) Json() corejson.Result {
	return corejson.New(it)
}

func (it Info) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Info) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

func (it Info) JsonModelAny() interface{} {
	return it
}

func (it *Info) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Info) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Info) IsChecksumEqualUsingInfo(
	isIgnoreOnAnyEmpty bool,
	another *Info,
) bool {
	return IsInfoChecksumEqual(
		isIgnoreOnAnyEmpty,
		it,
		another,
	)
}

func (it *Info) IsEqualDefault(
	right *Info,
) bool {
	return IsEqual(
		false,
		false,
		false,
		false,
		it,
		right)
}

func (it *Info) IsEqual(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	right *Info,
) bool {
	return IsEqual(
		isIgnoreModifiedTimeCompare,
		isIgnoreChmodCompare,
		isIgnoreChownCompare,
		isIgnoreCompareOnAnyEmpty,
		it,
		right)
}

func (it *Info) RwxWrapper() *chmodhelper.RwxWrapper {
	return chmodhelper.
		New.
		RwxWrapper.
		UsingFileModePtr(it.Chmod)
}

func (it *Info) chmodString() string {
	rwxWrapper := it.RwxWrapper()

	return fmt.Sprintf(chmodSprintFormat,
		rwxWrapper.ToFileModeString(),
		it.Chmod.String())
}

func (it *Info) nameValues() []namevalue.Instance {
	slice := []namevalue.Instance{
		{
			Name: "FullPath",
			Value: fmt.Sprintf(
				constants.SprintDoubleQuoteFormat,
				it.FullPath),
		},
		{
			Name:  "IsInvalid",
			Value: it.IsInvalid,
		},
		{
			Name:  "Chmod",
			Value: it.chmodString(),
		},
	}

	if it.IsFile {
		slice = append(slice,
			namevalue.Instance{
				Name:  "IsFile",
				Value: it.IsFile,
			},
			namevalue.Instance{
				Name:  "HexChecksum",
				Value: it.HexContentChecksum,
			},
			namevalue.Instance{
				Name:  "LastModified",
				Value: it.LastModified,
			},
			namevalue.Instance{
				Name:  "Size",
				Value: it.Size,
			})
	}

	if it.IsDir() {
		slice = append(slice,
			namevalue.Instance{
				Name:  "IsDirectory",
				Value: true,
			},
		)
	}

	if it.UserGroupId != nil {
		slice = append(slice,
			namevalue.Instance{
				Name:  "Chown/UserGroupId",
				Value: it.UserGroupId.String(),
			},
		)
	}

	return slice
}

func (it *Info) String() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.toString.IsInitialized() {
		return it.toString.String()
	}

	finalString := errcore.VarNameValuesJoiner(
		consts.FileInfoEachLineJoiner,
		it.nameValues()...)

	return it.
		toString.
		GetPlusSetOnUninitialized(
			finalString)
}
