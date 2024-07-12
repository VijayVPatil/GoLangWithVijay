# Decoding(Unmarshalling) Json

## decoding also called unmarshalling in golang , is a way to convert/unpack a json object to struct, slice , map

1. We use a package "encoding/json" from the standard library and use a function unmarshal to convert the json object to data.
2. This can be used to convert the jsn objects to structs and primitive types

#### Code 1.01

```
package main

import (
	"encoding/json"
	"fmt"
)

type details struct {
	Username string `json:"coursename"`
	Id       int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Items    []string `json:"tags,omitempty"`
}

func main() {
	DecodeJSON()
}

func DecodeJSON() {
	jsonData := []byte(`
	{
		"coursename": "Vijay",
		"Id": 123,
		"website": "Flipkart",
		"tags": ["Book","Pencil"]
   }
	`)
	wrongJsonData := []byte(`
        {
            "coursename":"Jay",
            :
        }
   `)
	var detailstwo details

	checkValid := json.Valid(jsonData)

	//Below check is only done to show Valid function use to identify wrong json data
	wrongJsonDataCheck := json.Valid(wrongJsonData)

	fmt.Println("Wrong Json data check is ", wrongJsonDataCheck)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &detailstwo)
		fmt.Printf("%#v\n", detailstwo)
	} else {
		fmt.Println("Json is not value")
	}

	// Some Cases where you just want to add data to key value

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("The Key is %v and Value is %v\n", k, v)
	}
}


```

Explanation of Code 1.01

1. Here in main function we first create a dummy json data in format of slice of bytes
2. We have a Valid() function from encoding/json package that checks for valid json format. If the format is correct it return 'True' else 'False'
3. We first create a place where the unmarshlled json data will be stored. We create a struct of the same field format as json object.
4. json.Unmarshal() function will take the two valus , firts will be the json object that is variable named 'jsonData' in above example and pointer to the struct instance as unmarshal function will direclty modify the struct instance rather than its copy.

### More about the function func Unmarshal(data []byte, v interface{}) error

The Unmarshal function takes the slice of bytes and a empty interface. Here reason for empty inteerface is An empty interface may hold values of any type. (Every type implements at least zero methods.)
