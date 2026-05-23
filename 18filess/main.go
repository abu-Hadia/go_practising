package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Golang files")
	content := "this needs to go store files  dfdfkdfkdfjdskfdk"

	file, error := os.Create("./mygofile.txt")

	// waxa uu soo muujinaya cilada haysata hadiiba ay jirto meeshay ka jirto

	if error != nil {
		panic(error)
	}

	length, err := io.WriteString(file, content)

	// function errorka handle garaynaayey ayaa halkan lagu call garaynayaaa
	checkNilError(err)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

// function filekaad aad abuurtay soo akhriyay content-giisa

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	// function errorka handle garaynaayey ayaa halkan lagu call garaynayaaa
	checkNilError(err)

	fmt.Println("The text inside the file is :\n", string(databyte))

}

// commmon syntax of go error handling
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
