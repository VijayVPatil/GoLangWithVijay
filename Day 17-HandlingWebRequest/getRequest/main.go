package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	// Web Request

	//Make a http request
	response, err := http.Get(url)

	//If any error like internet connection , 400 error then follwoing is executed
	if err != nil {
		panic(err)
	}

	fmt.Printf("The Type of response is %T\n", response)
	// fmt.Printf("The value of response is %v\n", response)

	/* WHY RESPONSE BODY NEEDS TO BE CLOSED
	The response body should be closed after the response is fully read.
	This is done to prevent resource leak of connections.
	If the response body is not closed then the connection will not be released and hence it cannot be reused */

	defer response.Body.Close() // resp. Body is an object that has a some data inside of it which we can read out and has a function that allows us to close the stream.
	//We always close it at end of the main function execution using Defer keyword

	/* Reading the Response */

	result, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result)         //The result will return in byte code
	fmt.Println(string(result)) //The result that is byte code is converted into string
}
