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
