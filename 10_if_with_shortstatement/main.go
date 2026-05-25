package main

import "fmt"

func main() {
	// items := 6
	// priceItem := 10
	// if total := items * priceItem; total >= 60 {
	// 	fmt.Println("expensive", total)
	// } else if total >= 40 {
	// 	fmt.Println("its Ok")
	// } else {
	// 	fmt.Println("its cheap mansha allah")
	// }

	a := 10
	b := 4
	sum := a + b
	multi := a * b
	divide := a / b
	if sum > 0 {
		fmt.Println("sum is", sum)
	} else if multi < 40 {
		fmt.Println("multiple:", multi)
	} else if divide > 1 {
		fmt.Println("divide is:", divide)
	}

}
