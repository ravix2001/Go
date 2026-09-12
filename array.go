// Go program to illustrate how to create
// an array using shorthand declaration 
// and accessing the elements of the 
// array using for loop
package main

import "fmt"

func main() {

	// Var array_name[length]Type
	var numbers[5] int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	fmt.Println("Numbers:", numbers)

	// array_name := [length]Type{item1, item2, item3,...itemN}
	// Shorthand declaration of array
	arr := [4]string{"Adam", "John", "Mark", "Henry"}

	// Accessing the elements of 
	// the array Using for loop
	fmt.Println("Elements of the array:")

	for i:= 0; i < 3; i++{
		fmt.Println(arr[i])
	}

}