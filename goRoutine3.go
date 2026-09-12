package main

import "fmt"
import "time"

func main(){
	go func (str string){
		for i:=0; i<5; i++ {
			fmt.Println(str)
			time.Sleep(500 * time.Millisecond)
		}
	}("Hello Go Routine")

	time.Sleep(time.Second)

	fmt.Println("Hello Main")
}