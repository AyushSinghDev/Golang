package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //usually these are pointers
var mut sync.Mutex    //usually pointers

func main() {
	// go greet("Hello")
	// greet("World!")
	websitelist := []string{
		"https://go.dev",
		"https://chaiaurcode.in",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

/*
func greet(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(20 * time.Millisecond)
		fmt.Println(s)
	}
}
*/

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops, there is some issue in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
