package scanner

import (
	"path/filepath"
	"strings"
)

func ShouldIgnore(path string, isDir bool) bool {
	path = filepath.ToSlash(path)
	lower := strings.ToLower(path)

	ignoreDirs := []string{
		".git",
		".idea",
		"target",
		"node_modules",
		"logs",
		"out",
		"build",
		"output",
	}

	for _, dir := range ignoreDirs {
		if lower == dir || strings.HasPrefix(lower, dir+"/") || strings.Contains(lower, "/"+dir+"/") {
			return true
		}
	}

	ignoreExts := []string{
		".class",
		".jar",
		".war",
		".ear",
		".exe",
		".dll",
		".log",
		".zip",
		".tar",
		".gz",
	}

	for _, ext := range ignoreExts {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}

	return false
}
