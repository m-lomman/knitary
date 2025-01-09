package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const knit = 'k'
const purl = 'p'

func getPatternContent(file string) string {
	// get the file from patterns folder & read content
	file = filepath.Join("patterns", file)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading pattern file: %w", err)
		return ""
	}

	// convert byte array to string -> lowercase
	contentStr := string(content)
	contentStr = strings.ToLower(contentStr)

	return contentStr
}

func knitToBinary(content string) string {
	print("executing knit to binary")
	knitBinary := ""
	for _, char := range content {
		if char == knit {
			knitBinary = knitBinary + "0 "
		} else if char == purl {
			knitBinary = knitBinary + "1 "
		} else if char == '\n' {
			knitBinary = knitBinary + "\n"
		} else {
			continue
		}
	}
	return knitBinary
}
