package main

import "fmt"

// Anonymous functions can form closures. 
// A closure is a function that can access and modify variables from its surrounding scope, 
// even after the outer function has finished executing.

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main(){
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

// count is declared inside counter().
// The anonymous function remembers the value of count.
// Every time next() is called, count is incremented.