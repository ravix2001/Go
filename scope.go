package main

import "fmt"

// global level scope
var myvariable1 = 100

func main() {

	fmt.Println(myvariable1)
	
	var myvariable2 int

	fmt.Println(myvariable2)
	
	myvariable3 := 200

	fmt.Println(myvariable3)
    
}