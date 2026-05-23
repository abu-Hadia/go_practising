package main

import "fmt"

func main() {
	fmt.Println("welcom to pointer")

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber1 := 22

	var ptr = &myNumber1
	fmt.Println("value of actual pointer is", ptr)
	fmt.Println("value of actual of pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("the new value is ", myNumber1)

}
