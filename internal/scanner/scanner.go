package scanner

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Type string `json:"type"`
}

type Manifest struct {
	ProjectName  string     `json:"projectName"`
	RootPath     string     `json:"rootPath"`
	TotalFiles   int        `json:"totalFiles"`
	JavaFiles    int        `json:"javaFiles"`
	XmlFiles     int        `json:"xmlFiles"`
	IgnoredFiles int        `json:"ignoredFiles"`
	Files        []FileInfo `json:"files"`
}

func Scan(root string) (*Manifest, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	manifest := &Manifest{
		ProjectName: filepath.Base(absRoot),
		RootPath:    absRoot,
		Files:       make([]FileInfo, 0),
	}

	err = filepath.WalkDir(absRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(absRoot, path)
		relPath = filepath.ToSlash(relPath)

		if relPath == "." {
			return nil
		}

		if ShouldIgnore(relPath, d.IsDir()) {
			manifest.IgnoredFiles++
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		fileType := detectType(relPath)

		manifest.TotalFiles++
		if fileType == "java" {
			manifest.JavaFiles++
		}
		if fileType == "xml" {
			manifest.XmlFiles++
		}

		manifest.Files = append(manifest.Files, FileInfo{
			Path: relPath,
			Size: info.Size(),
			Type: fileType,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return manifest, nil
}

func WriteManifest(manifest *Manifest, outputPath string) error {
	err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputPath, data, 0644)
}

func detectType(path string) string {
	lower := strings.ToLower(path)

	switch {
	case strings.HasSuffix(lower, ".java"):
		return "java"

	case strings.HasSuffix(lower, ".xml"):
		return "xml"

	case strings.HasSuffix(lower, ".go"):
		return "go"

	case strings.HasSuffix(lower, ".md"):
		return "markdown"

	case strings.HasSuffix(lower, ".json"):
		return "json"

	case strings.HasSuffix(lower, ".yml"),
		strings.HasSuffix(lower, ".yaml"):
		return "yaml"

	default:
		return "other"
	}
}
