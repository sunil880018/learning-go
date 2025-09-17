package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read body
		body, err := io.ReadAll(req.Body)
		data := json.RawMessage(body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer req.Body.Close()

		// Log the body to server console
		fmt.Println("Received Body:", data)

		// Send response back to client
		fmt.Fprintf(w, "You sent: %s", string(body))
	})

	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
