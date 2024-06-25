package main

import "fmt"

// isolated solution
func removeElement(nums *[]int, val int) int {
	founds := 0
	numLen := len(*nums)
	for i := 0; i < len(*nums); i++ {
		if (*nums)[i] == val {
			(*nums) = append((*nums)[:i], (*nums)[i+1:]...)
			founds++
			i--
		}
	}
	return numLen - founds
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	remove := 2

	fmt.Println("Nums", nums)
	fmt.Println("Remove", remove)

	result := removeElement(&nums, remove)

	fmt.Println("Result", result)
	fmt.Println("Nums", nums)
}
