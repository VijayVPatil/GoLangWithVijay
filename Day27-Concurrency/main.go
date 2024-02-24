package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wq sync.WaitGroup

func main() {

	// go greeter("Hello")
	// greeter("World")
	endPoints := []string{
		"https://google.com",
		"https://github.com",
		"https://boot.dev",
	}

	for _, v := range endPoints {
		go checkStatus(v)
		wq.Add(1)
	}

	wq.Wait()

}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func checkStatus(endPoints string) {
	defer wq.Done()
	res, err := http.Get(endPoints)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("Status code of %s is %d\n", endPoints, res.StatusCode)
	}
}
