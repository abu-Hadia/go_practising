package main

import (
	"fmt"
	"io"
	"net/http"
)

// define POST STRUCT
type POST struct {
	userID int    `json:"userid"`
	id     int    `json:"ID,omit empty"` //ID-ga uu soo celinayo IP-gu
	Title  string `json:"body"`
}

func main() {
	fmt.Println("WELCOME MY GO LANG FIRST GET REQUEST")

	//call all get  function-ka getPosts
	getPosts()
	// call GETPOST BY ID
	getPostByID(10)
}

//Function kan waaa GET function

func getPosts() error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("waxa dhacay error", err)
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("all Posts")
	fmt.Println(string(body))
	return nil
}

//Function kani waa GET POST BY ID

func getPostByID(id int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	resp, err := http.Get(url)

	//hadiiba uu wax error ahi dhaco
	if err != nil {
		fmt.Println("waxaa dhacay error!")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("GET BY ID %d:\n%s\n", id, string(body))

}
