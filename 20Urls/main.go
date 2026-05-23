package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://myurl.go:3000/learnGo?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println("how to handle url in GO lang")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)   // waxa soo qabanayaa https
	// fmt.Println(result.Host)     // waxa uu soo qabanayaa hostname
	// fmt.Println(result.Port())   // waxa uu soo qabanayaa port-ga
	fmt.Println(result.RawQuery) // waxa uu soo qabanayaa
	qprograms := result.Query()
	fmt.Printf("the type of query params are:", qprograms)
	fmt.Println(qprograms["coursename"])

	for _, val := range qprograms {
		fmt.Println("program is:", val)
	}

	partsOfurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=ahmed",
	}
	anotherurl := partsOfurl.String()
	fmt.Println(anotherurl)

}
