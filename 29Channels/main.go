package main

import (
	"fmt"
	"sync"
)

//GO CHANNELS

func main() {
	fmt.Println("Channels is goland")
	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()

	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// myCh <- 5
		// myCh <- 10
		myCh <- 0
		close(myCh)

		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
