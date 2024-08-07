package main

import "fmt"

//We use pointer receivers in  methods.  This helps us to make changes to the instance of the struct.

type Student struct {
	id   int
	name string
}

//Following is how we write struct type recivers in methods
func (s Student) showName() {
	fmt.Println(s.name)
}

//Following is how we write pointer recievers in methods as well as take values arguments
//If we want to make a change to teh struct fields we use pointer receivers
//In changeName, the name field is accessed explicitly by dereferencing s using (*s).seats. Now we can directly use s in place of (*s)
//This is a more verbose way to do the same thing.
func (s *Student) changeName(su string) {
	s.name = su
}

//Now using functions
func changeNames(s *Student, su string) {
	s.name = su
}

func main() {

	s1 := Student{
		id:   32,
		name: "John",
	}

	//Following will print the struct and its type
	fmt.Printf("The value of struct s1 is %v and its type is %T\n", s1, s1)

	//Following will call the showName() method
	s1.showName()

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	// var v Vertex
	// v.Scale(5)  // OK
	// p := &v
	// p.Scale(10) // OK

	fmt.Println("The struct before changeName method is called", s1)
	//Follwing will call the changeName method with pointer reciever
	s1.changeName("Jake")
	fmt.Println("The struct after changeName method is called", s1)

	fmt.Println("The struct before changeNames function is called", s1)
	//Following is how we call a function that takes a argument as pointer to struct and a value
	//functions with a pointer argument must take a pointer:
	changeNames(&s1, "Jerry")
	fmt.Println("The struct after changeNames function is called", s1)
}
