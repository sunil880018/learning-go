package main

import (
	"fmt"
	"time"
)

// User struct represents a user in the ticket booking system
type User struct {
	ID   string
	Name string
}

// BookTicket allows a user to book seats for an event
func (u *User) BookTicket(event *Event, seats int) (Booking, error) {
	return event.BookSeats(u, seats)
}

// CancelBooking allows a user to cancel a booking
func (u *User) CancelBooking(event *Event, bookingID string) (string, error) {
	return event.CancelBooking(u, bookingID)
}

// Booking struct represents a booking for an event
type Booking struct {
	BookingID string
	EventID   string
	UserID    string
	Seats     int
}

// Event struct represents an event for which tickets can be booked
type Event struct {
	ID             string
	Name           string
	TotalSeats     int
	AvailableSeats int
	Bookings       []Booking
}

// BookSeats books seats for a user if seats are available
func (e *Event) BookSeats(user *User, seats int) (Booking, error) {
	if seats > e.AvailableSeats {
		return Booking{}, fmt.Errorf("only %d seats available", e.AvailableSeats)
	}
	bookingID := fmt.Sprintf("%s-%d", e.ID, time.Now().UnixNano())
	booking := Booking{BookingID: bookingID, EventID: e.ID, UserID: user.ID, Seats: seats}
	e.Bookings = append(e.Bookings, booking)
	e.AvailableSeats -= seats
	return booking, nil
}

// CancelBooking cancels a booking if it exists and was made by the same user
func (e *Event) CancelBooking(user *User, bookingID string) (string, error) {
	for i, booking := range e.Bookings {
		if booking.BookingID == bookingID && booking.UserID == user.ID {
			e.Bookings = append(e.Bookings[:i], e.Bookings[i+1:]...)
			e.AvailableSeats += booking.Seats
			return fmt.Sprintf("Booking %s cancelled successfully", bookingID), nil
		}
	}
	return "", fmt.Errorf("booking not found or unauthorized")
}

// GetAvailableSeats returns the number of available seats
func (e *Event) GetAvailableSeats() int {
	return e.AvailableSeats
}

// BookingSystem struct manages events and users
type BookingSystem struct {
	Events []Event
	Users  []User
}

// AddEvent adds a new event to the booking system
func (bs *BookingSystem) AddEvent(name string, totalSeats int) Event {
	eventID := fmt.Sprintf("%s-%d", name, time.Now().UnixNano())
	event := Event{ID: eventID, Name: name, TotalSeats: totalSeats, AvailableSeats: totalSeats}
	bs.Events = append(bs.Events, event)
	return event
}

// AddUser adds a new user to the booking system
func (bs *BookingSystem) AddUser(name string) User {
	userID := fmt.Sprintf("%s-%d", name, time.Now().UnixNano())
	user := User{ID: userID, Name: name}
	bs.Users = append(bs.Users, user)
	return user
}

// FindEventByID finds an event by its ID
func (bs *BookingSystem) FindEventByID(eventID string) *Event {
	for i := range bs.Events {
		if bs.Events[i].ID == eventID {
			return &bs.Events[i]
		}
	}
	return nil
}

// FindUserByID finds a user by their ID
func (bs *BookingSystem) FindUserByID(userID string) *User {
	for i := range bs.Users {
		if bs.Users[i].ID == userID {
			return &bs.Users[i]
		}
	}
	return nil
}

func main() {
	bookingSystem := BookingSystem{}
	event1 := bookingSystem.AddEvent("Concert", 100)
	user1 := bookingSystem.AddUser("Alice")
	user2 := bookingSystem.AddUser("Bob")

	fmt.Println("Initial available seats:", event1.GetAvailableSeats())

	booking1, err := user1.BookTicket(&event1, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Booking 1:", booking1)
		fmt.Println("Available seats after booking 1:", event1.GetAvailableSeats())
	}

	booking2, err := user2.BookTicket(&event1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Booking 2:", booking2)
		fmt.Println("Available seats after booking 2:", event1.GetAvailableSeats())
	}

	cancelResult, err := user1.CancelBooking(&event1, booking1.BookingID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(cancelResult)
		fmt.Println("Available seats after cancellation:", event1.GetAvailableSeats())
	}
}
