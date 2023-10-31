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
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Vijay",
		"Id": 123,
		"website": "Flipkart",
		"tags": ["Book","Pencil"]
   }
	`)
	var detailstwo details

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &detailstwo)
		fmt.Printf("%#v\n", detailstwo)
	} else {
		fmt.Println("Json is not value")
	}
}
