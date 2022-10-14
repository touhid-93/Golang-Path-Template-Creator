package hashas

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type Variant byte

const (
	Invalid Variant = iota
	Md5
	Sha1
	Sha256
	Sha512
)

func (it Variant) NewHash() (hash.Hash, *errorwrapper.Wrapper) {
	switch it {
	case Invalid:
		return nil, errnew.Messages.Many(
			errtype.UnexpectedDefinition,
			it.Name()+"(HashMethod/Variant) is expected to be not defined. Thus nil hasher.",
		)
	case Md5:
		return md5.New(), nil
	case Sha1:
		return sha1.New(), nil
	case Sha256:
		return sha256.New(), nil
	case Sha512:
		return sha512.New(), nil
	default:
		return nil, errnew.Messages.Many(
			errtype.InvalidOption,
			it.Name()+" invalid option.",
			BasicEnumImpl.RangesInvalidMessage(),
		)
	}
}

func (it Variant) NewHashError() (hash.Hash, error) {
	switch it {
	case Invalid:
		return nil, errtype.UnexpectedDefinition.ReferencesCsvError(
			"(HashMethod/Variant) is expected to be not defined. Thus nil hasher.",
			it.Name(),
		)
	case Md5:
		return md5.New(), nil
	case Sha1:
		return sha1.New(), nil
	case Sha256:
		return sha256.New(), nil
	case Sha512:
		return sha512.New(), nil
	default:
		return nil, errtype.InvalidOption.ReferencesCsvError(
			BasicEnumImpl.RangesInvalidMessage(),
			it.Name(),
		)
	}
}

func (it Variant) HexSumOf(
	inputBytes []byte,
) *errstr.Result {
	return HexChecksumOfRawBytes(it, inputBytes)
}

func (it Variant) HexSumOfFile(
	fileName string,
) *errstr.Result {
	return HexChecksumOfFilePath(it, fileName)
}

func (it Variant) HexSumOfFileNoError(
	fullPath string,
) string {
	return HexChecksumOfFilePathNoError(
		false,
		it,
		fullPath)
}

func (it Variant) HexSumOfFileNoErrorIf(
	isSkipGenerate bool,
	fullPath string,
) string {
	if isSkipGenerate {
		return constants.EmptyString
	}

	return HexChecksumOfFilePathNoError(
		false,
		it,
		fullPath)
}

func (it Variant) HexSumOfFileIf(
	isSkipGenerate bool,
	fullPath string,
) *errstr.Result {
	if isSkipGenerate {
		return errstr.Empty.Result()
	}

	return HexChecksumOfFilePath(
		it,
		fullPath)
}

func (it Variant) SumOfFile(
	filePath string,
) *errbyte.Results {
	return SumOfFile(it, filePath)
}

func (it Variant) SumOf(
	inputBytes []byte,
) *errbyte.Results {
	return BytesChecksum(it, inputBytes)
}

func (it Variant) SumOfErrorBytes(
	errBytes *errbyte.Results,
) *errbyte.Results {
	return ErrorWrapperWithBytesChecksum(it, errBytes)
}

func (it *Variant) HexSumOfAny(
	item interface{},
) *errstr.Result {
	jsonResult := corejson.NewPtr(item)

	return it.HexOfJsonResult(jsonResult)
}

func (it *Variant) HexSumOfAnyIf(
	isGenerate bool,
	item interface{},
) *errstr.Result {
	if isGenerate {
		jsonResult := corejson.NewPtr(item)

		return it.HexOfJsonResult(jsonResult)
	}

	return errstr.Empty.Result()
}

func (it Variant) HexSumOfAnyItemsToCombinedSingleString(
	isSkipOnNil bool,
	items ...interface{},
) *errstr.Result {
	return HexChecksumOfAnyItemsToCombinedSingleString(
		isSkipOnNil,
		it,
		items...)
}

func (it Variant) HexSumOfAnyItems(
	isSkipOnNil bool,
	items ...interface{},
) *errstr.Results {
	return HexChecksumOfAnyItems(
		isSkipOnNil,
		it,
		items...)
}

func (it *Variant) SumOfJsonResult(
	result *corejson.Result,
) *errbyte.Results {
	if result == nil || result.Bytes == nil {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Null.Message(
				"cannot hash nil json result or nil bytes values!",
			))
	}

	if result.HasError() {
		return errbyte.New.Results.ErrorWrapper(
			errnew.Messages.Many(
				errtype.JsonSyntaxIssue,
				"cannot hash on error json results!",
				result.MeaningfulError().Error()))
	}

	return it.SumOf(result.Bytes)
}

