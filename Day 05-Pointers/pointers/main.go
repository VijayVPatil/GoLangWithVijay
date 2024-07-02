package main

import "fmt"

func main() {

	//Declaring Pointers
	var ptr *int

	//Following will Print Zero value of the pointer
	fmt.Println(ptr)

	nos := 34
	//Below will store address of nos in nos1 variable using '&'
	nos1 := &nos

	//Using '*' we can dereference the pointer to get the value at that address
	value := *nos1
	fmt.Println("The Address is :", nos1)
	fmt.Println("The Value is at address stored in nos1 extracted using value which is:", value)

	var m = "Vijay"
	n := &m

	//Below should give the output as "The value Vijay ,the address <Address>, the type *string"
	fmt.Printf("The value %v ,the address %v ,the type %T \n", *n, n, n)

	//Pointer to a Pointer
	a := 32

	var b *int = &a

	var c **int = &b

	fmt.Println("Value of a is ", a, " and value of pointer to a is that is stored in b", b)

	fmt.Printf("To access the value at address stored in b using dereferencing , the value will be %v\n", *b)

	fmt.Println("*c will give me address of pointer of a that is &a that is b", *c)
	fmt.Println("**c will give ", **c)

}
