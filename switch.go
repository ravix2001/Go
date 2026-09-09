package main

import "fmt"

func main(){
	day := 3

	switch day {
	case 1:
		fmt.Println("Sunday")
		
	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	default:
		fmt.Println("Invalid")
	}
}