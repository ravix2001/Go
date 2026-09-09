package main

import "fmt"

func double(num int){
	num *=2
}

func main(){
	x := 5
	double(x)
	fmt.Println("Double of x = ", x)
	// it will not give the updated value so pointer is used to update the value at that particular location
}