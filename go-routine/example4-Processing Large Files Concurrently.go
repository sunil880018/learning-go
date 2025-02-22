// Processing files in a single-threaded way is slow. Goroutines help by splitting the work.
// Example: Counting Words in a Large File in Parallel

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countWords(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wordCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordCount += len(strings.Fields(scanner.Text())) // Count words in a line
	}

	fmt.Printf("File: %s, Word Count: %d\n", filename, wordCount)
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go countWords(file, &wg)
	}

	wg.Wait()
	fmt.Println("All files processed!")
}

//  Goroutines are powerful for handling concurrency in Go.
//  They help speed up execution by running independent tasks in parallel.
