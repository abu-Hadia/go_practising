package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"Test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex    // pointer

func main() {

	// go waxaaad u sheegaysaa shaqooyinkan ka horaynaya waan concept-ga la yidhaa concurrency
	// greater("WELCOME GO ROUTINES")
	// go greater("WELCOME GO CONCURRENCY")

	websiteList := []string{
		"https://lco.dev",
		"https://google.com",
	}

	for _, web := range websiteList {
		go getstatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greater(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getstatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS Is endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
