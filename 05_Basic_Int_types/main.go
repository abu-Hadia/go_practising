package main

import (
	"fmt"
)

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

	var smallFload float64
	fmt.Println(smallFload)
	smallFload = 55.333
	fmt.Println(smallFload)

	var largeFloat float64
	fmt.Println(largeFloat)
	largeFloat = 999999.1
	fmt.Println(largeFloat)

	//complex numbers

	var mycomplex complex128
	mycomplex = complex(smallFload, largeFloat)
	fmt.Println(mycomplex)

	/// other data types

	var positivevalues uint8
	positivevalues = 100
	fmt.Println(positivevalues)

	/// negative data type

	var negativevalues int8
	negativevalues = -100
	fmt.Println(negativevalues)

	// typecase
	var mybye byte
	mybye = 'd'
	fmt.Println(mybye)

	// rune is alias of int8

	// Signed integers (can be negative and positive)
	// Type	Size	Range
	// int8	8-bit	-128 to 127
	// int16	16-bit	-32,768 to 32,767
	// int32	32-bit	~ -2B to 2B
	// int64	64-bit	very large range

	// 	Unsigned integers (only positive, no negatives)
	// Type	Size	Range
	// uint8	8-bit	0 to 255
	// uint16	16-bit	0 to 65,535
	// uint32	32-bit	0 to ~4B
	// uint64	64-bit	very large
	// 	int
	// Size depends on your system:
	// 32-bit system → same as int32
	// 64-bit system → same as int64
	// 👉 Most commonly used integer
	// uint
	// Same as int, but only positive
	// 👉 Avoid unless you really need unsigned behavior
	// 🔤 rune
	// Alias for int32
	// Used to represent Unicode characters

}
