// The Prototype Pattern is used when you need to create new objects by copying existing ones,
// rather than building them from scratch. It’s especially useful when object creation is
// expensive or complex.

// Scenario

// Imagine your system sends standardized JSON responses for different APIs (success, error, not found, etc.).
// Instead of building each response from scratch, you maintain prototype templates and clone them when needed.
package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Document struct {
	Title   string
	Content []string
}

func (d *Document) Clone() Prototype {
	// Deep copy of the slice to avoid shared reference
	contentCopy := make([]string, len(d.Content))
	copy(contentCopy, d.Content)

	return &Document{
		Title:   d.Title,
		Content: contentCopy,
	}
}

func (d *Document) Print() {
	fmt.Println("Title:", d.Title)
	fmt.Println("Content:", d.Content)
}

func main() {
	original := &Document{
		Title:   "Original",
		Content: []string{"Line 1", "Line 2"},
	}

	clone := original.Clone().(*Document) // Type assertion to *Document

	// Modify the clone
	clone.Title = "Cloned"
	clone.Content[0] = "Updated Line 1"

	fmt.Println("Original Document:")
	original.Print()

	fmt.Println("\nCloned Document:")
	clone.Print()
}
