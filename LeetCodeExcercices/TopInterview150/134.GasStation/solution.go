func rotate(nums *[]int, startIndex, endIndex int) {
	for startIndex < endIndex {
		(*nums)[startIndex], (*nums)[endIndex] = (*nums)[endIndex], (*nums)[startIndex]
		startIndex++
		endIndex--
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	var startIndex int
	var ng = len(gas)

	for c := 0; c < ng; c++ {
		startIndex = c
		var ngas = make([]int, ng)
		var ncost = make([]int, ng)

		ngas = gas
		ncost = cost
		c %= ng

		rotate(&ngas, 0, ng-1)
		rotate(&ngas, 0, c-1)
		rotate(&ngas, c, ng-1)

		rotate(&ncost, 0, ng-1)
		rotate(&ncost, 0, c-1)
		rotate(&ncost, c, ng-1)

		fmt.Printf("StarIndex %v\n", startIndex)
		fmt.Print("ngas\n")
		fmt.Print(ngas)
		fmt.Print("\n")
		fmt.Print("ncost\n")
		fmt.Print(ncost)
		fmt.Print("\n")

		tank := 0
		for r := 0; r < ng; r++ {
			if tank == 0 {
				tank += ngas[r]
			} else {
				fmt.Printf("Tank:%d-Cost[%d]:%d+Gas[%d]:%d\n", tank, r, ncost[r], r, ngas[r])
				tank = tank - ncost[r] + ngas[r]
				fmt.Printf("Tank:%d\n", tank)
				if tank <= 0 {
					fmt.Print("Can't complete circuit\n")
					break
				}
			}
			fmt.Print("Tank: ")
			fmt.Print(tank)
			fmt.Print("\n")
		}
		if (tank - ncost[0]) >= 0 {
			return startIndex + 1
		}
	}

	return -1
}