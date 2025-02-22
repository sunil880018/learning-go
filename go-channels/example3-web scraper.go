package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Function to fetch data from a URL
func fetchURL(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as done

	resp, err := http.Get(url) // Fetch data
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)                     // Read response body
	ch <- fmt.Sprintf("URL: %s, Length: %d", url, len(body)) // Send data to channel
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
	}

	ch := make(chan string, len(urls)) // Buffered channel to store results
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)                 // Increment counter
		go fetchURL(url, ch, &wg) // Start goroutine
	}

	wg.Wait() // Wait for all goroutines to finish
	close(ch) // Close channel after all work is done

	// Read results from channel
	for result := range ch {
		fmt.Println(result)
	}
}
