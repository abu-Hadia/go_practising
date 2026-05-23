package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	// user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter you Name please:")

	//comma || ERROR
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks ", input)
	fmt.Printf("the type of variable is %T", input)
}
