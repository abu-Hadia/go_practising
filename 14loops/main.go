package main

import "fmt"

func main() {
	fmt.Println("welcome Go LOOPS")

	days := []string{"Saterday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//ANOTHER EXAMPLE OF LOOP

	// for index, days := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, days)
	// }

	//another example

	// for _, days := range days {
	// 	fmt.Printf("index is and value is %v\n", days)
	// }

	// ANOTHER EXAMPLE loop inside if

	number := 0
	for number < 100 {

		if number == 2 {
			goto welcome
		}

		if number == 2 {
			number--
			continue
		}

		fmt.Println("the value is :", number)
		number++
	}

welcome:
	fmt.Println("welcome again masuul")

}
