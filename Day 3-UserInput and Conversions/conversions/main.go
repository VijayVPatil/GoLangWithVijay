package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Give your number:")

	//User input program
	reader := bufio.NewReader(os.Stdin)

	//comma ok syntax
	input, _ := reader.ReadString('\n')

	//Here we need to add one to given user Input number
	//we need to convert input which is of type string to numeric
	number1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to number", number1+1)
	}
}
