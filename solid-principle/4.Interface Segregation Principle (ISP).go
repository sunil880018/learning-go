// Clients should not be forced to implement interfaces they do not use. Instead of a large interface,
// split it into smaller, more specific ones.
package main

import (
	"fmt"
)

// Define smaller interfaces instead of one large one
type Printer interface {
	Print() error
}

type Scanner interface {
	Scan() error
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print() error {
	fmt.Println("Printing document...")
	return nil
}

func (m *MultiFunctionPrinter) Scan() error {
	fmt.Println("Scanning document...")
	return nil
}

func main() {
	printer := &MultiFunctionPrinter{}

	// We can use printer as both a Printer and Scanner independently
	printer.Print()
	printer.Scan()
}

// Printer and Scanner interfaces are separate, so a struct can implement only what it needs.
