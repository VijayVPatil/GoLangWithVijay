package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Number int
	Status bool
}

func main() {

	vijay := User{"Vijay", "vj@slack", 12, true}

	//Calling the method
	vijay.GetStatus()

	vijay.NewName()
	// fmt.Printf("User is %+v", vijay.Name)

}

// Defining methods in golang
func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

//Setting new value

func (u User) NewName() {
	u.Name = "Tushar"
	fmt.Println("New Name is :", u.Name)
}
