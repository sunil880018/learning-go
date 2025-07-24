✅ What is context in Go?
The context package in Go is used to carry deadlines, timeouts, and cancellation 
signals across API boundaries and goroutines.

🧠 Why do we need it?
Let's say you're making an API request to a server or database:

What if the client cancels the request?

Or it takes too long?
Or you want to pass some data (like request ID)?

We use context so we can:

Cancel work cleanly (avoid memory leaks)
Set timeouts
Pass request-scoped values


| Use Case          | Example                                      |
| ----------------- | -------------------------------------------- |
| Timeout           | Cancel DB query after 2 seconds              |
| Cancel goroutines | Stop worker if main request is cancelled     |
| Pass values       | Share request ID or user info through layers |

package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Finished work")
    case <-ctx.Done():
        fmt.Println("Timeout:", ctx.Err()) // Output: Timeout: context deadline exceeded
    }
}

In the above:

We set a 2-second timeout
But the work takes 3 seconds
So context cancels it


func fetchData(ctx context.Context) {
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Fetched data")
    case <-ctx.Done():
        fmt.Println("Stopped early:", ctx.Err())
    }
}

🧵 When you spawn goroutines, always pass context
Cancel them
Avoid goroutines running forever
