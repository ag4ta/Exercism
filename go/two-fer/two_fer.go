// Package twofer provides tools to display "One for X, one for me"
package twofer

import "fmt"

// ShareWith returns a string with "One for X, one for me" where X is replaced with given name or with 'you' if it's empty
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
