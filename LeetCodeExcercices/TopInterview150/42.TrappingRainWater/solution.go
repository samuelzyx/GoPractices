func trap(height []int) int {
	n := len(height)
	unitWater := 0
	peak := -1
	peakIndex := 0
	prevPeak := 0
	possibleUnitWater := 0
	h := 0
	fmt.Printf("peak:%d\n", peak)
	fmt.Printf("prevPeak:%d\n", prevPeak)
	for i := 0; i < n; i++ {
		h = height[i]
		fmt.Printf("\nh:%d\n", h)

		fmt.Printf("Line 16 (h>=peak):%v (%d>=%d)\n", (h >= peak), h, peak)
		if h >= peak {
			peak = h
			peakIndex = i
			fmt.Printf("peak:%d\n", peak)
			fmt.Printf("peakIndex:%d\n", peakIndex)
			fmt.Printf("Line 21 (prevPeak==0):%v (%d==0)\n", (prevPeak == 0), prevPeak)
			if prevPeak == 0 {
				fmt.Printf("prevPeak:%d\n", prevPeak)
				prevPeak = h
			} else {
				fmt.Printf("Line 26 (prevPeak<=h):%v (%d<=%d)\n", (prevPeak <= h), prevPeak, h)
				if prevPeak <= h {
					unitWater += possibleUnitWater
					fmt.Printf("unitWater:%d\n", unitWater)
					possibleUnitWater = 0
					fmt.Printf("possibleUnitWater:%d\n", possibleUnitWater)
					prevPeak = h
					fmt.Printf("prevPeak:%d\n", prevPeak)
				}
			}
		} else {
			fmt.Printf("Line 38 ((prevPeak!=0) && (i!=n-1)):%v (((%d!=0) && (%d!=%d-1))\n", ((prevPeak != 0) && (i != n-1)), prevPeak, i, n)
			if (prevPeak != 0) && (i != n-1) {
				possibleUnitWater += prevPeak - h
				fmt.Printf("possibleUnitWater:%d\n", possibleUnitWater)
			} else {
				fmt.Printf("Line 43 (possibleUnitWater != 0):%v (%d!=0)\n", (possibleUnitWater != 0), possibleUnitWater)
				if possibleUnitWater != 0 && i == n-1 && prevPeak != h {
					possibleUnitWater = 0
					fmt.Printf("possibleUnitWater:%d\n", possibleUnitWater)
					prevPeak--
					fmt.Printf("prevPeak:%d\n", prevPeak)
					i = peakIndex + 1
					fmt.Printf("i:%d\n", i)
					peakIndex++
					fmt.Printf("prevPeak:%d\n", prevPeak)
					peak = prevPeak
					fmt.Printf("peak:%d\n", peak)
					fmt.Printf("/////////////////Reset\n")
				}
			}
		}
	}
	return unitWater
}