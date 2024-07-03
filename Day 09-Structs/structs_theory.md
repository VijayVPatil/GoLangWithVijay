# STRUCT in Golang

1. Struct in golang is a composition of different simpler types like int, float64/float32 , string etc .
2. It is a user defined type.

#### Declaration

Structs are declared as below

```
type <identifier> struct{
    <variable name> type
    .
    .
    .
}
```

ex:

```
type Employee struct{
    empId int
    empName string
}
```

#### How to define a struct

We can define struct as below:

```

package main

import "fmt"

type Employee struct {
	empId   int
	empName string
}

func main(){

var emp1 Employee

fmt.Println(emp1)
}
```

Output of above will be :

```
{0 }
```

When the struct is not initalised with any values , the zero value of each type used in struct is set.

#### Struct initalization

```

package main

import "fmt"

type Employee struct {
	empId   int
	empName string
}

func main(){

//Defining a Struct
var emp1 Employee

//Initalizing a struct
emp1 = Employee{
		empId:   3444,
		empName: "Vijay",
	}
fmt.Println(emp1)
}

fmt.Printf("The value of struct instance is %+v and type is %T\n", emp1, emp1) //+v gives more information

//Accessing Employee Struct field using instance
fmt.Printf("Empoyee Id is : %v ,Employee Name is %v\n", emp1.empId, emp1.empName)

```

We can also access the fields of struct using dot '.' operator, Check above example where we have accessed Employee Id and name. The type of the above 'emp1' instance is main.Employee that is the structure.

#### Pointer to a Struct

Pointers as we know can store memmory address of another variable.

###### Please checkout the pointers section if new to Pointers

We can also store the address of the structure in a variable

```
    empAdd := &emp1
	fmt.Println("The emp1 id is ", (*empAdd).empId, "and emp1 name is ", (*empAdd).empName)

	fmt.Println("The emp1 id is ", empAdd.empId, "and emp1 name is ", empAdd.empName)

    fmt.Printf(""The type of pointer to the structure is %T\n",empAdd)
```

The address
