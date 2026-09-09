package main

import "fmt"

func double(num int) (int){
	num *=2
	// Returning updated value
	return num
}

func main(){
	x := 5

	result := double(x)
	fmt.Println("Double of x = ", result)
	
}