package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const knit = 'k'
const purl = 'p'

var (
	boldMagenta = "\x1b[1m\x1b[35m"
	boldYellow  = "\x1b[1m\x1b[33m"
	reset       = "\x1b[0m"
)

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
			knitBinary = knitBinary + boldMagenta + "0 " + reset
		} else if char == purl {
			knitBinary = knitBinary + boldYellow + "1 " + reset
		} else if char == '\n' {
			knitBinary = knitBinary + "\n"
		} else {
			continue
		}
	}
	return knitBinary
}
