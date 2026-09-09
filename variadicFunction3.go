package main

import "fmt"

// greeting is a regular parameter.
// names is a variadic parameter.
// The for loop iterates through all names.
// The variadic parameter must always appear last in the function definition.
// Variadic functions can only have one variadic parameter, and it must be the last parameter.
// Cannot have multiple variadic parameters in a single function definition.

func greet(greeting string, names ...string){
	for _, name := range names{
		fmt.Println(greeting, name)
	}
}

func main(){
	greet("Hello", "John", "Mark")
	greet("Welcome", "Ravi")
}
