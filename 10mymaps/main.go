package main

import "fmt"

func main() {

	fmt.Println("GO MAPS")
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RP"] = "ruby"
	language["PY"] = "python"
	fmt.Println("List of all languages", language)
	fmt.Println("RP shorts for:", language["RP"])

	// delete value from map
	delete(language, "RP")
	fmt.Println("One language is delete ", language)

	// LOOPS
	for _, value := range language {
		fmt.Printf("For key v, value is %v\n", value)
	}

}
