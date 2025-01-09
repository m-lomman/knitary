package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getPatternContent(file string) string {
	file = filepath.Join("patterns", file)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading pattern file: %w", err)
		return ""
	}
	return string(content)
}