func (it *Variant) HexOfJsonResult(
	result *corejson.Result,
) *errstr.Result {
	bytesResult := it.SumOfJsonResult(result)

	if bytesResult.HasError() {
		return errstr.New.Result.ErrorWrapper(bytesResult.ErrorWrapper)
	}

	toString := bytesResult.NonEmptyString(
		convertBytesResultsToEncodedHexString)

	return errstr.New.Result.ValueOnly(toString)
}

func (it Variant) IsUndefined() bool {
	return it == Invalid
}

func (it Variant) IsMd5() bool {
	return it == Md5
}

func (it Variant) IsSha1() bool {
	return it == Sha1
}

func (it Variant) IsSha256() bool {
	return it == Sha256
}

func (it Variant) IsSha512() bool {
	return it == Sha512
}

func (it Variant) ValueUInt16() uint16 {
	return uint16(it)
}

func (it Variant) AllNameValues() []string {
	return BasicEnumImpl.AllNameValues()
}

func (it Variant) OnlySupportedErr(names ...string) error {
	return BasicEnumImpl.OnlySupportedErr(names...)
}

func (it Variant) OnlySupportedMsgErr(message string, names ...string) error {
	return BasicEnumImpl.OnlySupportedMsgErr(message, names...)
}

func (it Variant) IntegerEnumRanges() []int {
	return BasicEnumImpl.IntegerEnumRanges()
}

func (it Variant) MinMaxAny() (min, max interface{}) {
	return BasicEnumImpl.MinMaxAny()
}

func (it Variant) MinValueString() string {
	return BasicEnumImpl.MinValueString()
}

func (it Variant) MaxValueString() string {
	return BasicEnumImpl.MaxValueString()
}

func (it Variant) MaxInt() int {
	return BasicEnumImpl.MaxInt()
}

func (it Variant) MinInt() int {
	return BasicEnumImpl.MinInt()
}

func (it Variant) RangesDynamicMap() map[string]interface{} {
	return BasicEnumImpl.RangesDynamicMap()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) IsAnyNamesOf(names ...string) bool {
	return BasicEnumImpl.IsAnyNamesOf(it.ValueByte(), names...)
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it Variant) IsAnyValuesEqual(anyByteValues ...byte) bool {
	return BasicEnumImpl.IsAnyOf(it.ValueByte(), anyByteValues...)
}

func (it Variant) IsByteValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsValueEqual(value byte) bool {
	return it.ValueByte() == value
}

func (it Variant) ValueInt8() int8 {
	return int8(it)
}

func (it Variant) ValueInt16() int16 {
	return int16(it)
}

func (it Variant) ValueInt32() int32 {
	return int32(it)
}

func (it Variant) ValueString() string {
	return it.ToNumberString()
}

func (it Variant) Format(format string) (compiled string) {
	return BasicEnumImpl.Format(format, it.ValueByte())
}

func (it Variant) EnumType() enuminf.EnumTyper {
	return BasicEnumImpl.EnumType()
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.ValueByte())
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.ValueByte())
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	dataConv, err := it.UnmarshallEnumToValue(
		data)

	if err == nil {
		*it = Variant(dataConv)
	}

	return err
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) IsEqual(level Variant) bool {
	return level == it
}

func (it Variant) IsAboveOrEqual(level Variant) bool {
	return level.ValueByte() >= it.ValueByte()
}

func (it Variant) IsLowerOrEqual(level Variant) bool {
	return level.ValueByte() <= it.ValueByte()
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsAnyOf(anyOfItems ...Variant) bool {
	for _, item := range anyOfItems {
		if item == it {
			return true
		}
	}

	return false
}

func (it Variant) UnmarshallEnumToValue(
	jsonUnmarshallingValue []byte,
) (byte, error) {
	return BasicEnumImpl.UnmarshallToValue(
		true,
		jsonUnmarshallingValue)
}

func (it Variant) MaxByte() byte {
	return BasicEnumImpl.Max()
}

func (it Variant) MinByte() byte {
	return BasicEnumImpl.Min()
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) RangesByte() []byte {
	return BasicEnumImpl.Ranges()
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.ValueByte())
}

func (it *Variant) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it Variant) Json() corejson.Result {
	return corejson.New(it)
}

func (it Variant) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Variant) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Variant) AsJsoner() corejson.Jsoner {
	return it
}

func (it Variant) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}

func (it Variant) AsBasicByteEnumContractsBinder() enuminf.BasicByteEnumContractsBinder {
	return &it
}

func (it Variant) AsBasicEnumContractsBinder() enuminf.BasicEnumContractsBinder {
	return &it
}

func (it Variant) ToPtr() *Variant {
	return &it
}
