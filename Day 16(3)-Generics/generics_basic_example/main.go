package main

import "fmt"

func valueFloat(a float64) float64 {
	return a
}

func valueInt(b int) int {
	return b
}

func valueTypes[T any](value T) T {
	return value
}

func main() {
	fmt.Println(valueInt(2))
	fmt.Println(valueFloat(3.3))

	//Below is using a generic function

	fmt.Println(valueTypes(2))
	fmt.Println(valueTypes(3.3))
}
