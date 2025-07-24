//  In Go, select { ... } is used to wait on multiple channel operations.
//  select default means:
// It provides a non-blocking way to use select.

select {
case msg := <-ch:
	fmt.Println("Received:", msg)
default:
	fmt.Println("No message received")
}

// Explanation:
// If ch has a value ready, it will read from it.

// If not, the default block runs immediately, without waiting.

