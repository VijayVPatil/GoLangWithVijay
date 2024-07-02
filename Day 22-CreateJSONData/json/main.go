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
	EncodeJSON()
}

func EncodeJSON() {
	detailsOne := []details{
		{"Vijay", 123, "Flipkart", "abc123", []string{"Book", "Pencil"}},
		{"Varun", 223, "Amazon", "bcd123", []string{"Yogurt", "Pen"}},
		{"Josh", 263, "Myntra", "ssc123", nil},
	}

	//Packaging this data as JSON data

	//finalJson, err := json.Marshal(detailsOne) //Using Marshal we find it hard to read the output(not in formated output) so just ike marshal we have marshal indent
	finalJson, err := json.MarshalIndent(detailsOne, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
