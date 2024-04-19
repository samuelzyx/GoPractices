Read the next code and write what will be the result
///
package main

import "fmt"

func main() {
	sports := make([]string, 5)
	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo"

	xs := sports[1:3:3]
	xs[0] = "CHANGED"

	inspectSlice(sports)
	inspectSlice(xs)
}

func inspectSlice(xs []string) {
	fmt.Printf("\nlen: %v, \ncap: %v, \n", len(xs), cap(xs))
	for i := range xs {
		fmt.Printf("%p \t %v \n", &xs[i], xs[i])
	}
}
///

Answer 

len: 5, 
cap: 5, 
0xc000100050 	 ski 
0xc000100060 	 CHANGED 
0xc000100070 	 swim 
0xc000100080 	 sail 
0xc000100090 	 sumo 

len: 2, 
cap: 2, 
0xc000100060 	 CHANGED 
0xc000100070 	 swim 

Program exited.