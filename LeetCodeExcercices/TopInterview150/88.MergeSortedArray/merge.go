//isolated solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	o := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[o]
		o++
	}
	slices.Sort(nums1)
	return
}