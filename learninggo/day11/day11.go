package main

import (
	"fmt"
	"goapp-giahuy/learninggo/day11/cat"
	"goapp-giahuy/learninggo/day11/dog"
	"goapp-giahuy/learninggo/day11/mouse"
	"goapp-giahuy/learninggo/day11/service"
)

func MakeSound(a service.Animal) {
	fmt.Printf("Day la tieng cua %s \n", a.GetName())
	fmt.Println(a.Speak())
}

func MakeSoundPlus(p service.AnimalPlus) {
	fmt.Printf("Day la tieng cua %s \n", p.GetName())
	fmt.Println(p.Speak())
	fmt.Println(p.Eat())
}

func main() {
	// --- 1. Con Chó ---
	mydog, err := dog.New("Bully")
	if err != nil {
		panic(err)
	}
	MakeSound(mydog)

	fmt.Println("-=-=-=-=-=-=-=-=-=-")

	// --- 2. Con Mèo ---
	mycat, err := cat.New("Pon")
	if err != nil {
		panic(err)
	}
	MakeSoundPlus(mycat)

	fmt.Println("-=-=-=-=-=-=-=-=-=-")

	// --- 3. Con Chuột ---
	mymouse, err := mouse.New("Mr Ti")
	if err != nil {
		panic(err)
	}
	// Chỉ in ra tên con chuột như trong video mẫu
	fmt.Println(mymouse.GetName())
}
