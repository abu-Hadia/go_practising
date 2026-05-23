package main

import "fmt"

func main() {
	views1 := 100
	views2 := 300
	totalviews := views1 + views2
	like := 3000
	like++
	like++
	avgviews := totalviews / 2
	fmt.Println(totalviews, like, avgviews)

}
