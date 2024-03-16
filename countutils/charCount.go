package countutils

func CountBytes(fileContent string) int {
	fileBytes := []byte(fileContent)
	return len(fileBytes)
}
