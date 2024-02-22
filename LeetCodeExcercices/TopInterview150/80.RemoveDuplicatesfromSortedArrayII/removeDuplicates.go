func removeDuplicates(nums []int) int {
	count := 0
	currentNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentNum == nums[i] {
			count++
			if count > 1 {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		} else {
			count = 0
			currentNum = nums[i]
		}
	}
	return len(nums)
}