package main

import (
	"fmt"
)

// func main() {

// 	var num1, num2 float64
// 	var operator string

// 	//input
// 	fmt.Println("Enter first Number:")
// 	fmt.Scan(&num1)
// 	fmt.Println("Enter the operator")
// 	fmt.Scan(&operator)
// 	fmt.Println("Enter the second Number:")
// 	fmt.Scan(&num2)

// 	//process
// 	// if operator == "+" {
// 	// 	fmt.Println("Addition is:", num1+num2)
// 	// } else if operator == "-" {
// 	// 	fmt.Println("substraction is:", num1-num2)
// 	// } else if operator == "*" {
// 	// 	fmt.Println("multiplication is:", num1*num2)
// 	// } else if operator == "/" {
// 	// 	fmt.Println("division is:", num1/num2)
// 	// } else {
// 	// 	fmt.Println("invalid format")
// 	// }

// 	switch operator {
// 	case "+":
// 		fmt.Println(num1, "+", num2, "=", num1+num2)

// 	case "-":
// 		fmt.Println(num1, "-", num2, "=", num1-num2)
// 	case "*":
// 		fmt.Println(num1, "*", num2, "=", num1*num2)
// 	case "/":
// 		if num2 == 0 {
// 			fmt.Println("division by zero is not allowed")
// 		} else {
// 			fmt.Println(num1, "/", num2, "=", num1/num2)
// 		}
// 	default:
// 		fmt.Println("invalid operator")
// 	}

// }

func main() {

	for {
		var num1, num2 float64
		var operator string

		//input
		fmt.Println("Enter first Number:")
		fmt.Scan(&num1)
		fmt.Println("Enter the operator")
		fmt.Scan(&operator)
		fmt.Println("Enter the second Number:")
		fmt.Scan(&num2)

		//process
		// if operator == "+" {
		// 	fmt.Println("Addition is:", num1+num2)
		// } else if operator == "-" {
		// 	fmt.Println("substraction is:", num1-num2)
		// } else if operator == "*" {
		// 	fmt.Println("multiplication is:", num1*num2)
		// } else if operator == "/" {
		// 	fmt.Println("division is:", num1/num2)
		// } else {
		// 	fmt.Println("invalid format")
		// }

		switch operator {
		case "+":
			fmt.Println(num1, "+", num2, "=", num1+num2)

		case "-":
			fmt.Println(num1, "-", num2, "=", num1-num2)
		case "*":
			fmt.Println(num1, "*", num2, "=", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("division by zero is not allowed")
			} else {
				fmt.Println(num1, "/", num2, "=", num1/num2)
			}
		default:
			fmt.Println("invalid operator")
		}
		fmt.Println()
	}

}
