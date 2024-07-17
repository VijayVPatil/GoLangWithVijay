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

//Follwing method shows how to handle nil underlying types of interface
func (t *Teacher) testingNil() {
	if t == nil {
		fmt.Println("<nil> underlying type")
		return
	}
	fmt.Println(t.pages)
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

//Below show how we define a empty interface
type Empty interface {
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

	//Follwoing show how a interface value store value of different type if the type implements menthods that interface definition have
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

	//Following is how we declare a empty interface
	var emp Empty
	emp = st1
	emp = te1
	fmt.Println(emp)

	//Type assertions and type switch example
	typeSwitch(st1)
	typeSwitch(te1)
	typeSwitch(th)
}

func typeSwitch(i interface{}) {
	switch v := i.(type) { //i.(type) gives us the underlying type
	case Student:
		fmt.Printf("The type is of %v and %T\n", v, v)
	case Teacher:
		fmt.Printf("The type is of %v and %T\n", v, v)
	default:
		fmt.Printf("The type is of %v and %T\n", v, v)
	}
}
