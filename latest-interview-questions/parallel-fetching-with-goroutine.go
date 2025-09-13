package main

import (
	"fmt"
)

func fetchUserData(ch chan string) {
	ch <- "User Data: Sunil Kumar"
}

func fetchOrders(ch chan string) {
	ch <- "Orders: [#123, #456, #789]"
}

func fetchNotifications(ch chan string) {
	ch <- "Notifications: [Welcome, Order Shipped]"
}

func main() {
	// Channels to collect results
	userChan := make(chan string)
	orderChan := make(chan string)
	notificationChan := make(chan string)

	// Run services concurrently
	go fetchUserData(userChan)
	go fetchOrders(orderChan)
	go fetchNotifications(notificationChan)

	// Collect results
	// Print combined response

	userData := <-userChan
	orders := <-orderChan
	notifications := <-notificationChan

	fmt.Println("Aggregated Response:")
	fmt.Println(userData)
	fmt.Println(orders)
	fmt.Println(notifications)
}
