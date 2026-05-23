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
	//floating

	float1 := 45.1
	float2 := 49.9
	avgfloat := (float1 + float2) / 2
	fmt.Println(avgfloat)

}
