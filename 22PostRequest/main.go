package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id,omitempty"` // ID is returned by API
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	fmt.Println("CLASSKANI WAA POST -GU HOREEYEY EE GO LANG")
	newpost := post{UserID: 3, Title: "Manager", Body: "this is the body"}
	createpost(newpost)

	//CALL UPDATE POST FUNCTION
	updatedPost := post{UserID: 1, Title: "Updated Title", Body: "Update Body"}
	updatePost(1, updatedPost)
}

//CREATE POST

func createpost(post post) {

	data, _ := json.Marshal(post)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(data))

	// hadii ciladi dhacdo
	if err != nil {
		fmt.Println("cilada yaa dhacday fadlan baadh", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("created Post:")
	fmt.Println(body)
	deletePost(1)
}

//update Post

func updatePost(id int, post post) {

	data, _ := json.Marshal(post)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Cilad ayaa dhacday !", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Cilad ayaa dhacday !")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Updated Post ID %d:\n%s\n", id, string(body))
}

// DELETE POST

func deletePost(id int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		fmt.Println("Cilad ayaa dhacday iska jir", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("cilada yaa hadana dhacday", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Deleted Post ID %d. Status:%s\n", id, resp.Status)
}
