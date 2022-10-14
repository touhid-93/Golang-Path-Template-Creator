package pathcompiler

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/enum/envtype"
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type Basic struct {
	AppName       string
	AppNameLower  string
	ProductionMap map[osmixtype.Variant]*Specific
	TestMap       map[osmixtype.Variant]*Specific
}

func (it *Basic) CurrentOs() *Specific {
	return it.By(CurrentOsType)
}

func (it *Basic) TestCurrentOs() *Specific {
	return it.TestBy(CurrentOsType)
}

func (it *Basic) ByEnv(
	envType enuminf.EnvironmentTyper,
) *Specific {
	if envType.IsAnyTestEnv() {
		return it.TestCurrentOs()
	}

	return it.CurrentOs()
}

func (it *Basic) By(
	osType osmixtype.Variant,
) *Specific {
	return it.selectByOsType(
		it.ProductionMap,
		osType)
}

func (it *Basic) ByEnvOs(
	envVariant envtype.Variant,
	osType osmixtype.Variant,
) *Specific {
	return it.ByEnvFlagOs(
		envVariant.IsAnyTestEnv(),
		osType)
}

func (it *Basic) ByEnvFlagOs(
	isTestEnv bool,
	osType osmixtype.Variant,
) *Specific {
	if isTestEnv {
		return it.selectByOsType(
			it.TestMap,
			osType)
	}

	return it.selectByOsType(
		it.ProductionMap,
		osType)
}

func (it *Basic) TestBy(
	osType osmixtype.Variant,
) *Specific {
	return it.selectByOsType(
		it.TestMap,
		osType)
}

func (it *Basic) selectByOsType(
	projectionMap map[osmixtype.Variant]*Specific,
	osType osmixtype.Variant,
) *Specific {
	if len(projectionMap) == 0 {
		return nil
	}

	switch osType {
	case osmixtype.Ubuntu,
		osmixtype.Debian,
		osmixtype.ArchLinux,
		osmixtype.RedHatEnterpriseLinux,
		osmixtype.Centos:
		specific, has := projectionMap[osType]

		if has {
			return specific
		}

		return it.selectByOsType(
			projectionMap,
			osmixtype.Linux)
	case osmixtype.Linux:
		specific, has := projectionMap[osType]

		if has {
			return specific
		}

		return it.selectByOsType(
			projectionMap,
			osmixtype.Unix)
	case osmixtype.MacOs:
		specific, has := projectionMap[osType]

		if has {
			return specific
		}

		return it.selectByOsType(
			projectionMap,
			osmixtype.Unix)
	case osmixtype.Unix, osmixtype.Windows:
		specific, has := projectionMap[osType]

		if has {
			return specific
		}

		return it.selectByOsType(
			projectionMap,
			osmixtype.AnyOs)
	case osmixtype.AnyOs:
		specific, has := projectionMap[osType]

		if has {
			return specific
		}
	}

	panic(errnew.OutOfRange(
		osmixtype.Invalid,
		osmixtype.Max(),
		osmixtype.RangesInvalidErr(),
		"given current value "+osType.NameValue()+" is not supported."))
}

func (it *Basic) JoinWithVarAppRoot(
	relativePaths ...string,
) string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.JoinWithVarAppRoot(
		relativePaths...)
}

func (it *Basic) JoinWithTempRoot(
	relativePaths ...string,
) string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.JoinWithTempRoot(
		relativePaths...)
}

func (it *Basic) JoinWithDefaultInstructionsRoot(
	relativePaths ...string,
) string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.JoinWithDefaultInstructionsRoot(
		relativePaths...)
}

func (it *Basic) JoinWithInstructionTempRoot(
	relativePaths ...string,
) string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.JoinWithInstructionTempRoot(
		relativePaths...)
}

func (it *Basic) JoinWithAppConfigRoot(
	relativePaths ...string,
) string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.JoinWithAppConfigRoot(
		relativePaths...)
}

// DefaultConfigFilePath
//
//  eg. unix :
//  pathsconst.UnixVarAppRoot + "/config/default-config.json" =>
//  "/var/opt/{app-name}/config/default-config.json"
func (it *Basic) DefaultConfigFilePath() string {
	specific := it.By(CurrentOsType)

	if specific == nil {
		return ""
	}

	return specific.DefaultConfigFilePath
}

func (it Basic) ToPtr() *Basic {
	return &it
}

func (it *Basic) ToNonPtr() Basic {
	if it == nil {
		return Basic{}
	}

	return *it
}

func (it *Basic) ReflectSetTo(toPtr interface{}) error {
	return coredynamic.ReflectSetFromTo(it, toPtr)
}

func (it *Basic) ReflectSetToErrWrap(toPtr interface{}) *errorwrapper.Wrapper {
	return errnew.Reflect.SetFromTo(it, toPtr)
}

func (it *Basic) PrettyJsonString() string {
	return corejson.NewPtr(it).PrettyJsonString()
}

func (it *Basic) Json() corejson.Result {
	return corejson.New(it)
}

func (it *Basic) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Basic) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it Basic) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
