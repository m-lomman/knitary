package main

import (
	"fmt"
)

func main() {
	content := getPatternContent("testfile.txt")
	fmt.Println(content)

	patternContent := getPatternContent("testPattern.txt")
	patternBinary := knitToBinary(patternContent)
	println("\n" + patternContent)

	println("\n\n" + patternBinary + "\n\n")

	patternContent = getPatternContent("towelPattern.txt")
	patternBinary = knitToBinary(patternContent)
	println("\n" + patternContent)

	println("\n\n" + patternBinary)
}
