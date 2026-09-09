package main

import "fmt"

func divide(a, b int) (int, int){
	quotient := a/b
	remainder := a%b

	return quotient, remainder
}

func main(){
	quotient, remainder := divide(10,3)
	fmt.Println("Quotient = ", quotient)
	fmt.Println("Remainder", remainder)
}