package main

import "fmt"

func main() {
	fmt.Println("KUSOO DHAWOOW KHUDARALAHAYGA")
	// array
	var fruitList [3]string
	fruitList[0] = "\nTufaax \n"
	fruitList[1] = "muus \n"
	fruitList[2] = "liinmacaan \n"

	fmt.Println("Imisa Khudaara ayaad haysaa :", fruitList)
	fmt.Println("the bigest array is", len(fruitList))

	// another example
	var vegtabList = [4]string{"xabxab", "liinmacaan", "kaabash", "cinab"}
	fmt.Println("vegtable list is: ", len(vegtabList))

}
