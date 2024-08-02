package main

import "fmt"

func addFloat(a, b float64) float64 {
	return a + b
}

func addInt(a, b int) int {
	return a + b
}

//generic function with type constraint
func addTypes[T numbersOperation](a, b T) T {
	return a + b
}

//generic function with type constraint and type set interface
type numbersOperation interface {
	int | float64
}

func main() {
	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(2.2, 3.3))

	//Below is using a generic function with type constraint and type set interface
	fmt.Println(addTypes(2, 3))
	fmt.Println(addTypes(2.2, 3.3))
}
