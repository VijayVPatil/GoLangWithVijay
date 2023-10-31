package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	performPostJsonRequest()
}

func performPostJsonRequest() {
	const MYURL = "http://localhost:8000/post"

	//Fake JSON Payload

	requestBody := strings.NewReader(`
	{
		"Name":"Vijay",
		"Number":"01",
		"Twitter Handel":"@vijayOnTwt"

	}
	
	`)

	response, err := http.Post(MYURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	result, _ := io.ReadAll(response.Body)

	fmt.Println(string(result))
}
