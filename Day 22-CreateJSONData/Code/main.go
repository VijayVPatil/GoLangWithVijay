package main

import (
	"encoding/json"
	"fmt"
)

//Please check the step by step explanation in Day 22-CreateJSONData/encoding.md

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,-"`
	Age   int    `json:age`
	Class *Class `json:",omitempty"`
}

type Class struct {
	NoOfStudents  int    `json:"noOfStudents"`
	NameOfTeacher string `json:"nameOfTeacher"`
}

func main() {

	s1 := Student{

		Name: "Vijay",
		Age:  12,
	}

	st, error := json.Marshal(s1)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(st))
}
