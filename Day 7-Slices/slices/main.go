package main

import (
	"fmt"
	"sort"
)

func main() {
	//Slices

	// var fruitList = []string{"Apple", "Tomato", "Peach"}

	// fruitList = append(fruitList, "Mango", "Banana")

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	//Making Slices using make keyword
	highScores := make([]int, 4)

	highScores[0] = 434
	highScores[1] = 256
	highScores[2] = 745
	highScores[3] = 776

	fmt.Printf("The values are %v and type is %T \n", highScores, highScores)

	highScores = append(highScores, 45, 676, 342)
	fmt.Printf("The values are %v and type is %T \n", highScores, highScores)

	sort.Ints(highScores) // Ints sorts a slice of ints in increasing order
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) //IntsAreSorted reports whether the slice x is sorted in increasing order.

	//Declare a slice

	var slicess = []int{1, 2, 3, 4}
	fmt.Println(slicess)

	//How to remove a value from slices based on indexes

	var courses = []string{"react", "js", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

}
