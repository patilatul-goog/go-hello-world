package main

import (
	"fmt"
	// Import the local greeting package using the module path + subdirectory
	"github.com/patilatul-goog/go-hello-world/greeting" // Make sure 'example.com/hello' matches your go.mod module path

	// Import a third-party package
	"rsc.io/quote"
)

// main is the entry point of the program.
func main() {
	// Get a greeting from the local package
	message := greeting.GetGreeting("Gladys")
	fmt.Println(message)

	// Get a quote from the third-party package
	fmt.Println("And here's a Go proverb:")
	fmt.Println(quote.Go())
}
