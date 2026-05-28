// package main

// import "fmt"

// func main() {
// 	// items := 6
// 	// priceItem := 10
// 	// if total := items * priceItem; total >= 60 {
// 	// 	fmt.Println("expensive", total)
// 	// } else if total >= 40 {
// 	// 	fmt.Println("its Ok")
// 	// } else {
// 	// 	fmt.Println("its cheap mansha allah")
// 	// }

// 	a := 10
// 	b := 4
// 	sum := a + b
// 	multi := a * b
// 	divide := a / b
// 	if sum > 0 {
// 		fmt.Println("sum is", sum)
// 	} else if multi < 40 {
// 		fmt.Println("multiple:", multi)
// 	} else if divide > 1 {
// 		fmt.Println("divide is:", divide)
// 	}

// }

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
