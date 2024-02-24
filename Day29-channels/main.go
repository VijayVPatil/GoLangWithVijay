package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")
	//Channel is a way or a pipeline through which multiple go routines interact

	//To create a channel
	mych := make(chan int, 2) //This tells that we are making a channel with buffer which is used to interact with multiple go routines and we pass a integer between them
	wg := &sync.WaitGroup{}

	//mych <- 5 // Pushing a 5 value in channel
	//fmt.Println(<-mych)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych) //closing my channel
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
