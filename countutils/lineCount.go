package countutils

func CountLines(fileContent string) int {
	fileBytes := []byte(fileContent)

	var lines int = 0
	for charPtr := 0; charPtr < len(fileBytes); charPtr++ {
		if int(fileBytes[charPtr]) == 10 {
			lines++
		}
	}
	return lines
}
