package main

import "fmt"

// init() Function
// The init() function is executed automatically when a package is initialized, before the main() function runs.

// It is used for tasks such as:

// Initializing global variables
// Loading configuration values
// Setting up database connections
// Registering dependencies

func init(){
	fmt.Println("Welcome to init() function")
}

func init(){
	fmt.Println("Database initialization")
}

// The main() function is the entry point of an executable Go program. It is automatically executed when the program starts and is defined inside the main package.

// Must be declared inside the main package.
// Does not accept any parameters.
// Does not return any value.
// Is automatically executed by Go.
// There can be only one main() function in an executable program.

func main(){
	fmt.Print("Welcome to main() function")
}