package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

func main() {
	fileName := "./text.txt"
	content, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content:", content)
}
