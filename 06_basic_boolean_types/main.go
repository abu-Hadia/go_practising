package main

import "fmt"

func main() {

	//boolean
	isLogging := true
	hasSubscription := true
	isAdmin := true

	cancancelpost := isAdmin && hasSubscription
	canupdate := isLogging || isAdmin
	fmt.Println("you can cancel:", cancancelpost)
	fmt.Println("you can update it:", canupdate)

}
