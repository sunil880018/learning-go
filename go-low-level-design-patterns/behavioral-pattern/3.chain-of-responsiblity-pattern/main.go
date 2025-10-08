// The Chain of Responsibility (CoR) Pattern in Go allows you to pass a request along a chain of handlers.
// Each handler can process the request, modify it, or pass it along. It’s widely used for middleware,
// validation, and request pipelines.

// 🧩 Scenario

// Suppose you have an HTTP request pipeline where requests must go through:

// Authentication

// Logging

// Validation

// Instead of calling each handler manually, you can chain them dynamically.

package main

import "fmt"

// Handler interface
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

// BaseHandler — optional default implementation
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// Concrete handler 1 — Authentication
type AuthHandler struct {
	BaseHandler
}

func (h *AuthHandler) Handle(request string) {
	fmt.Println("[Auth] Checking authentication for:", request)
	// Call next in chain
	h.BaseHandler.Handle(request)
}

// Concrete handler 2 — Logging
type LoggingHandler struct {
	BaseHandler
}

func (h *LoggingHandler) Handle(request string) {
	fmt.Println("[Logging] Logging request:", request)
	// Call next in chain
	h.BaseHandler.Handle(request)
}

// Concrete handler 3 — Validation
type ValidationHandler struct {
	BaseHandler
}

func (h *ValidationHandler) Handle(request string) {
	fmt.Println("[Validation] Validating request:", request)
	h.BaseHandler.Handle(request)
}

// Client code
func main() {
	// Create handlers
	auth := &AuthHandler{}
	logging := &LoggingHandler{}
	validation := &ValidationHandler{}

	// Chain them: auth -> logging -> validation
	auth.SetNext(logging).SetNext(validation)

	// Start handling a request
	auth.Handle("POST /api/payment")
}

// Explanation

// Handler interface (Handler) → Defines SetNext and Handle.

// Concrete handlers (AuthHandler, LoggingHandler, ValidationHandler) → Perform actions and optionally call next.

// Client → Builds a chain and sends requests to the first handler.

// Decoupling → Handlers don’t need to know the full chain, just the next handler.

// ⚡ Real-world Use Cases

// Web middleware — Express, Gin, or HTTP request pipelines.

// Authorization/Validation — Check permissions, validate data, log requests sequentially.

// Event handling — Sequential processing in event-driven systems.

// Payment workflows — Check fraud → Validate funds → Process payment → Notify.
