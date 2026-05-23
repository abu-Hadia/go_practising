package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("WELCOME TO MY TIME STUDY FOR GO LANG")
	timepresent := time.Now()
	fmt.Println(timepresent)
	fmt.Println(timepresent.Format("02-01-2006 09:00:00 wednesday"))

	//ceate date
	createDate := time.Date(2024, time.January, 10, 23, 34, 0, 2, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
