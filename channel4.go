package main

import "fmt"

func main(){

	channel := make(chan string, 5)

	channel <- "Ravi"
	channel <- "Pandit"
	
	fmt.Println("Length of the channel is ", len(channel))
	fmt.Println("Capacity of the channel is ", cap(channel))

}