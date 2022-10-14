package knowndir

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/core/ostype"
	"gitlab.com/evatix-go/enum/osarchs"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/knowndir/os"
)

type AliasStruct struct {
	Name      string
	Alias     Alias
	Linux     *os.Linux
	Windows   *os.Windows
	Mac       *os.Mac
	specifics os.DirSpecifier
}

func (a *AliasStruct) Specifics() os.DirSpecifier {
	if a.specifics != nil {
		return a.specifics
	}

	if osconsts.IsWindows {
		a.specifics = a.Windows
	} else if osconsts.IsLinux {
		a.specifics = a.Linux
	} else if osconsts.IsDarwinOrMacOs {
		a.specifics = a.Mac
	}

	return a.specifics
}

func (a *AliasStruct) HasLinux() bool {
	return a.Linux != nil
}

func (a *AliasStruct) HasWindows() bool {
	return a.Windows != nil
}

func (a *AliasStruct) HasMac() bool {
	return a.Mac != nil
}

func (a *AliasStruct) IsLinuxEmpty() bool {
	return a.Linux == nil
}

func (a *AliasStruct) IsWindowsEmpty() bool {
	return a.Windows == nil
}

func (a *AliasStruct) IsMacEmpty() bool {
	return a.Mac == nil
}

func (a *AliasStruct) IsOsSpecificsEmpty() bool {
	return a.Specifics() == nil
}

func (a *AliasStruct) HasOsSpecifics() bool {
	return a.Specifics() != nil
}

func (a *AliasStruct) GetDirSpecifier(
	osType ostype.Variation,
) (
	os.DirSpecifier,
	*errorwrapper.Wrapper,
) {
	switch osType {
	case ostype.Windows:
		return a.Windows, nil
	case ostype.Linux:
		return a.Linux, nil
	case ostype.DarwinOrMacOs:
		return a.Mac, nil
	default:
		msg := "GetDirSpecifier failed because non supported os given"

		return nil,
			errnew.NotSupportedOption(
				"os-type",
				osType.String(),
				msg)
	}
}

func (a *AliasStruct) GetCurrentDir() (
	string, *errorwrapper.Wrapper,
) {
	specifier, err := a.GetDirSpecifier(ostype.GetCurrentVariant())

	if err.HasError() || specifier == nil {
		return constants.EmptyString, err
	}

	return specifier.Generated(), err
}

func (a *AliasStruct) GetDir(
	architecture osarchs.Architecture,
	osType ostype.Variation,
) (
	string, *errorwrapper.Wrapper,
) {
	specifier, err := a.GetDirSpecifier(osType)

	if err.HasError() || specifier == nil {
		return constants.EmptyString, err
	}

	return specifier.GetDir(architecture), err
}

func (a *AliasStruct) GetDirMust(
	architecture osarchs.Architecture,
	osType ostype.Variation,
) string {
	specifier, err := a.GetDir(architecture, osType)
	err.HandleError()

	return specifier
}
