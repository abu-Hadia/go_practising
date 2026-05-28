package main

import "fmt"

const v = 199
const x float64 = 333.22
const y string = "Ahmed"

func main() {

	// there are two types of constants
	// and they are
	// 1. typed
	// 	Constants that have an explicitly declared type.
	// const age int = 25
	// const price float64 = 99.99
	// const name string = "Go"

	// 2. untyped
	// 	2. Untyped constants (most important in Go)

	// Constants declared without a type.

	// const age = 25
	// const price = 99.99
	// const name = "Go"

	// var a int = v

	// var v float64 = x
	// var c string = y
	// fmt.Println(c)
	// fmt.Println(v)
	// fmt.Println(a)
	const (
		firstname string = "ahmed"
		lastname  string = "Abdilahi"
		fullname  string = firstname + " " + lastname
	)
	fmt.Println(fullname)

}
