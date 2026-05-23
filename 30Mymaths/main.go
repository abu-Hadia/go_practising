package main

import (
	"fmt"
	"math/big"

	//	"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("WELCOME TO MY GO LAND MATHS")

	// var myNumberOne int = 3
	// var myNumberTwo float64 = 5.5

	// // fmt.Println("Sum of Two numbers:", myNumberOne+int(myNumberTwo))

	// random numbers
	// rand.NewSource(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 2)

	// random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
