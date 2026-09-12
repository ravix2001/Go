package main

import "fmt"
import "time"

func task1(ch chan string){
	ch <- "Task 1 completed"
}

func task2(ch chan string){
	// Adding delay to change the time of execution 
	time.Sleep(time.Second)
	ch <- "Task 2 completed"
}

func main(){

	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	select {
	case msg1 := <- ch1:
		fmt.Println(msg1)
	case msg2 := <- ch2:
		fmt.Println(msg2)
	}
}