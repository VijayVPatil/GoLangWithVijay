package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//User input program
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	//comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("The value is: ", input)
	fmt.Printf("Type of the value is %T", input)
}
