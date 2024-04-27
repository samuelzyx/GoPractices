func solution(n int) int {
	strNum := strconv.Itoa(n) //Int to String
	nArrStr := strings.Split(strNum, "")
	result := 0

	for i := 0; i < len(nArrStr); i++ {
		num, _ := strconv.Atoi(nArrStr[i]) //String to int
		result += num
	}

	return result
}
