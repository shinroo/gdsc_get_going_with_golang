// Package genx shows off Go's code generation tool go:generate
package genx

import "fmt"

//go:generate ./newList.sh Person

// Person represents a person's name and age
type Person struct {
	Name string
	Age  int
}

// String returns a string representation of a person
func (p *Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}
