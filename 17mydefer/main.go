package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello World defer")
	defer fmt.Println("two defer")
	fmt.Println("no defer")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
}
