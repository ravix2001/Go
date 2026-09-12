package main

import "fmt"

type User struct{
	name string
	address string
	pincode int
}

func main(){
	user1 := User{}
	fmt.Println(user1)

	user2 := User{"Ravi", "Biratnagar", 56613}
	fmt.Println(user2)

	user3 := User{name: "Ravi", address: "Biratnagar", pincode: 56613}
	fmt.Println(user3)

	// user4 := User{"Biratnagar", 56613}
	// fmt.Println(user4)

	user5 := User{address: "Biratnagar", pincode: 56613}
	fmt.Println(user5)

	// Accessing struct fields
    // using the dot operator
    fmt.Println("User Name: ", user2.name)
    fmt.Println("User Address: ", user2.address)
	fmt.Println("User Pincode: ", user2.pincode)

    // Assigning a new value
    // to a struct field
    user2.address = "Dharan"
	fmt.Println("User Address: ", user2.address)
}