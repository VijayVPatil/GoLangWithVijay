package main

import "fmt"

func main() {

	//Creating Pointers
	var ptr *int
	fmt.Println(ptr)

	nos := 34

	nos1 := &nos
	value := *nos1
	fmt.Println("The Address is :", nos1)
	fmt.Println("The Value is :", value)
}
