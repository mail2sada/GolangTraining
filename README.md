# GolangTraining
This will cover day wise Golang training content.

# Getting Started with Golang.

### Overview
	Introduction to Golang
	Installing Golang 
	Hello World! in Golang
	Golang vs C++
	Golang vs Java

### Packages and Modules in Go
	Overview
	Modules
	Understanding Package and Also Creating a module
	Creating an application before Module
	Exported vs Unexported Names
	Nested Packages
	Aliasing in importing packages
	Init Functions
	Order of execution of a Go program
	Blank Identifier in import
	Package Naming Convention
	Types of Modules
	Package vs Module
	Add a dependency to your project
	Directly adding it to the go.mod file
	Do a go get
	Add the dependency to your source code and do a go mod tidy
	Adding a vendor directory
	Module Import Path
	Importing package within same module
	Importing package from different module locally
	Selecting the version of library
	Differ in minor or patch version	
	Differ in major version
	go mod command
	Direct vs Indirect Dependencies in go.mod file

### Fundamentals
	Identifiers in Golang
	Keywords in Golang
	Data Types
	Variables
	Constants
	Rune in Golang
	Operators in Golang
	Scope of Variables
	Type Casting
	var Keyword in Golang
	Short Declaration Operator(:=)
	var keyword vs short declaration operator
	
### Control Statements
	Decision Making Statements
	Loops in Golang
	Loop Control Statements
	Switch Statement in Go
	Deadlock and Default Case in Select Statement


### Functions & Methods
	What are the Functions?
	Variadic Function
	Anonymous Function
	main and init function
	Function Arguments
	Function Returning Multiple Values
	Named Return Values
	Blank Identifier
	Defer
	Methods
	Methods With Same Name
	Panic and recover
	
### Error Handling
	Overview
	Using type which implements error interface
	Advantages of using error as a type
	Different ways of creating an error
	Using errors.New(“some_error_message”)
	Using fmt.Errorf(“error is %s”, “some_error_message”).
	Creating Custom error
	Ignoring errors
	Wrapping of error
	Unwrap an error
	Check if two error are equal
	Using the equality operator (==)
	Using the Is function of errors package
	Get the underlying error from an error represented by the error interface
	Using the .({type}) assert
	Using the As function of errors package
	Runtime Error Panic
	Calling the panic function explicitly
	Panic with defer
	Recover in golang
	Panic/Recover and Goroutine
	Printing stack trace
	Return value of the function when panic is recovered
	
### String
	Strings in Golang
	Different ways to compare Strings
	Different ways to concatenate two strings
	Trimming a String in Golang
	Splitting a String in Golang
	Check if the given characters are present in String
	Repeating a String for Specific Number of Times
	Finding the index value of specified string
	Counting the Number of Repeated Characters in String
### Structure
	Structures
	Structure Equality
	Nested Structure
	Anonymous Structure and Fields
	Promoted Fields in Structure
	Promoted Methods in Structure
	Function as a Field in Structure
	
### Arrays & Slices
	Arrays in Golang
	Copying an Array into Another Array in Golang
	Passing an Array to a Function in Golang
	Slices in Golang
	Slice Composite Literal
	Copying one Slice into another Slice
	Passing a Slice to Function
	Comparing two Slices in Golang
	Checking the Equality of Slices in Golang
	Sorting a Slice in Golang
	Trimming a Slice in Golang
	Splitting a Slice in Golang
	Slice Sort, Reverse, Search Functions
	
### Maps
	Overview
	Allowed Key types in a Map
	Allowed Value types in a Map
	Creating a Map
	Using the map[<key_type>]<value_type> format
	Using Make
	Map Operations
		Add a key value pair
		Update a key-value pair
		Get the value corresponding to a key
		Delete a key value pair
		Check if a key exists
		Functions on Maps
	Zero Value
	Maps are referenced data types
	Iterate over a map
	Maps are not safe for concurrent use
	Slice Sort, Reverse, Search Functions

### Pointers
	Pointers
	Pointer to Pointer (Double Pointer)
	Pointers to a Function
	Returning Pointer from a Function
	Pointer to an Array as Function Argument
	Pointer to Struct
	Comparing Pointers
	Finding the Capacity of the Pointer
	Finding the Length of the Pointer

### Interfaces
	Interfaces
	Multiple Interfaces
	Embedding Interfaces
	Polymorphism Using Interfaces


### Concurrency
	Goroutines – Concurrency
	Select Statement
	Multiple Goroutines
	Goroutine vs Thread
	Channel in Golang
	Unidirectional Channel in Golang
	Synchronization in Golang 
		WaitGroup
		Mutex

### Date and Time
	Overview
	Structure
	Create a new time
	Using time.Now()
	Using time.Date()
	Understanding Duration
	Add or Subtract to a time
	Add to time
	Subtract to time
	Time Parsing/Formatting
	Time Parse Example
	Time Formatting Example
	Time Diff
	Time Conversion
	Convert time between different timezones
	Timers
	Tickers

### Context Package
	Overview
	Context Interface
	Creating New Context
		context.Background()
		context.ToDo()
	Context Tree
		Two level tree
		Three level tree
		Multi level tree
	Deriving From Context
		context.WithValue()
		context.WithCancel()
		context.WithTimeout()
		context.WithDeadline()
	Best Practices

### Basic Http Project Implementation
	Overview
	Request
	Response
	Pair of API signature and its handler
	Mux
	Listener
		Using server’s ListenAndServe function
		Using http's ListenAndServe function
	GRPC and JSON

	

