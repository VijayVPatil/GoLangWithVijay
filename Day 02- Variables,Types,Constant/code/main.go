package main

import "fmt"

func main() {

	//Below is how we declare a variable in go
	//We have declared the Variable username using var key word of type string
	var username string = "Vijay"
	fmt.Printf("Variable is of type: %T\n", username) // We use Printf

	//Below is how we declare muliple variables at once
	var b, c int = 1, 2
	fmt.Println("Multiple variables b and c is declared at once respectively: ", b, c)

	//Below is how we declare using short hand operator(:=)
	name1 := "Vijay"
	fmt.Println("The name is ", name1)

	//Boolean
	var isThisBoolean bool = true
	fmt.Printf("Variable is of type: %T\n", isThisBoolean)

	//Integer Types!
	//These can be type specific like uint8 , uint16, uint32, uint64
	//We usually go with int which is same as uint 32 or 64 bits
	//uint8
	var intNo uint8 = 125
	fmt.Printf("Variable is of type: %T\n", intNo)
	var intNoo int = 2342
	fmt.Printf("Variable is of type: %T\n", intNoo)

	//Float Types!
	//We have float32 and float64
	//Now float64 gives us a more precise result
	var smallFloat float32 = 234.2324224322235
	fmt.Printf("Variable is of type: %T and value is %v\n", smallFloat, smallFloat)

	var smallFloatt float64 = 234.2324224322235

	fmt.Printf("Variable is of type: %T and value is %v\n", smallFloatt, smallFloatt)

	//Implicit Way to declare the Variables
	//We need not to specidy the type as the go identifies it during the run time
	var name = "Vijay"
	fmt.Printf("The value is %v and type is %T", name, name)

	//We can also avoid using var and go identifies this but
	//need to go with :=   . This is called type inference
	//Type inference is only allowed inside main method and not outside
	noOfStudents := 32
	fmt.Printf("The value is %v and type is %T", noOfStudents, noOfStudents)

	//Constant are declared with const keyword and declared with captial letter starting as below
	//Constants can be declared outside main method but can not use type inference(:=)
	const Name string = "Vjay"
	fmt.Printf("The value is %v and type is %T", Name, Name)

}
