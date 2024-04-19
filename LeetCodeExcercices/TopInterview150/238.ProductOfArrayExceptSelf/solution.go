func productExceptSelf(nums []int) []int {
	var n = len(nums)

	if n == 0 {
		return []int{}
	}

	var (
		i, temp = 1, 1
	)

	var prod = make([]int, n)

	for j := 0; j < n; j++ {
		prod[j] = 1
	}

	for i = 0; i < n; i++ {
		prod[i] = temp
		temp *= nums[i]
	}

	temp = 1

	for i = n - 1; i >= 0; i-- {
		prod[i] *= temp
		temp *= nums[i]
	}

	return prod
}