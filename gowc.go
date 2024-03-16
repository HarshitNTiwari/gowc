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

	flag.BoolVar(&bytesVar, "c", false, "count no. of bytes in the file")
	flag.BoolVar(&linesVar, "l", false, "count no. of lines in the file")
	flag.BoolVar(&wordsVar, "w", false, "count no. of words in the file")
	flag.BoolVar(&charVar, "m", false, "count no. of chars in the file")
	flag.Parse()

	filenames := flag.Args()

	fmt.Println(filenames)

	for fileIdx := 0; fileIdx < len(filenames); fileIdx++ {
		fileContent, _ := fileutils.ReadTextFile(string(filenames[fileIdx]))

		if bytesVar {
			charCount := countutils.CountBytes(fileContent)
			fmt.Print(charCount, " ")
		}
		if linesVar {
			lineCount := countutils.CountLines(fileContent)
			fmt.Print(lineCount, " ")
		}

		if wordsVar {
			wordCount := countutils.CountWords(fileContent)
			fmt.Print(wordCount, " ")
		}

		fmt.Println(filenames[fileIdx])
	}

}
