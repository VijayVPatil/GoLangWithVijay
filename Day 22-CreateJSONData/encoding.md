# Encoding in Golang

## Encoding also called marshalling in golang , is a way to convert a data structue into json object

1. We use a package "encoding/json" from the standard library and use a function marshal to convert the data to json object.
2. This can be used to convert the structs into serializable json object.
3. Lets refer the code snippet in CreateJSONData\json\main.go

###### Code(1.01)

```
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    int
	Name  string
	Class Class
}

type Class struct {
	NoOfStudents  int
	NameOfTeacher string
}

func main() {

	s1 := Student{
		Id:   23,
		Name: "Vijay",
		Class: Class{
			NoOfStudents:  32,
			NameOfTeacher: "Mr. John",
		},
	}

    //st, error := json.Marshal(s1)
    st,error := json.MarshalIndent(s1, "", "\t")
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(st))
}

```

#### Explanation of Code(1.01)

1. Marshal() function converts the struct to slice of bytes. We also have MarshalIndent function that takes a any type value with prefix and indent to format the output
2. It takes a struct or pointer to strut as a argument and return a slice of bytes
3. The slice of bytes can then be converted to string as seen above.
4. They are converted using UTF-8 encoding

## Omitempty property

In scenarios where our struct fields don't have any values , we do not want them to pass as there zero values rather we won't pass them

##### Code(1.02)

Please check explanation

```
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,-"`
	Class Class
}

type Class struct {
	NoOfStudents  int    `json:"noOfStudents"`
	NameOfTeacher string `json:"nameOfTeacher"`
}

func main() {

	s1 := Student{

		Name: "Vijay",
		Class: Class{
			NoOfStudents:  32,
			NameOfTeacher: "Mr. John",
		},
	}

	st, error := json.Marshal(s1)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(st))
}

```

#### Explanation of Code(1.02)

1. Lets say in some scenario the s1 student don't have a ID yet , so struct initalization don't include the ID field.
2. When we add a omitempty property for the Student struct , now if ID value set to it’s default value (which is 0 for the int type), the key itself will be omitted from the JSON object. Try running the program without omitempty tag in json , we see the fields with zero value
3. The same will apply if a string is empty "", or if a pointer is nil, or if a slice has zero elements in it.

### Are you running above programs if no please do. Lets go!!!!!!!!!!!!!!!!!!!!

4. What if we want to completly omit a field,then we use the tag '-' .Just like in Student structs with json tag we have attached '-'

5. We have a limitation when using omit empty tag , let's say we use omit empty tag with the embedded struct class that has fleds with zero values. It wll still show in json object.
6. To solve this we use struct pointers , struct pointers have empty value as nil. Try with below

```
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,-"`
    Age int `json:age`
	Class *Class `json:",omitempty"`
}

type Class struct {
	NoOfStudents  int    `json:"noOfStudents"`
	NameOfTeacher string `json:"nameOfTeacher"`
}

func main() {

	s1 := Student{

		Name: "Vijay",
        Age: 12,

	}

	st, error := json.Marshal(s1)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(st))
}

```
