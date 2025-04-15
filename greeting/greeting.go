// Package greeting provides functions for creating greetings.
package greeting // This package name must match the directory name

import "fmt"

// GetGreeting returns a greeting for the named person.
// Note: The function name starts with an uppercase letter to make it exported (public).
func GetGreeting(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// You could add other functions here specific to greetings.
// func GetFarewell(name string) string { ... }
