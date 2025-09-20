package main

import "fmt"

// Product
type User struct {
	FirstName string
	LastName  string
	Age       int
	City      string
	Country   string
	Email     string
	Phone     string
}

// Builder
type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) WithName(first, last string) *UserBuilder {
	b.user.FirstName = first
	b.user.LastName = last
	return b
}

func (b *UserBuilder) WithAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) WithAddress(city, country string) *UserBuilder {
	b.user.City = city
	b.user.Country = country
	return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) WithPhone(phone string) *UserBuilder {
	b.user.Phone = phone
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}

// Usage
func main() {
	user := NewUserBuilder().
		WithName("Sunil", "Kumar").
		WithAge(28)
		// WithAddress("Delhi", "India").
		// WithEmail("sunil@example.com").
		// Build()

	fmt.Printf("%+v\n", user)
}

// ✅ Concept

// 	Used to construct complex objects step by step.
// 	Helps when:
// 	The object has many optional fields/configurations.
// 	You want to avoid huge constructors with dozens of parameters.

// Instead of:

// 	user := User{
//   FirstName: "Sunil",
//   LastName:  "Kumar",
//   Age:       28,
//   City:      "Delhi",
//   Country:   "India",
//   Phone:     "9999999999",
//   Email:     "x@y.com",
// }

// You do:
// user := NewUserBuilder().
//           WithName("Sunil", "Kumar").
//           WithAge(28).
//           WithAddress("Delhi", "India").
//           WithEmail("x@y.com").
//           Build()

// Real-World Go Use Cases for Builder Pattern

// 	Configuration objects

// 	Example: HTTP client, Kafka producer, Redis client → they often have optional params.

// 	SQL Query builder

// 	Building queries dynamically (squirrel, gorm do this internally).

// 	Complex struct creation

// 	Example: Creating a User with optional fields (address, phone, email).

// 	Infrastructure setup

// 	Creating cloud resources (EC2 with optional VPC, tags, IAM roles).
