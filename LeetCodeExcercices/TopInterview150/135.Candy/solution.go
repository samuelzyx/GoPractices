func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	totalCandies, upDiff, downDiff, peak := 1, 0, 0, 0
	fmt.Printf("TotalCandys:%d\n", totalCandies)
	for i := 0; i < n; i++ {
		fmt.Printf("\nCurrentRating:%d\n", ratings[i])
		if (i + 1) < n {
			fmt.Printf("NextRating:%d\n", ratings[i+1])
			currentRating, nextRating := ratings[i+1], ratings[i]
			if currentRating > nextRating {
				fmt.Print("CurrentRating > NextRating\n")
				upDiff++
				downDiff = 0
				peak = upDiff
				totalCandies += 1 + upDiff
				fmt.Printf("Increment upDiff:%d\n", upDiff)
				fmt.Printf("downDiff 0:%d\n", downDiff)
				fmt.Printf("Peak equals upDiff:%d\n", peak)
				fmt.Printf("Increment totalCandies 1+upDiff:%d\n", totalCandies)
			} else if currentRating == nextRating {
				fmt.Print("CurrentRating == NextRating\n")
				upDiff, downDiff, peak = 0, 0, 0
				totalCandies++
				fmt.Printf("Increment totalCandies:%d\n", totalCandies)
			} else {
				fmt.Print("CurrentRating < NextRating\n")
				upDiff = 0
				downDiff++
				totalCandies += 1 + downDiff
				fmt.Printf("upDiff 0:%d\n", upDiff)
				fmt.Printf("Increment downDiff:%d\n", downDiff)
				fmt.Printf("Increment totalCandies 1+downDiff:%d\n", totalCandies)
				if peak >= downDiff {
					totalCandies--
					fmt.Printf("If peak>=downDiff, Decrement totalCandies:%d\n", totalCandies)
				}
			}
		}
	}
	fmt.Printf("TotalCandys:%d\n", totalCandies)
	return totalCandies
}
