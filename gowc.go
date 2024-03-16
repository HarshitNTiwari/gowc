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

	flag.BoolVar(&bytesVar, "c", false, "Count the number of bytes in the input file")
	flag.BoolVar(&linesVar, "l", false, "Count the number of new line characters in the input file")
	flag.BoolVar(&wordsVar, "w", false, "Count the number of words in the input file")
	flag.BoolVar(&charVar, "m", false, "Count the number of characters in the input file")
	flag.Parse()

	filenames := flag.Args()

	fmt.Println(filenames)

	for fileIdx := 0; fileIdx < len(filenames); fileIdx++ {
		fileContent, _ := fileutils.ReadTextFile(string(filenames[fileIdx]))

		if linesVar {
			lineCount := countutils.CountLines(fileContent)
			fmt.Print(lineCount, " ")
		}

		if wordsVar {
			wordCount := countutils.CountWords(fileContent)
			fmt.Print(wordCount, " ")
		}

		if bytesVar {
			charCount := countutils.CountBytes(fileContent)
			fmt.Print(charCount, " ")
		}

		fmt.Println(filenames[fileIdx])
	}

}
