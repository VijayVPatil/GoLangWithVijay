package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	//Web Request

	PerformGetRequest()
}

func PerformGetRequest() {

	const MYURL = "http://localhost:8000/get"

	response, err := http.Get(MYURL)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	result, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result)

	//Converting Byte to String Way 1
	fmt.Println(string(result))

	//Converting Byte to String Way 2
	var responseString strings.Builder
	byteCount, _ := responseString.Write(result)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
