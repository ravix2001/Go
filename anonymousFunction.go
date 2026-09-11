package main

import "fmt"

func main(){
	func(){
		fmt.Println("Hello World!")
	}()

	// Assigning an anonymous function to a variable
	value := func(){
		fmt.Println("Hello Ravi!")
	}

	value()

	// Passing arguments in anonymous function
	func(name string){
		fmt.Println("Hello",name)
	}("Ravi")
}