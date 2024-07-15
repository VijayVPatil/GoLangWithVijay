# Interfaces in Go

Interfaces by far has been super interesting in term on inital understanding and code walkthrough. I have tried to explain every aspect of code with a example I find easy.

## So what are interfaces?

1. Interfaces are a type that contains set of method signatures.
2. A value of a interface type can hold any value that implements the methods defined in Interface
   Following is how we define a interface

```
type variable_name interface{
    method1
    method2
    .
    .
    .
}

```

3. In simple words, when a type let's say a 'struct' provides all the definition(methods) for all the methods in an interface , then it is said that that type implements the interface. Heavy sentence? let's understand with a example

#### Example code 1.01

```

package main

import "fmt"

type Student struct {
	name  string
	pages int
}

func (s Student) reading() int {
	return s.pages
}

type Teacher struct {
	name  string
	pages int
}

func (t *Teacher) reading() int {
	return t.pages
}

type thermometer int

func (t thermometer) reading() int {
	if t < 0 {
		return int(-t)
	}
	return int(t)
}

type Result interface {
	reading() int
}

func main() {

	var report Result

	th := thermometer(24)
	st1 := Student{
		name:  "Rick",
		pages: 22,
	}

	te1 := Teacher{
		name:  "Morty",
		pages: 34,
	}

	report = th
    fmt.Println(report.reading())

	report = st1
    fmt.Println(report.reading())

	// report = te1 // This will throw an error

	report = &te1

	fmt.Println(report.reading())

}


```

Explanation of above (Code 1.01)"

1. As seen we have 2 struct type 'Student' , 'Teacher' and a non-struct type 'Thermometer'(which is just renamed as int) . Each type implement 'reading' method .
2. We have also defined a interface that have a method definition 'reading() int' .
3. Before jumping on how the interface will work lets go back to the points mentioned at the start(Point 2) and reference it with the example.
4. When a type, here it is 'Student'and 'Teacher' a struct type and a 'thermomter' a non struct(int) type provides all the method definitions here 'reading()' for all the methods in an interface(interface also have reading method defined) then it is said that 'Student', 'Teacher' and 'thermomter' implements the interface.
5. If you closely observe the reading method implemented by Student struct and Thermometer have value receiver and reading method implemented by Teacher struct have pointer receiver. This distinction matters for implementing interfaces and we will see why just ahead.
6. Now in the main function variable 'report' of 'Result' can hold any value that satisfies the Result interface.
7. In code the 'th' and 'st1' variable is assigned to 'report' throwing no error demonstrating that both types satisfy the 'Result' interface.
8. If the reading method for Teacher was defined with a pointer receiver, then only pointers to Teacher would satisfy the Result interface, not the Teacher values themselves. Therefore, the assignment reprot = te1 would not work ,instead you would need to assign a pointer to te1.
