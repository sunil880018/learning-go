// The Observer Design Pattern in Go is used when you want to notify multiple objects
// automatically when the state of another object changes. It’s commonly used in event systems,
// messaging, and real-time updates.

// 🧩 Scenario

// Suppose you have a stock price system and multiple clients (e.g., dashboards, alerts)
// need to get notified whenever the stock price changes.

package main

import "fmt"

// Observer interface
type Observer interface {
	Update(price float64)
}

// Subject interface
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify()
}

// Concrete subject — Stock
type Stock struct {
	name      string
	price     float64
	observers []Observer
}

func (s *Stock) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Stock) Unregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Stock) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.price)
	}
}

// Set new price and notify observers
func (s *Stock) SetPrice(price float64) {
	fmt.Printf("\nStock %s price changed to $%.2f\n", s.name, price)
	s.price = price
	s.Notify()
}

// Concrete observer — Dashboard
type Dashboard struct {
	id string
}

func (d *Dashboard) Update(price float64) {
	fmt.Printf("[Dashboard %s] Stock price updated: $%.2f\n", d.id, price)
}

// Concrete observer — Alert system
type Alert struct {
	threshold float64
}

func (a *Alert) Update(price float64) {
	if price > a.threshold {
		fmt.Printf("[ALERT] Stock price crossed threshold: $%.2f\n", price)
	}
}

// Client code
func main() {
	stock := &Stock{name: "GOOG"}

	d1 := &Dashboard{id: "D1"}
	d2 := &Dashboard{id: "D2"}
	alert := &Alert{threshold: 1500.0}

	stock.Register(d1)
	stock.Register(d2)
	stock.Register(alert)

	stock.SetPrice(1490.0)
	stock.SetPrice(1510.0)
}

// Explanation

// Subject (Stock) → Maintains a list of observers and notifies them on state changes.

// Observers (Dashboard, Alert) → Implement the Update() method and react to notifications.

// Client → Registers observers with the subject and changes the stock price.

// Decoupling → Stock doesn’t know details about dashboards or alerts — only calls Update().

// ⚡ Real-world Use Cases

// Stock market tickers.

// Social media notifications (followers get updates).

// Event-driven architectures (Kafka consumers, WebSockets).

// Logging, metrics, or monitoring systems.
