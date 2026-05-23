package main

import "fmt"

func main() {
	fmt.Println("STRUCTS IN GOLANG")

	// no inheritance in golang / no super or parent
	ahmed := User{"\nAHMED ABDILAHI\n", "bushaar2018@gmail.com\n", true, 20}
	fmt.Println(ahmed)
	fmt.Printf("ahmed details are : %+v\n", ahmed)
	fmt.Printf("Name is %v and Email is %v.", ahmed.Name, ahmed.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
