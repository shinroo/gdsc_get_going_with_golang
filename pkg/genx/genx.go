package genx

import "fmt"

//go:generate ./newList.sh Person

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}
