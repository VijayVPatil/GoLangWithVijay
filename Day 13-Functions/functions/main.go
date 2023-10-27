package main

import "fmt"

func main() {
	//Functions in GoLang

	//Example 1
	// greeter()

	//Example 2
	// result := operation(1, 3)
	// fmt.Println(result)

	//Example 3
	// result1 := adder(1, 2, 3, 4, 3)
	// fmt.Println(result1)

	//Example 4

	result, message := add(2, 3, 35, 45)
	fmt.Println(result)
	fmt.Println(message)
}

//--------------------------------------
// Example 1
//-----------------------------------
// func greeter() {
// 	fmt.Println("Hello")
// }

//---------------------------------
//Example 2
//--------------------------------

// func operation(val1 int, val2 int) int {
// 	return val1 + val2
// }

//-------------------------
//Example 3 Variadic Parameters
//---------------------------

// func adder(values ...int) int {
// 	total := 0

// 	for val := range values {
// 		total += val
// 	}

// 	return total
// }

//----------------------------------
//Another way to do Example 3
//-----------------------------------
// func adder(values ...int) int {
// 	total := 0

// 	for val, _ := range values {
// 		total += val
// 	}

// 	return total
// }

//Returning Mutiple parameters

func add(value ...int) (int, string) {

	total := 0

	for val, _ := range value {
		total += val
	}

	return total, "GoLangWithVijay"
}
