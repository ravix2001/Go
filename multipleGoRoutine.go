package main

import "fmt"
import "time"

func callName(names []string){
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func callMarks(marks []int){
	for j := 0; j < len(marks); j++{
		// time.Sleep(10 * time.Millisecond)
		fmt.Println(marks[j])
	}
}

func main(){
	fmt.Println("=====Main Go Routine Start=====")

	names := [] string{"Ravi", "John", "Mike"}

	marks := [] int{10, 12, 13}

	go callName(names)
	go callMarks(marks)

	time.Sleep(time.Second)
	fmt.Println("=====Main Go Routine End=====")
}