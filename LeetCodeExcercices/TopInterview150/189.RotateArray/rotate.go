func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // Handle the case where k is larger than the array length

	// Reverse the entire array
	reverse(&nums, 0, n-1)

	// Reverse the first k elements
	reverse(&nums, 0, k-1)

	// Reverse the remaining elements (n-k)
	reverse(&nums, k, n-1)
}

// Helper function to reverse a subarray from startIndex to endIndex
func reverse(nums *[]int, startIndex, endIndex int) {
	for startIndex < endIndex {
		(*nums)[startIndex], (*nums)[endIndex] = (*nums)[endIndex], (*nums)[startIndex]
		startIndex++
		endIndex--
	}
}