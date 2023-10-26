package main

import "fmt"

func main() {
	//MAPS
	//We can make use of make key word to declare a map .
	// languages :=make(map[key]value)
	//In place of key and value we give our datatype
	languages := make(map[string]string)

	//Add Values in Maps
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("The Map is :", languages)

	//Prining Individual Values
	fmt.Println("Long Form of JS is :", languages["JS"])
	fmt.Println("Long Form of RB is :", languages["RB"])
	fmt.Println("Long Form of PY is :", languages["PY"])

	//Delete individual Values
	delete(languages, "RB")
	fmt.Println("The Map is :", languages)

	//Loops in MAPS
	for key, value := range languages {
		fmt.Printf("For Key %v ,value is %v\n", key, value)
	}

	//If we only want values and not key we can use comma ok syntax

	for _, value := range languages {
		fmt.Printf("For Key v ,value is %v\n", value)

	}
}
