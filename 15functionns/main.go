package main

import "fmt"

func main() {
	// fmt.Println("welcome G lang Functions")
	greater()
	anotherfun()

	result := add(12, 4)
	fmt.Println("the Result is:", result)

	//call sumall func
	sumallresul, mymessage := sumall(1, 2, 3)
	fmt.Println("Sum all of the values will be:", sumallresul)
	fmt.Println("kani wa messsagekayga labaaad ee functionkan", mymessage)

}

// functionka waxaad ka dhex call garaynaysaa main function-ka dhexdiisa
func greater() {
	fmt.Println("Hello From GO lang Function ")
}

// functionka waxaad ka dhex call garaynaysaa main function-ka dhexdiisa
func anotherfun() {
	fmt.Println("this is another example of GO lang Func")
}

// function added two values

func add(Valueone int, ValueTwo int) int {
	return Valueone * ValueTwo
}

// function qabanay inuu xisaabiyo dhaman values-ka oo dhan adoon garanayn valuesku intay noqon doonaan

func sumall(values ...int) (int, string) {
	total := 0

	// isticmaal for loop
	for _, val := range values {
		total += val
	}
	return total, "kani waa function kayga "
}
