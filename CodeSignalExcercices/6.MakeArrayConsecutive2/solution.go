import "sort"

func solution(statues []int) int {
	additionalStatues := 0
	smallest := 0
	n := len(statues)

	sort.Slice(statues, func(i, j int) bool { return statues[i] < statues[j] })

	if n <= 0 {
		return 0
	}

	smallest = statues[0]

	for i := 1; i < n; i++ {
		if statues[i] == (smallest + 1) {
			smallest++
		} else {
			additionalStatues += statues[i] - (smallest + 1)
			smallest += statues[i] - (smallest)
		}
	}

	return additionalStatues
}
