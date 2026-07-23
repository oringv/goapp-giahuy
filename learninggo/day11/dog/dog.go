package dog

import "fmt"

type Dog struct {
	Name string
}

// Hàm khởi tạo Dog giống video
func New(name string) (*Dog, error) {
	if name == "" {
		return nil, fmt.Errorf("Ten khong duoc de trong")
	}
	return &Dog{Name: name}, nil
}

func (d *Dog) Speak() string   { return "Gau gau!" }
func (d *Dog) GetName() string { return d.Name }
func (d *Dog) Eat() string     { return "Cho an xuong" }
