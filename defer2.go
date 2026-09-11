package main

import "fmt"

// Go program to illustrate
// multiple defer statements, to illustrate LIFO policy

func add(a1, a2 int) int {
    res := a1 + a2
    fmt.Println("Result: ", res)
    return 0
}

func main(){

	fmt.Println("Start")

	defer fmt.Println("End")
	defer add(10,20)
	defer add(30,50)
}