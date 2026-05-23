package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(("WELCOME TO MY FIRST CONVERSION"))
	fmt.Print("please enter the rate of pizza between 1 and 5 \n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks masuul sare", input)

	//converion

	numberReading, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add 1 to your rating :", numberReading+1)
	}

}
