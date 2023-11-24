package main

import "fmt"

//Example 1: Here, we are creating a structure named Employee which has two variables.
//In the main function, create the instance of the struct i.e. emp. After that, you can
//pass the address of the struct to the pointer which represents the pointer to struct
//concept. There is no need to use dereferencing explicitly as it will give the same
//result as you can see in the below program (two times ABC).

// taking a structure
type Employee struct {

	// taking variables
	name  string
	empid int
}

// Main Function
func main() {
	// Creating the instance of the
	// Employee struct type
	empInstance := Employee{"liem", 19078}

	fmt.Println(empInstance) //&empInstance: 0xc000008048
	fmt.Println(empInstance.name)

	// Here, it is the pointer to the struct
	empPointer := &empInstance

	fmt.Println(empPointer)

	// accessing the struct fields(liem employee's name)
	// using a pointer but here we are not using
	// dereferencing explicitly

	name := empPointer.name
	fmt.Println(name)

	// same as above by explicitly using
	// dereferencing concept
	// means the result will be the same

	namePointer := (*empPointer).name
	fmt.Println(namePointer)
}
