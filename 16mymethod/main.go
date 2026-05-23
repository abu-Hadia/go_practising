package main

import "fmt"

func main() {
	fmt.Println("STRUCTS IN GOLANG")

	// no inheritance in golang / no super or parent
	ahmed := User{"\nAHMED ABDILAHI\n", "bushaar2018@gmail.com\n", false, 20}
	fmt.Println(ahmed)
	fmt.Printf("ahmed details are : %+v\n", ahmed)
	fmt.Printf("Name is %v and Email is %v.", ahmed.Name, ahmed.Email)
	ahmed.GetStatus()
	ahmed.NewEmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// METHOD

func (u User) GetStatus() {
	fmt.Println("is user status is active", u.Status)
}

// func or method

func (u User) NewEmail() {
	u.Email = "ayuusuf@telesom.com"
	fmt.Println("My New Email is:", u.Email)
}
