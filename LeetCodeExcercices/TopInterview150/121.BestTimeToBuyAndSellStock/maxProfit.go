func maxProfit(prices []int) int {
	lenPrice := len(prices)

	currentMaxProf := 0
	maxProf := 0

	for i := 1; i < lenPrice; i++ {
		currentMaxProf += prices[i] - prices[i-1]

		if currentMaxProf < 0 {
			currentMaxProf = 0
		}
		if maxProf < currentMaxProf {
			maxProf = currentMaxProf
		}
	}

	return maxProf
}