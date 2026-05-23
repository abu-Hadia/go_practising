package main

import "fmt"

//public variable
const LoginToken string = "ahmed abdilahi"

func main() {
	// DECLARING GO VARIABLES
	var username string = "Ahmed"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	// GO BOOLEAN
	var islogin bool = true
	fmt.Println(islogin)
	fmt.Printf("variable is of type: %T \n", islogin)

	// GO SMALL VALUE

	var smallval uint16 = 256
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T \n", smallval)

	var smallfloat float64 = 256.333544333333
	fmt.Println(smallfloat)
	fmt.Printf("variable is off type : %T \n", smallfloat)

	// default values and some alias
	var anothervariables int
	fmt.Println(anothervariables)
	fmt.Printf("variable is type of : %T\n", anothervariables)

	//implicit type
	var website = "welcome to my website"
	fmt.Println(website)

	// no variable style
	numerofusers := 30000
	fmt.Println(numerofusers)

	//calling public variable inside a function
	fmt.Println(LoginToken)
	fmt.Printf("this is public variabl and type is : %T\n", LoginToken)
}
