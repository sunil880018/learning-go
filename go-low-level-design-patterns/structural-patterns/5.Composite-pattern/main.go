// The Composite Design Pattern in Go lets you treat individual objects and groups of objects uniformly.
// It’s great for tree-like structures, e.g., file systems, UI components, or menu hierarchies

// Scenario

// Suppose you want to represent a file system with files and folders. Both should support a Display()
// method, and folders can contain files or other folders.

package main

import "fmt"

// Component interface
type FileSystemItem interface {
	Display(indent string)
}

// Leaf — File
type File struct {
	name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent+"File:", f.name)
}

// Composite — Folder
type Folder struct {
	name     string
	children []FileSystemItem
}

func (f *Folder) Add(item FileSystemItem) {
	f.children = append(f.children, item)
}

func (f *Folder) Display(indent string) {
	fmt.Println(indent+"Folder:", f.name)
	for _, child := range f.children {
		child.Display(indent + "  ")
	}
}

// Client code
func main() {
	// Create leaf nodes
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}

	// Create folders (composite)
	folder1 := &Folder{name: "Documents"}
	folder2 := &Folder{name: "Pictures"}

	// Add files to folders
	folder1.Add(file1)
	folder1.Add(file2)

	file3 := &File{name: "image1.png"}
	folder2.Add(file3)

	// Root folder
	root := &Folder{name: "Root"}
	root.Add(folder1)
	root.Add(folder2)

	// Display the entire file system
	root.Display("")
}

// Explanation

// Component (FileSystemItem) → Common interface for leaf and composite.

// Leaf (File) → Basic object with no children.

// Composite (Folder) → Contains children, which can be files or other folders.

// Client → Can treat files and folders uniformly.
