package main

import "fmt"

func double(num *int){
	*num *=2
}

func main(){
	x := 5

	// double(x)
	// fmt.Println("Double of x = ", x)
	// it will not give the updated value so pointer is used so that we can pass address of the value and update the value at that particular address

	// Using address (Pointer)
	double(&x)
	fmt.Println("Double of x = ", x)

}