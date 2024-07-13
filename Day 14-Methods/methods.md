# Methods in Golang

1. In Golang , methods are similar to function , it has a name , takes arguments as input and also return value. But it also comes with a receiver argument.
2. Methods helps to access the properties of reciever. Reciever can be a struct or non struct type.
3. When you create a method in your code the receiver and receiver type must be present in the same package.

How to define a method

```
func(reciever_argument_name  type_of_reciever)method_name(paramters_list)(returns){

}

```

Example :

#### code 1.01

```
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


```

### Explannation of code 1.01

1. We can call the method by using the variable of struct declared
2. The methods have a reciever argument that can have a struct and non struct type. AS in above example we have a taken a Student struct type.
3. We can also have a pointer reciver in methods if we want to make direct changes to the struct fields.
4. We can call a function that takes a argument as pointer to struct and a value. Functions with a pointer argument must take a pointer.
