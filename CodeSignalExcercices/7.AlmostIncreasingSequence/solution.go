func solution(sequence []int) bool {
	countRemoveNum := 0
	lenSeq := len(sequence)

	for i := 1; i < lenSeq; i++ {
		if sequence[i-1] >= sequence[i] {
			countRemoveNum++
			if countRemoveNum > 1 {
				return false
			}
			if i > 1 && i < lenSeq-1 
				&& sequence[i-2] >= sequence[i] 
				&& sequence[i-1] >= sequence[i+1] {
				return false
			}
		}
	}

	return true
}
