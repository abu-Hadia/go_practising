package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to GO Slices")

	fruList := []string{"\napple\n", "banana\n", "peach\n"}
	fmt.Printf("the type of frulist is %T\n", fruList)

	//add another list to that slices
	fruList = append(fruList, "Tufaax\n", "Cinab\n")
	fmt.Println(fruList)

	// start slices from 1

	fmt.Println("first Array/Slices is missing\n")
	fruList = append(fruList[1:])
	fmt.Println(fruList)

	// Start and End
	fmt.Println("print only allowed slices\n")
	fruList = append(fruList[1:3])
	fmt.Println(fruList)

	// make another example of slices

	highscore := make([]int, 3)
	highscore[0] = 432
	highscore[1] = 452
	highscore[2] = 498
	//highscore[3] = 476
	fmt.Println(highscore)

	// add nod defined values to the slices
	highscore = append(highscore, 123, 542)
	fmt.Println(highscore)

	//sort slices
	sort.Ints(highscore)
	fmt.Println(highscore)

	//check if sorted
	fmt.Println(sort.IntsAreSorted(highscore))

	// how to remove value from slices bases on index
	// courses:= waxa la mid tahay uun var courses=[]string

	courses := []string{"reacts", "javascript", "java", "go"}
	fmt.Println(courses)

	var index int = 3

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
