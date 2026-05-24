package main

import "fmt"

func main() {
	items := 6
	priceItem := 10
	if total := items * priceItem; total >= 60 {
		fmt.Println("expensive", total)
	} else if total >= 40 {
		fmt.Println("its Ok")
	} else {
		fmt.Println("its cheap mansha allah")
	}

}
