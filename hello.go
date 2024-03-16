package main

import (
	"flag"
	"fmt"
	"gowc/countutils"
	"gowc/fileutils"
)

func main() {
	var (
		bytesVar bool
		linesVar bool
		wordsVar bool
		charVar  bool
	)

	flag.BoolVar(&bytesVar, "c", true, "count no. of bytes in the file")
	flag.BoolVar(&linesVar, "l", true, "count no. of lines in the file")
	flag.BoolVar(&wordsVar, "w", true, "count no. of words in the file")
	flag.BoolVar(&charVar, "m", true, "count no. of chars in the file")
	flag.Parse()

	// fmt.Println(bytesVar)
	// fmt.Println(linesVar)

	fileContent, _ := fileutils.ReadTextFile("data/test.txt")
	charCount := countutils.CountChars(fileContent)
	fmt.Println(charCount)
	wordCount := countutils.CountWords(fileContent)
	fmt.Println(wordCount)

}
