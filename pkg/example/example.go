// Package example demonstrates go comment documentation.
package example

// CustomInt is an alternative representation of an int
type CustomInt int

const (
	// Constant represents the number 1
	Constant = 1
)

// Example returns 1
func Example() int {
	return 1
}
