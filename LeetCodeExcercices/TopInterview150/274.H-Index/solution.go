func hIndex(citations []int) int {
	hIndex := 0
	sort.Slice(citations, func(i, j int) bool { return citations[i] > citations[j] })
	for i, element := range citations {
		if i+1 <= element {
			hIndex++
		} else {
			return hIndex
		}
	}
	return hIndex
}