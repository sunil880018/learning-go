package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Home Page - Handles GET and POST
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Welcome to the Home Page! (GET Request)")
	} else if r.Method == http.MethodPost {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "Received POST Request with body: %s", string(body))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// About Page - Only Handles GET
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "This is the About Page.")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server is running on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
