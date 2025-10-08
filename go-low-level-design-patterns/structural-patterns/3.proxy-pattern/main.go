// Proxy Design Pattern in Go is used to provide a surrogate or placeholder for another object to
// control access to it. It’s commonly used for lazy initialization, access control, logging,
// caching, or remote proxies.

// Scenario

// Suppose you have a heavy image object (e.g., large image files) that takes time to load.
// You want to load the image only when needed, without changing the client code. This is a
// classic virtual proxy use case.

package main

import "fmt"

// Subject interface
type Image interface {
	Display()
}

// RealSubject — actual image (heavy object)
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading image from disk:", filename)
	return &RealImage{filename: filename}
}

func (ri *RealImage) Display() {
	fmt.Println("Displaying image:", ri.filename)
}

// Proxy — controls access to RealImage
type ProxyImage struct {
	filename string
	real     *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (pi *ProxyImage) Display() {
	if pi.real == nil {
		pi.real = NewRealImage(pi.filename) // lazy initialization
	}
	fmt.Println("Proxy delegating display to RealImage")
	pi.real.Display()
}

// Client code
func main() {
	image1 := NewProxyImage("photo1.jpg")
	image2 := NewProxyImage("photo2.jpg")

	// Images are not loaded yet
	fmt.Println("First call to display image1:")
	image1.Display() // Loads and displays

	fmt.Println("\nSecond call to display image1:")
	image1.Display() // Already loaded, just displays

	fmt.Println("\nCall to display image2:")
	image2.Display() // Loads and displays

	fmt.Println("\nSecond call to display image2:")
	image2.Display() // Already loaded, just displays
}

// Explanation

// Subject (Image) → Interface used by the client.

// RealSubject (RealImage) → The actual heavy object.

// Proxy (ProxyImage) → Controls access to RealImage, e.g., lazy loading.

// Client → Works with Image interface and doesn’t know whether it’s dealing with a proxy or the real object.
