package main

import "fmt"

func main() {

	// for i := 1; i <= 40; i++ {
	// 	if i < 40 {
	// 		fmt.Println("soo good", i)
	// 	} else {
	// 		fmt.Println("stop iteration")
	// 	}
	// }

	// For LOOP WHICH PRINT *

	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//REVERSE FROM LARGE TO SMALL
	// for i := 10; i >= 1; i-- {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// loop for S

	n := 7

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			// top row
			if i == 0 ||
				// middle
				i == n/2 ||
				//upper right side
				(i < n/2 && j == 0) ||
				// lower right side
				(i > n/2 && j == n-1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
	}
}
