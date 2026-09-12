package main

import "fmt"

type Employee struct {
    firstName, lastName string
    age, salary int
}

func main() {

    // passing the address of struct variable
    // emp is a pointer to the Employee struct
    emp := &Employee{"Sam", "Anderson", 55, 6000}

    // (*emp).firstName is the syntax to access
    // the firstName field of the emp struct
    fmt.Println("First Name:", (*emp).firstName)
    fmt.Println("Age:", (*emp).age)

	// The Golang gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field
	fmt.Println("First Name: ", emp.firstName)
    fmt.Println("Age: ", emp.age)
}