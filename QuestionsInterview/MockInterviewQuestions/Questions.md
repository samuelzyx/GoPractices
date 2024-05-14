Tell me more about Golang?

- Is a language it used normally by the easy to undertand of the sintaxis that use.
- It's was desinged by Google en 2009.
- It's open source.
- It's was based in C.
- It's can handle concurrency with the GoRoutines, Channels, Wait Groups.
- It's have great performance of memory, lightweight processes and use Garbage Collector similar like the Java JVM but don't use a VM.

How run a program in Golang?

- After install the Golang version requiered in the SO, we can go to the directory were is allowed the go.mod file and execute a command: > go mod tidy and/or execute command: > go run main.go (or name of the file that we want to run)

What is a interface in Golang?

- Is similar that the interfaces in Java, but is a type (interface) that can allow multiple methods signatures that only define the name of them, the params that will be recieve or return but only that and in the implementation it will be describe all the logic inside the functions, in that way the interface can helps us to define certain contract of functions that the implementation needs to have.

What are the objects in Golang?

- The objects are the instance of a data type

What are the types in Golang?

- The types are the templates that defines the structure and behavior of a value. Can be classified into Basic types or Compound types:
   * Basic types: these are the predefined types, such as int, float64, string or bool
   * Compound types: these are created by the user, such structs, slices, arrays and maps.

What is the difference beetween Slice and Array in Golang?

- The array length is fixed and in slice are dinamically
- Arrays are passing throw functions by copy and the slices by reference
- In arrays we can modify the values of the elements but not decrement or increment the length of elements in the array, and in the slice yes we can do that
- Slices have property capacity and arrays not

Difference beetween capacity and lenght in golang slices?

- The capacity is the size of the memory reference of the slice with reallocating memory also if the capacity is exceed this will increase at the double of the existing size
- The lenght is the current size of accesible elements of the slice

How can we have the type of an object?

- Using package fmt with the %T format in the sintaxis of the print: fmt.Printf("Type of variable: %T\n",variable) //Output: Type of variable: string
- Using package reflect: reflect.TypeOf(variable)
- Using assertions we can have a boolean that the conversion is was successful or not   
	s, ok := i.(string)
	fmt.Println(s, ok)

Where we will use square breakers or curly breakers?

- Square breakers are common used into arrays to assing or access values
- Curly breakers are common used to assing multiple variables, or initialized a slice

How works the garbage collector in Golang?

- Gargabe collector is a automated system that manages memory allocation and deallocation during the program execution, it's help us to prevent memory leaks and ensure efficient memory usage by reclaimed unused memory that is not longer reference by the program execution

What are tags and the purpose of them?

- The tags are used to associate metada to the variables of the struct definition for example were we want to specify the decoded format like JSON or XML in the fields or for some libriaries that helps to validate fields like formats that are requireds or with minimum sizes or so on
- The purpose of them is indicate to the compiler or the library or tool that we want to provide additional information metadata to the fields

How you can compare two objects in golang?

- We can use comparison oparators like equals == , not equals !=, less than <, great than > and so on.
- We can compare arrays with reflect.DeepEqual() function to compare each value of the array
   Syntax: func DeepEqual(x, y interface{}) bool
   Parameters: This function takes two parameters with value of any type, i.e. x, y.
   Return Value: This function returns the boolean value. 
- We can use for loops in order to compare each value from same length size or slices or arrays
- We can use == operator with structs comparing each field with the other field of the struct

How works maps, the limitations that the have?

- Maps are defined by a unorder collection of key and value pairs elements
- We can defined a new maps with make(map[typeofkey]typeofvalue)
- We can access to the values of the list by the key reference map[key]
- We can increment or decremente the collection of elements
- The limitations can be the unorder collection, the unique key for entired collection, there are some limited types to use as key

What's does the command 'go build'?
The go command is used to compile packages and dependencies in the current directory and generate executable files or object files.

What's does the command 'go test'?
The go command is used to run all the test files generated in the project and obtain all the results

What's does the command 'go run'?
The go command is used to execute the current project or main file that execute the project

What's does the command 'go clean'?
The go command is used to remove object files and catched files generated by the 'go build', 'go install' and 'go test' commands. It helps us to cleaning up the workspace by removing any generated files that are not needed anymore.

What's a empty interface?  
An empty interface, denotad by the 'interface{}' type, is an interface with zero methods. It's most general interface type because it does not specify any method signatures, witch mean any types satisfies an empty interface. Essentially, an empty interface can hold any value of any type.

What's a variadic function?
A variadic function is a function that can accept a variable number of arguments of a specific type. This means you can pass zero or more arguments of the specified type to the function when calling it. The sintaxis for defining a variadic function in Go is to append '...' to the type of the last parameter of the function. This tells the compiler that the function can accept multiple values of that type.

func sum(nums ...int) int{
   total := 0
   for _, num := range nums{
      total += num
   }
   return total
}

In which layer do you normally puts the router and the handlers?
In the controller or can be separatly in the file handler

What means that a variable is not exported?
Means that the variable name starts with lowercase and are only visible/accesible by the package where they are defined. The cannot be accesed or used by code outside of the package

package mypackage

// Exported variable
var ExportedVariable int = 42

// Unexported variable
var unexportedVariable int = 10

It's true that if you declare a variable without value golang assign a value by default? 
Yes, for numeric types it will initialized with zero value, if is boolean it will be false, string empty string, pointer it will nil, same for slice, map and channel it will be nil

A function can return multiple values?
Yes is possible you can specified how many values and the types of the values

What is a annonymus function?
Is a function that dont have a name is normally defined with func(a int)b int{}

What is a receiver parameter?
It is the variable/parameter used at the beginning of the method to indicate to which structure it will be associated, In other words, it will refer to the structure to which the method coming from an interface will be associated.

What does the defer keyword?
It is to tell golang to execute the lines of code just before the end of the complete function. It is commonly used with goroutines to close channels.

