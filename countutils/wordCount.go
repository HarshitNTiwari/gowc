package countutils

func checkSeparator(i int) bool {
	if (i >= 9 && i <= 10) || (i >= 32 && i <= 47) || (i >= 58 && i <= 64) || (i >= 91 && i <= 96) || (i >= 123 && i <= 126) {
		return true
	}
	return false
}

func CountWords(fileContent string) int {
	fileBytes := []byte(fileContent)
	var words int = 0
	var currWord bool = false
	for charPtr := 0; charPtr < len(fileBytes); charPtr++ {
		if !checkSeparator(int(fileBytes[charPtr])) {
			if !currWord {
				currWord = true
				words++
			}
		} else {
			currWord = false
		}
	}
	return words
}
