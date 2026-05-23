package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("SWICH  AND CASE FOR GO LANG")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(4) + 1

	fmt.Println("value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice values is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("you can move 2")
		fallthrough
	default:
		fmt.Println("what is that !")
	}

}
