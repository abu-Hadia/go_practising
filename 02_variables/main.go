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
	Name := "Ahmed"
	age := 40.4
	fmt.Println("Name:", Name)
	fmt.Println("Age", age)

	//Grouped variables
	// var (
	// 	JName  string = "jama"
	// 	Salary int    = 4000
	// 	exper  int    = 5
	// )
	// fmt.Println("JName:", JName)
	// fmt.Println("Salary:", Salary)
	// fmt.Println("Exper:", exper)

	///ANOTHER TYPE YOU CAN DECLARE VAR

	subscriber := 400 //----infer  it detects automatically the data type of the variable.
	subscriber = subscriber + 1000

	product, service := "sugar", 40

	fmt.Println(subscriber, product, service)

}
