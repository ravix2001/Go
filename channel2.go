package main

import "fmt"

func add(ch chan int){
	// receiving data
	data := <-ch
	fmt.Printf("Sum = %d\n", 250 + data)

	// can be directly received
	// fmt.Printf("Sum = %d\n", 250 + <-ch)
}

func main(){
	fmt.Println("Sarting main method")

	channel := make(chan int)

	go add(channel)

	// sending data
	channel <- 25

	fmt.Println("Ending main method")
}