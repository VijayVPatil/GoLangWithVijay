package main

import "fmt"

func addFloat(a, b float64) float64 {
	return a + b
}

func addInt(a, b int) int {
	return a + b
}

func addTypes[T int | float64](a, b T) T {
	return a + b
}

func main() {

    type def

	fmt.Println(addInt(2, 3))
	fmt.Println(addFloat(2.2, 3.3))

	//Below is using a generic function with type constraint

	fmt.Println(addTypes(2, 3))
	fmt.Println(addTypes(2.2, 3.3))
}
