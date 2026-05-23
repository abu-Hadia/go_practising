package main

import (
	"fmt"
	"strings"
)

func main() {

	Firstname := "Ahmed"
	Lastname := "Abdilahi"
	fullName := Firstname + " " + Lastname
	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}
