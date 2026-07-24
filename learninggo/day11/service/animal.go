package service

type Animal interface {
	Speak() string
	GetName() string
}

type AnimalPlus interface {
	Animal // Kế thừa lại Animal (có Speak và GetName)
	Eat() string
}
