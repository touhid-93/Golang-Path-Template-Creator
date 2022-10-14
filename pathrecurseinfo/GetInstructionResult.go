package pathrecurseinfo

import (
	"io/fs"
	"path/filepath"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// GetInstructionResult returns result
func GetInstructionResult(instruction *Instruction) *Result {
	if instruction == nil || instruction.Root == "" {
		return InvalidResult(
			constants.EmptyString,
			errnew.Path.Empty(),
			nil)
	}

	expand := expandpath.ExpandVariablesIf(
		instruction.IsExpandEnvironmentVar,
		instruction.Root)

	normalizedRoot := normalize.PathUsingSingleIf(
		instruction.IsNormalize,
		expand)

	pathStat := chmodhelper.GetPathExistStat(
		normalizedRoot)

	if pathStat.HasError() {
		errWrap := errnew.
			Path.
			Error(
				errtype.MissingPathsOrInvalidPaths,
				pathStat.Error,
				normalizedRoot)

		return InvalidResult(
			normalizedRoot,
			errWrap,
			pathStat)
	}

	if !pathStat.IsExist {
		// not exist
		errWrap := errnew.
			Path.
			Messages(
				errtype.MissingPathsOrInvalidPaths,
				normalizedRoot)

		return InvalidResult(
			normalizedRoot,
			errWrap,
			pathStat)
	}

	if pathStat.IsFile() {
		if instruction.IsRelativePath {
			normalizedRoot = strings.TrimPrefix(normalizedRoot, normalizedRoot)
		}

		return &Result{
			Root:            normalizedRoot,
			PathStat:        pathStat,
			IsInvalidResult: false,
			PathsResult: &PathsResult{
				ExpandingPaths: corestr.New.SimpleSlice.SpreadStrings(
					normalizedRoot),
				IsExist: true,
				IsFile:  true,
				IsDir:   false,
			},
			IsRelative:   instruction.IsRelativePath,
			ErrorWrapper: nil,
		}
	}

	if !instruction.IsRecursive {
		return nonRecursiveResult(
			normalizedRoot,
			instruction,
			pathStat)
	}

	paths := make(
		[]string,
		0,
		constants.ArbitraryCapacity32)

	var sliceErr []string
	isExcludeAny := instruction.HasExcludingRootNames()
	nameExcludes := instruction.ExcludingNamesHashset()
	isExcludeRoot := instruction.IsExcludeRoot
	isRelativePath := instruction.IsRelativePath
	excludingPaths := instruction.ExcludingPathsHashset()
	hasAnyExcludingPaths := excludingPaths.HasAnyItem()

	finalErr := filepath.Walk(
		normalizedRoot,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				sliceErr = append(sliceErr, err.Error()+" - "+path)

				return err
			}

			if info == nil {
				sliceErr = append(sliceErr, "Nil file info - "+path)

				return err
			}

			if isExcludeRoot && normalizedRoot == path {
				return nil
			}

			isDir := info.IsDir()
			isExcludeRootCondition := isExcludeAny &&
				nameExcludes.Has(info.Name())
			isRootNameExclude := isExcludeRootCondition &&
				normalizedRoot != path &&
				strings.TrimPrefix(path, normalizedRoot)[1:] == info.Name()

			if isRootNameExclude && isDir {
				return filepath.SkipDir
			}

			if isRootNameExclude && !isDir {
				return nil
			}

			isExcludingPath := hasAnyExcludingPaths &&
				excludingPaths.Has(path)

			if isExcludingPath && isDir {
				return filepath.SkipDir
			}

			if isExcludingPath {
				return nil
			}

			finalizedPath := path
			if isRelativePath {
				finalizedPath = strings.Replace(
					finalizedPath,
					normalizedRoot,
					constants.EmptyString,
					1)
			}

			if isRelativePath && finalizedPath != "" {
				finalizedPath = finalizedPath[1:]
			}

			if finalizedPath == "" {
				return nil
			}

			switch {
			case instruction.IsIncludeAll:
				paths = append(paths, finalizedPath)
			case instruction.IsIncludeDirsOnly && info.IsDir():
				paths = append(paths, finalizedPath)
			case instruction.IsIncludeFilesOnly && !info.IsDir():
				paths = append(paths, finalizedPath)
			}

			return nil
		},
	)

	if finalErr != nil {
		sliceErr = append(sliceErr, finalErr.Error())
	}

	compiledErr := errcore.SliceToError(sliceErr)

	return &Result{
		Root:            normalizedRoot,
		PathStat:        pathStat,
		IsInvalidResult: compiledErr != nil,
		PathsResult: &PathsResult{
			ExpandingPaths: corestr.New.SimpleSlice.Strings(paths),
			IsExist:        true,
			IsFile:         false,
			IsDir:          true,
		},
		IsRelative: instruction.IsRelativePath,
		ErrorWrapper: errnew.
			Path.
			Error(
				errtype.PathExpand,
				compiledErr,
				normalizedRoot),
	}
}
