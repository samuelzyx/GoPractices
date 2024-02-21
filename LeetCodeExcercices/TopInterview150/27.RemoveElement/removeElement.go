//isolated solution
func removeElement(nums []int, val int) int {
	founds := 0
	numLen := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			founds++
			i--
		}
	}
	return numLen - founds
}