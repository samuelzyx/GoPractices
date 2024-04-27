func solution(inputString string) bool {
	splitStr := strings.Split(inputString, "")

	for i, j := 0, len(splitStr)-1; i < j; i, j = i+1, j-1 { //Reverse array
		splitStr[i], splitStr[j] = splitStr[j], splitStr[i]
	}

	reverseStr := strings.Join(splitStr, "")

	return reverseStr == inputString
}
