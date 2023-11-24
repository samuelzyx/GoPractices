package main

import "fmt"

//Example 2: You can also modify the values of the structure members or
//structure literals by using the pointer as shown below:

// taking a structure
type Employee struct {

	// taking variables
	name  string
	empid int
}

// Main Function
func main() {

	// creating the instance of the
	// Employee struct type
	empInstance := Employee{"Sam", 170390}

	// Here, it is the pointer to the struct
	empPointer := &empInstance

	// displaying the values
	fmt.Println(empPointer)

	// updating the value of name
	empPointer.name = "Samuel Ivan"

	fmt.Println(empPointer)
}
