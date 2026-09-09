package main

import "fmt"

func main(){

	var num1 = 4
	var num2 int

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	
	fmt.Println("Sum = ", num1 + num2)

	var a, b, c = 10, "Ravi", true

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	fmt.Printf("a = %d (%T)\n", a, a)
	fmt.Printf("b = %s (%T)\n", b, b)
	fmt.Printf("c = %t (%T)\n", c, c)

	x, y, z := 10, "Ravi", true

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)
	

	var number, name = 10, "Ravi"

	number = 20
	name = "Ravi Pandit"

	fmt.Println("Number = ", number)
	fmt.Println("Name = ", name)

	const newNumber, newName = 30, "Mahi"

	// Cannot update the value of constants
	// newNumber =  40
	// newName = "M.S. Dhoni"

	fmt.Println("New Number = ", newNumber)
	fmt.Println("New Name = ", newName)

}