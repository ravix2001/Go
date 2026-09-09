package main

import "fmt"

func main() {

	// create channel using var keyword
	var channel chan int
	fmt.Println("Value of the channel: ", channel)
	fmt.Printf("Type of the channel: %T\n", channel)

	// create channel using make() function
	channel1 := make(chan int)
	fmt.Println("Value of the channel: ", channel1)
	fmt.Printf("Type of the channel: %T\n", channel1)
	
}
