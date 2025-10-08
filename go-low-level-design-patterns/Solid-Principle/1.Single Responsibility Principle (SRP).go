// A struct or function should have only one reason to change, meaning it should focus on a single responsibility.
package main

import (
	"fmt"
)

// Single Responsibility: User struct is responsible only for user data.
type User struct {
	Name  string
	Email string
}

// Single Responsibility: Logger is only responsible for logging.
type Logger struct{}

func (l *Logger) Log(info string) {
	fmt.Println("Log:", info)
}

// Single Responsibility: UserService is responsible for user operations.
type UserService struct {
	logger *Logger
}

func (s *UserService) CreateUser(user User) {
	fmt.Printf("User created: %s with email %s\n", user.Name, user.Email)
	s.logger.Log("User created: " + user.Name)
}

func main() {
	logger := &Logger{}
	userService := &UserService{logger}
	user := User{Name: "Alice", Email: "alice@example.com"}
	userService.CreateUser(user)
}
