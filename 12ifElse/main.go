package main

import "fmt"

func main() {
	fmt.Println("Welcome to IF else Statement for GO LANG")

	loginCount := 4
	var result string
	if loginCount <= 3 {
		result = "Good User"
	} else if loginCount >= 3 {
		result = "You are Out of loginout"
	} else {
		result = "Good bye"
	}

	fmt.Println(result)

	// more example about ifelse stat for go lang
	// % waxa loo yaqaanaa remainder laba number marka laysu qaybiyo maxaa baaqi noqda
	if 9%777 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	/// another exammple

	if num := 1; num < 5 {
		fmt.Println("number is less then 5")
	} else {
		fmt.Println("num is greater then 1")
	}
}
