package main

import "fmt"

// In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns. In other words, defer function or method call arguments evaluate instantly, but they don't execute until the nearby functions returns. You can create a deferred method, or function, or anonymous function by using the defer keyword.

// Syntax:

// // Function
// defer func func_name(parameter_list Type)return_type{
// // Code
// }

// // Method
// defer func (receiver Type) method_name(parameter_list){
// // Code
// }

// defer func (parameter_list)(return_type){
// // code
// }()

// Important Points:

// In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order as shown in Example 2.
// In the defer statements, the arguments are evaluated when the defer statement is executed, not when it is called.
// Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.

func mul(a int, b int) int {
	result := a * b
	fmt.Println("Result:", result)
	return 0
}

func show(){
	fmt.Println("Hello")
}

func main(){

	// Calling mul() function
    // Here mul function behaves
    // like a normal function
	mul(10,20)

	// Calling mul()function
    // Using defer keyword
    // Here the mul() function
    // is defer function
	defer mul(10,40)

	show()

}

// First, we call mul function normally(without the defer keyword), i.e, mul(10,20) and it executes when the function is called(Output: Result : 200 ).
// Second, we call mul() function as a defer function using defer keyword, i.e, defer mul(10,40) and it executes(Output: Result: 400 ) when all the surrounding methods return.