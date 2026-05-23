package main

import (
	"fmt"
	"io/ioutil" // In modern Go, prefer "io" package's io.ReadAll
	"net/http"
)

func main() {
	fmt.Println("Its my first line of http request in GO lang")

	// --- FIX START ---
	// Declare and initialize the 'url' variable with the URL you want to fetch
	const url = "https://www.google.com" // Or any other URL you want to request
	// --- FIX END ---

	response, err := http.Get(url) // Now 'url' is defined

	if err != nil {
		// It's generally better to print the error and return, rather than panic,
		// unless it's an unrecoverable application-level error.
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return // Exit the function gracefully
	}

	// Corrected Printf format string for type
	// %T prints the type of the value
	fmt.Printf("Response is type of %T\n", response)

	// Always defer closing the response body to prevent resource leaks
	defer response.Body.Close()

	// Use io.ReadAll for reading response bodies in modern Go
	// (ioutil.ReadAll is deprecated but still works for now)
	databye, err := ioutil.ReadAll(response.Body)

	if err != nil {
		// Again, better to print and return than panic for I/O errors
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	content := string(databye) // Convert the byte slice to a string
	fmt.Println(content)       // Print the content of the response body
}
