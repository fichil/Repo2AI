package packer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fichil/Repo2AI/internal/scanner"
)

const outputDir = "output"

func Generate(manifest *scanner.Manifest) error {
	if manifest == nil {
		return fmt.Errorf("manifest is nil")
	}

	groups := map[string][]scanner.FileInfo{}

	for _, file := range manifest.Files {
		category := file.Category
		if category == "" {
			category = "others"
		}
		groups[category] = append(groups[category], file)
	}

	for category, files := range groups {
		if len(files) == 0 {
			continue
		}

		err := writeCategoryPack(manifest.RootPath, category, files)
		if err != nil {
			return err
		}
	}

	fmt.Println("AI Context Pack generated:", outputDir)
	return nil
}

func writeCategoryPack(rootPath string, category string, files []scanner.FileInfo) error {
	fileName := fmt.Sprintf("%s_01.md", category)
	outputPath := filepath.Join(outputDir, fileName)

	var builder strings.Builder

	builder.WriteString("# Repo2AI Context Pack\n\n")
	builder.WriteString("Category: ")
	builder.WriteString(category)
	builder.WriteString("\n\n")

	for _, file := range files {
		fullPath := filepath.Join(rootPath, filepath.FromSlash(file.Path))

		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("read file failed: %s, error: %w", file.Path, err)
		}

		builder.WriteString("---\n\n")
		builder.WriteString("## ")
		builder.WriteString(file.Path)
		builder.WriteString("\n\n")

		builder.WriteString("```")
		builder.WriteString(languageByType(file.Type))
		builder.WriteString("\n")
		builder.Write(content)
		builder.WriteString("\n```\n\n")
	}

	err := os.WriteFile(outputPath, []byte(builder.String()), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Context pack generated:", filepath.ToSlash(outputPath))
	return nil
}

func languageByType(fileType string) string {
	switch fileType {
	case "java":
		return "java"
	case "xml":
		return "xml"
	case "json":
		return "json"
	case "yaml":
		return "yaml"
	case "properties":
		return "properties"
	case "sql":
		return "sql"
	case "go":
		return "go"
	case "markdown":
		return "markdown"
	default:
		return ""
	}
}
