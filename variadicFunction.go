package main

import "fmt"

// Variadic functions allow a function to accept a variable number of arguments of the same type. 
// Instead of defining a fixed number of parameters, the ... syntax is used, 
// making functions more flexible and suitable for scenarios where the number of inputs is unknown beforehand.
func sum(nums ...int) (int){
	total := 0

	for _, value := range nums{
		total += value
	}
	return total
}

func main(){
	fmt.Println("Sum = ", sum(1,2,3,4))
	fmt.Println("Sum = ", sum(1,2,3,4,5,6,7,8))
}