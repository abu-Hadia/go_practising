package main

import "fmt"

func main() {

	//boolean
	isLogin := true
	isadmin := false
	hassubscription := true

	canDelete := isLogin && isadmin
	canPost := isadmin || hassubscription

	fmt.Println(canDelete, canPost)

}
