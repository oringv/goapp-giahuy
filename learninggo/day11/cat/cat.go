package cat

import "fmt"

type Cat struct {
	Name string
}

func New(name string) (*Cat, error) {
	if name == "" {
		return nil, fmt.Errorf("Ten khong duoc de trong")
	}
	return &Cat{Name: name}, nil
}

func (c *Cat) Speak() string   { return "Meo meo" }
func (c *Cat) GetName() string { return c.Name }
func (c *Cat) Eat() string     { return "Meo an ca" }
