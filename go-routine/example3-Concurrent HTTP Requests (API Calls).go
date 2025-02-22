// When making multiple API calls, running them sequentially slows down execution. Goroutines help speed this up.

// Fetching Data from Multiple APIs Concurrently

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when done

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", url)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response from %s: %s\n", url, body)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()
	fmt.Println("All API calls completed!")
}
