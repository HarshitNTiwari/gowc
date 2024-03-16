package countutils

func CountChars(fileContent string) int {
	fileBytes := []byte(fileContent)
	return len(fileBytes)
}
