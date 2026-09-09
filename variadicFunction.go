package main

import "fmt"

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