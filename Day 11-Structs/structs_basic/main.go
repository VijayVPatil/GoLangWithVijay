package main

import "fmt"

//No Inheritance in Golang, No Super or Parent or Child

//Defining Structs
// AS it should be accessible outside, We give first letter as Capital
type Employee struct {
	empId   int
	empName string
}

func main() {
	//Structs
	//Declaring Structs
	var emp1 Employee

	fmt.Println("The zero value of above struct instance", emp1)

	//Initalizing Structs
	emp1 = Employee{
		empId:   3444,
		empName: "Vijay",
	}

	fmt.Printf("The value of struct instance is %+v and type is %T\n", emp1, emp1) //+v gives more information

	// Accessing fields of struct using '.' operator
	fmt.Printf("Empoyee Id is : %v ,Employee Name is %v\n", emp1.empId, emp1.empName)

	//Pointer to structs
	var varAddress *Employee
	varAddress = &emp1
	fmt.Println(varAddress.empId, varAddress.empName)

	//Pointer to a struct using short hand operator
	empAdd := &emp1

	///Accessing the fields to emp1 using empAdd by dereferencing
	fmt.Println("The emp1 id using dereferencing operator is ", (*empAdd).empId, "and emp1 name using dereferencing operator is ", (*empAdd).empName)

	///We can also directly access Employee fields without dereferncing
	fmt.Println("The emp1 id is ", empAdd.empId, "and emp1 name is ", empAdd.empName)

	fmt.Printf("The type of pointer to the structure is %T\n", empAdd)

}
