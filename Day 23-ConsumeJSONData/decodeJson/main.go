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
