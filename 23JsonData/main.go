package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Pricce   int      // Typo: Should be Price
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("WELCOME JSON DATA IN GO LANG")
	jsondata()
	Decodejson()
}

func jsondata() {
	jsonda := []course{
		{"Go Lang", 300, "learningCode.com", "Abc123@@", []string{"web-dev", "js"}},
		{"Java Lang", 400, "learningCode.com", "Abc123@@", []string{"web-dev", "js"}},
		{"React Lang", 500, "learningCode.com", "Abc123@@", []string{"web-dev", "js"}},
		{"HTML Lang", 200, "learningCode.com", "Abc123@@", nil},
	}

	// package this data as json data
	finaljson, err := json.MarshalIndent(jsonda, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}

// CONSUME DATA

func Decodejson() {
	jsondatafrom := []byte(`
{
		"coursename": "React Lang",
        "Pricce": 500,
        "website": "learningCode.com",
        "tags": [
        "web-dev",
        "js"]
}
	`)
	var golan course
	checkvalid := json.Valid(jsondatafrom)

	if checkvalid {
		fmt.Println("json is valid")
		json.Unmarshal(jsondatafrom, &golan)
		fmt.Printf("%#v\n", golan)
	} else {
		fmt.Println("Json is not Valid")
	}

	// add some data in to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsondatafrom, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	//for loop

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T", k, v, v)
	}

}
