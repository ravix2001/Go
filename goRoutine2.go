package main

import "fmt"
import "time"

func display(str string){
	for i:=0; i<5; i++ {
		fmt.Println(str)
	}
}

func main(){
	go display("Hello Go Routine!")

	time.Sleep(time.Second)
	display("Hello Main!")
}

// time.Sleep() pauses the main goroutine
// This gives time for other goroutines to execute