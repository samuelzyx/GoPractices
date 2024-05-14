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

- Is similar that the interfaces in Java, but is a type (struct) that can allow multiple functions that only define the name of them the params that will be recieve or return but only that and in the implementation it will be describe all the logic inside the functions, in that way the interface can helps us to define certain contract of functions that the implementation needs to have.

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

- The capacity is the size of the memory reference of the slice with reallocating memory
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