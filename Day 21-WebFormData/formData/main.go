package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	performPostFormJsonRequest()
}

func performPostFormJsonRequest() {
	const MYURL = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "vijay")
	data.Add("lastname", "patil")
	data.Add("email", "vj@slack")

	response, err := http.PostForm(MYURL, data)

	if err != nil {
		panic(response)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(response)
	}

	fmt.Println(string(content))
}
