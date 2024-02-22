func majorityElement(nums []int) int {
	result := 0
	majority := 0
	for _, element := range nums {
		if majority == 0 {
			result = element
		}
		if element == result {
			majority++
		} else {
			majority--
		}
	}
	return result
}