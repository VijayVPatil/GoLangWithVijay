package main

import "fmt"

//No Inheritance in Golang, No Super or Parent or Child

//Defining Structs
// AS it should be accessible outside, We give first letter
// as Capital
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	//Structs

	vijay := User{"Vijay", "vj@gmail.com", true, 23}

	fmt.Println(vijay)
	fmt.Printf("Details are : %+v\n", vijay) //%+v gives more detail
	fmt.Printf("Name is : %v , Email is %v , Status is %v , age is %v", vijay.Name, vijay.Email, vijay.Status, vijay.Age)

}
