package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var arr1 [7]int = [7]int{7, 3, 6, 0, 4, 9, 10}
	s, _ := json.Marshal(arr1)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr1), cap(arr1))
	var arr2 []int = arr1[1:3]
	s, _ = json.Marshal(arr2)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr2), cap(arr2))
	arr2 = arr2[0 : len(arr2)+2]
	s, _ = json.Marshal(arr2)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr2), cap(arr2))
	for k := range arr2 {
		arr2[k] += 1
	}
	s, _ = json.Marshal(arr2)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr2), cap(arr2))
	s, _ = json.Marshal(arr1)
	fmt.Printf("%s\n", string(s))

	var arr3 []int = []int{1, 2, 3}
	s, _ = json.Marshal(arr3)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 4)
	s, _ = json.Marshal(arr3)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 5)
	s, _ = json.Marshal(arr3)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr3), cap(arr3))

	arr4 := make([]int, 3, 9)
	s, _ = json.Marshal(arr4)
	fmt.Printf("%s\n", string(s))
	fmt.Printf("Len: %d, Capacity: %d\n", len(arr4), cap(arr4))

}
