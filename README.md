# gowc

## Unix-like wc command-line tool written in Go

This is a Go implementation to mimic the Unix's wc utility that dislays the number of lines, words and bytes contained in an input file - with a few changes.

- In context of this tool, counting lines means counting the number of new line characters in the input file.
- A word is defined as a string of characters delimited by characters which are not alphabets or digits.

### Usage

1. Clone the repo and put the input text files in the root folder.

2. Following is the Synopsis for the tool:

```cmd
gowc [-options] [file ...]
```

3. Following are the options that can be used as flags:

```cmd
-c     Count the number of bytes in the input file
-l     Count the number of new line characters in the input file
-w     Count the number of words in the input file
```

The default values of these flags is `false`.

4. Here is an example to get the number of words in the input text file:

```cmd
go run ./gowc.go -w file.txt
```

5. When using multiple flags together, the order of output will be: lines, words, bytes and file name.

6. If more than one input file is specified, a line for each file is printed showing the results.
