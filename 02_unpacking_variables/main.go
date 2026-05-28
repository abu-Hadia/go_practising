package main

import "fmt"

var myvar int = 200

func main() {

	fmt.Println(myvar)
	myfunnc()
}

func myfunnc() {
	fmt.Println("myvariable", myvar)
}
