package main

import "fmt"

func main() {

    // taking a normal variable
    var x int = 5748
    
    // declaration of pointer
    var p *int
    
    // initialization of pointer
    p = &x

    // displaying the result
    fmt.Println("Value stored in x = ", x)
    fmt.Println("Address of x = ", &x)
    fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Value pointed by variable p (*p) = ", *p)
	fmt.Println("Address of p = ", &p)

	*p = 500

	fmt.Println("Updated Value stored in x = ", x)
	fmt.Println("Updated Value pointed by variable p (*p) = ", *p)

}