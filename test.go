package main

import (
	"encoding/json"
	"log"
	"os"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	todo := Todo{
		UserID:    1,
		ID:        101,
		Title:     "Learn Go",
		Completed: false,
	}

	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(todo)
	if err != nil {
		log.Panic(err)
	}
}
