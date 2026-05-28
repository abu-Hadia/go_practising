package main

import "fmt"

var myvar int = 200

func main() {

	fmt.Println(myvar)
	myfunnc()

	// multiple variable declarateion
	var year, age, name = 2026, 30, "Ahmed"
	fmt.Println(year, name, age)
}

func myfunnc() {
	fmt.Println("myvariable", myvar)
}
