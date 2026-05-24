package main

import "fmt"

func main() {

	grade := 50

	if grade > 50 && grade < 60 {
		fmt.Println("you are your grade is C")
	} else if grade >= 60 && grade < 75 {
		fmt.Println("your grade is B")
	} else if grade >= 75 && grade < 90 {
		fmt.Println("you grade is B+")
	} else if grade >= 90 {
		fmt.Println("you grade is A+")
	} else {
		fmt.Println("No grade ")
	}

}
