package mouse

import "fmt"

type Mouse struct {
	Name string
}

func New(name string) (*Mouse, error) {
	if name == "" {
		return nil, fmt.Errorf("Ten chuot khong duoc de trong")
	}
	return &Mouse{Name: name}, nil
}

func (m *Mouse) Speak() string   { return "Chit chit!" }
func (m *Mouse) GetName() string { return m.Name }
