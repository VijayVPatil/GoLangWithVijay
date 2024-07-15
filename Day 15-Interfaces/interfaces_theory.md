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

// Function taking interface a argument value
func workFunction(r Result) {

	fmt.Printf("The reading is %v \n", r.reading())
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

	//report = te1

	report = &te1

	fmt.Println(report.reading())

	workFunction(th)
	workFunction(st1)
	workFunction(&te1)

	st2 := Student{
		name:  "Ken",
		pages: 3,
	}

	report = st2

	//Calling a method on an interface value executes the method of the same name on its underlying type.
	report.reading()
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

### Points to Note

1. You don't need to declare that a type implements an interface; it happens automatically if the required methods are present.
2. Any type, in any package, can implement the interface as long as it provides the required methods.
3. There is no need for the implementing types to declare that they satisfy the interface explicitly.

### Interfaces under the hood!

1. When we call a method on an interface value, Go uses the type information to let the call to the appropriate method on the underlying type. In code Example code 1.01 ,internally report hold the type information of st2(struct type). When we call report.reading(), Go checks the type information stored in report to find out the underlying type, which is st2. Go then calls the reading() method on the value st2.
