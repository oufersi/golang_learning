package main

import(
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Alice")
	fmt.Println(message)
	fmt.Println(Get_rand_num() + 1)
}