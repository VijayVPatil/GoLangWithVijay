package main

import "fmt"

func main() {
	//Switch Case in Golang

	diceNumber := 1

	fmt.Println("Value of Dice is ", diceNumber)

	// Switch CAse syntax
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the Dice Again")
	default:
		fmt.Println("What was that?")
	}
}
