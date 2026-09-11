package main

import "fmt"

// In the below program, the func calculator(a, b int) (mul int, div int) line of code contains the named return arguments. 
// The return statement at the end of function doesn't contain any parameters. 
// Go compiler automatically returns the parameters.

// function having named arguments
func calculator(a int, b int) (sum int, product int) {
	// here, simple assignment will
    // initialize the values to it
	sum = a + b
	product = a * b
	// here we have return keyword
    // without any resultant parameters
	return
}

func main(){
	s, p := calculator(5, 8)
	fmt.Println("Sum = ", s)
	fmt.Println("Product = ", p)
}