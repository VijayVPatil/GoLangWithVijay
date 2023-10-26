package main

import "fmt"

func main() {
	//For Loop

	days := []string{"Mon", "Tue", "Sat", "Sum"}
	fmt.Println(days)

	//One way to loop a Slice
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("---------------")
	//Other way to loop

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("---------------")

	//For each
	for index, value := range days {
		fmt.Printf("The index position is %v and value at that index is %v\n", index, value)
	}
	fmt.Println("---------------")
	//For each using commaa ok syntax
	//if we want to ignore index value
	for _, value := range days {
		fmt.Printf("The index position is and value at that index is %v\n", value)

	}
	fmt.Println("---------------")

	rougueValue := 1
	//Following syntax is like while loop in other languages
	for rougueValue < 10 {
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

	//break and continue useage
	fmt.Println("Break usage -------------------")

	rougueValu := 1
	for rougueValu < 10 {
		if rougueValu == 5 {
			break
		}
		fmt.Println("Value is ", rougueValu)
		rougueValu++
	}

	fmt.Println("Continue -------------------")
	//Continue just skips the process
	rougueVal := 1
	for rougueVal < 10 {
		if rougueVal == 5 {
			rougueVal++
			continue
		}
		fmt.Println("Value is ", rougueVal)
		rougueVal++
	}

	//Go to in golang

	fmt.Println("Go to in GoLang")
	meme := 1

	for meme < 10 {
		fmt.Println("The value of meme is ", meme)
		if meme == 5 {
			goto hey
		}
		meme++
	}

hey:
	fmt.Println("Hello")

}
