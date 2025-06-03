// 1. What is Go (Golang)?
// Briefly describe Go and its key features (compiled, statically typed, concurrency support, simplicity, etc.).

// 2. What are the key features of Golang?
// Simplicity

// Concurrency (goroutines, channels)

// Fast compilation

// Garbage Collection

// Strong standard library

// Cross-platform support

// 3. What is the difference between var, :=, and const in Go?
// var declares a variable with a type.

// := is shorthand for declaring and initializing a variable (only inside functions).

// const declares a constant value.

// 4. What is a goroutine?
// Lightweight thread managed by Go runtime. Created using the go keyword.

// 5. What is a channel in Go?
// A channel is a way to communicate between goroutines — allows them to synchronize and share data safely.

// 6. What is a slice in Go? How is it different from an array?
// Slice is a dynamically-sized, flexible view into an array.

// Arrays are fixed in size; slices are more commonly used in real applications.

// 7. Explain defer, panic, and recover in Go.
// defer: delays execution of a function until the surrounding function returns.

// panic: stops normal execution and starts panicking.

// recover: regains control of a panicking goroutine.

// 8. How is error handling done in Go?
// Go uses error values (returning error type) instead of exceptions. Example:

// if err != nil {
//   // handle error
// }
// 9. What is a struct in Go?
// Struct is a collection of fields. Similar to classes but without methods.

// 10. What is an interface in Go?
// Interface defines a set of method signatures. A type implements an interface by implementing its methods.

// 11. What is the purpose of the blank identifier _ in Go?
// Used to ignore values or imports.

// 12. How is concurrency achieved in Go?
// Through goroutines and channels.

// 13. What is the difference between buffered and unbuffered channels?
// Unbuffered channels block until both sender and receiver are ready.

// Buffered channels can store a fixed number of elements without blocking immediately.

// 14. How do you create a custom package in Go?
// Create a folder with .go files, declare package name, and expose functions/variables with capital letters.

// 15. What are Go’s visibility rules (public/private)?
// If an identifier starts with an uppercase letter, it's exported (public); otherwise, it's unexported (private).

// 16. What is the select statement in Go?
// Used with channels to wait on multiple channel operations.

// 17. What is the difference between pointer and value receivers in Go methods?
// Value receiver copies the value; pointer receiver allows modifying the original value.

// 18. How is memory management handled in Go?
// Go has built-in garbage collection.

// 19. What are maps in Go?
// Built-in associative data type (key-value pairs).

// 20. How do you handle JSON in Go?
// Use encoding/json package with Marshal and Unmarshal functions.

