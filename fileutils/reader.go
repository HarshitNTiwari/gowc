package fileutils

import (
	"os"
)

func ReadTextFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}
