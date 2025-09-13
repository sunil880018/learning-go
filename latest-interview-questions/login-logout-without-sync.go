package main

import (
	"fmt"
	"time"
)

var users = map[string]string{
	"alice": "1234",
	"bob":   "qwerty",
}

type LoginRequest struct {
	User string
	Pass string
	Resp chan string
}

func loginService(reqs <-chan LoginRequest) {
	for req := range reqs {
		if pass, ok := users[req.User]; ok && pass == req.Pass {
			req.Resp <- "Login success: " + req.User
		} else {
			req.Resp <- "Login failed: " + req.User
		}
	}
}

func main() {
	reqChan := make(chan LoginRequest)
	go loginService(reqChan)

	tryLogin := func(u, p string) {
		respCh := make(chan string)
		reqChan <- LoginRequest{User: u, Pass: p, Resp: respCh}
		fmt.Println(<-respCh)
	}

	go tryLogin("alice", "1234")
	go tryLogin("bob", "wrong")
	go tryLogin("charlie", "nope")

	time.Sleep(500 * time.Millisecond)
}
