package main

import "fmt"

func main() {
	var quantity int
	var flag bool
	var length, width float64
	var customerName string

	fmt.Println("Zero value from quantity", quantity)
	fmt.Println("Zero value from flag", flag)
	fmt.Println("Zero value from length", length, "and with", width)
	fmt.Println("Zero value from customerName", customerName)

	quantity = 4
	customerName = "Homedepot"
	flag = true
	length, width = 1.2, 1.4

	fmt.Println("Customer Name:", customerName)
	if flag {
		fmt.Println("He has ordered", quantity, "sheets")
	} else {
		fmt.Println("He hasn't ordered", quantity, "sheets")
	}
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

}
