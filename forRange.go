package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fullName := "Ravi Pandit"

	for index, char := range fullName {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	employee := map[int]string{
		101: "Alice",
		102: "Bob",
		103: "Jane",
	}

	for key, value := range employee {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}
}
