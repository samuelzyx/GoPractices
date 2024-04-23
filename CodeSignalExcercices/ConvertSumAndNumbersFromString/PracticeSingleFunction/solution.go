func solution(n int) int {
	strNum := strconv.Itoa(n)
	nArrStr := strings.Split(strNum, "")
	result := 0

	for i := 0; i < len(nArrStr); i++ {
		num, _ := strconv.Atoi(nArrStr[i])
		result += num
	}

	return result
}
