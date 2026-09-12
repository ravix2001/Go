package main

import "fmt"

func display(str string){
	for i:=0; i<5; i++ {
		fmt.Println(str)
	}
}

func main(){
	go display("Hello Go Routine!")

	display("Hello Main!")
}

// go display(...) runs in a separate goroutine
// display("Hello, Main!") runs in the main goroutine
// The program may exit before the goroutine finishes