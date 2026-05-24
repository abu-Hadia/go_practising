package main

import "fmt"

func main() {
	// var Name string
	// var Age int
	// fmt.Println("Name:", Name)
	// fmt.Println("Age:", Age)
	// var cash float64
	// fmt.Println("Cash:", cash)

	/// create a variable and assign a value at the same time
	// Name := "Ahmed"
	// age := 40.4
	// fmt.Println("Name:", Name)
	// fmt.Println("Age", age)

	// Grouped variables
	// var (
	// 	JName  string = "jama"
	// 	Salary int    = 4000
	// 	exper  int    = 5
	// )
	// fmt.Println("JName:", JName)
	// fmt.Println("Salary:", Salary)
	// fmt.Println("Exper:", exper)

	///ANOTHER TYPE YOU CAN DECLARE VAR

	// subscriber := 400 //----infer  it detects automatically the data type of the variable.
	// subscriber = subscriber + 1000

	// product, service := "sugar", 40

	// fmt.Println(subscriber, product, service)
	// numb2 := 300
	// number3 := 40
	// fmt.Println("Total is:", numb2+number3)

	//small calculation

	num1 := 30
	num2 := 40
	total := num1 * num2
	fmt.Println("Total is:", total)
	if total > 1000 {
		fmt.Println("Total is greater than 1000:", total)
	} else {
		fmt.Println("Total is less than 1000 so division Result is:", float64(num1)/float64(num2))
	}

}
