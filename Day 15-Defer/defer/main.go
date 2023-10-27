package main

import "fmt"

func main() {

	// Example 1)
	// defer fmt.Println("How are you")
	// fmt.Println("Hello")

	//Example 2)
	// defer fmt.Println("How are you")
	// defer fmt.Println("Hi")
	// defer fmt.Println("Everyone")
	// fmt.Println("Hello")

	defer fmt.Println("How are you")
	defer fmt.Println("Hi")
	defer fmt.Println("Everyone")
	fmt.Println("Hello")
	Defered()
}

func Defered() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
