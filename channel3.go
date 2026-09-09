package main

import "fmt"

func myFunc(ch chan string){
	for i := 0; i<5; i++ {
		ch <- "Hello, Ravi"
	}
	close(ch)
}

func main(){

	// creating channel
	channel := make(chan string)

	// calling GoRoutine
	go myFunc(channel)

	for {
		res, status := <-channel

		if status == false {
			fmt.Printf("Channel closed (%t)\n", status)
			break
		}else{
			fmt.Printf("Channel open: %s (%t)\n", res, status)
		}
	}
}