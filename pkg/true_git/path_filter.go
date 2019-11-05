package true_git

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
)

type PathFilter struct {
	// BasePath can be path to the directory or a single file
	BasePath                   string
	IncludePaths, ExcludePaths []string
}

func (f *PathFilter) IsFilePathValid(filePath string) bool {
	return IsFilePathValid(filePath, f.BasePath, f.IncludePaths, f.ExcludePaths)
}

func (f *PathFilter) TrimFileBasePath(filePath string) string {
	return TrimFileBasePath(filePath, f.BasePath)
}

func (f *PathFilter) String() string {
	return fmt.Sprintf("BasePath=`%s`, IncludePaths=%v, ExcludePaths=%v", f.BasePath, f.IncludePaths, f.ExcludePaths)
}

func IsFilePathValid(filePath, basePath string, includePaths, excludePaths []string) bool {
	if !IsFileInBasePath(filePath, basePath) {
		return false
	}

	filePathInBasePath := TrimFileBasePath(filePath, basePath)

	if len(includePaths) > 0 {
		if !IsFilePathMatchesOneOfPatterns(filePathInBasePath, includePaths) {
			return false
		}
	}

	if len(excludePaths) > 0 {
		if IsFilePathMatchesOneOfPatterns(filePathInBasePath, excludePaths) {
			return false
		}
	}

	return true
}

/*
 * basePath can be path to a directory or a file.
 */
func IsFileInBasePath(filePath, basePath string) bool {
	filePath = NormalizeAbsolutePath(filePath)
	basePath = NormalizeAbsolutePath(basePath)

	if filePath == basePath {
		return true
	}

	return strings.HasPrefix(filePath, NormalizeDirectoryPrefix(basePath))
}

func IsFilePathMatchesPattern(filePath, pattern string) bool {
	matched, err := doublestar.PathMatch(pattern, filePath)
	if err != nil {
		panic(err)
	}
	if matched {
		return true
	}

	matched, err = doublestar.PathMatch(filepath.Join(pattern, "**", "*"), filePath)
	if err != nil {
		panic(err)
	}
	if matched {
		return true
	}

	return false
}

func IsFilePathMatchesOneOfPatterns(filePath string, patterns []string) bool {
	for _, pattern := range patterns {
		if IsFileInBasePath(filePath, pattern) {
			return true
		}

		if IsFilePathMatchesPattern(filePath, pattern) {
			return true
		}
	}

	return false
}

func TrimFileBasePath(filePath, basePath string) string {
	filePath = NormalizeAbsolutePath(filePath)
	basePath = NormalizeAbsolutePath(basePath)

	if filePath == basePath {
		// filePath path is always a path to a file, not a directory.
		// Thus if BasePath is equal filePath, then BasePath is a path to the file.
		// Return file name in this case by convention.
		return filepath.Base(filePath)
	}

	return strings.TrimPrefix(filePath, NormalizeDirectoryPrefix(basePath))
}

func NormalizeAbsolutePath(p string) string {
	return path.Clean(path.Join("/", p))
}

func NormalizeDirectoryPrefix(directoryPrefix string) string {
	if directoryPrefix == "/" {
		return directoryPrefix
	}
	return directoryPrefix + "/"
}
