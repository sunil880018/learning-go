package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	p := Product{ID: 1, Name: "Laptop", Price: 999.99}

	// Marshal Go struct to JSON
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	data := string(jsonBytes)
	fmt.Println(data) // {"id":1,"name":"Laptop","price":999.99}

	err = json.Unmarshal([]byte(data), &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
}

// another way

// ------------- use this when --------------------
// When you dont know all the values are of the same type use interfac with map

// y := map[string]interface{}{
// 		"status":  200,
// 		"message": "OTP verified successfully",
// 		"data": map[string]interface{}{
// 			"user_id":       12,
// 			"session_token": "abcd1234efgh5678",
// 		},
// 	}
// 	fmt.Println(y)
