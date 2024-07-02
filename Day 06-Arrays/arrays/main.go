package main

import "fmt"

func main() {
	//ARRAYS in GoLang

	//First Way to declare Array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println(fruitList)

	//Second Way to declare Array
	var veglist = [3]string{"Potato", "Tomato", "Cabbage"}

	fmt.Println(veglist)
}
