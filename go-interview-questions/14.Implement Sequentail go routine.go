package main

import (
	"fmt"
)

type Notification struct{}

func (n Notification) SendEmailNotification(i int, done chan bool) {
	fmt.Println("Email", i)
	done <- true
}

func (n Notification) SendSMSNotification(i int, done chan bool) {
	fmt.Println("SMS", i)
	done <- true
}

func main() {
	notify := Notification{}
	for i := 1; i <= 10; i++ {
		emailDone := make(chan bool)
		smsDone := make(chan bool)

		go notify.SendEmailNotification(i, emailDone)
		<-emailDone // Wait until email is done

		go notify.SendSMSNotification(i, smsDone)
		<-smsDone // Wait until SMS is done
	}
}
