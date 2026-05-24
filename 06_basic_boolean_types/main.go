package main

import "fmt"

func main() {

	//boolean
	// isLogging := true
	// hasSubscription := true
	// isAdmin := true

	// cancancelpost := isAdmin && hasSubscription
	// canupdate := isLogging || isAdmin
	// fmt.Println("you can cancel:", cancancelpost)
	// fmt.Println("you can update it:", canupdate)
	var Age int = 29
	Adult := Age > 13
	if Adult {
		fmt.Println("adult is eligible to delete file:", Adult)
	} else {
		fmt.Println("you are not eligible to delete file")
	}

}
