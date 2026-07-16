package main

import "fmt"

const MONAN = "Chao"

func main() {

	// const MONAN = "Com tam"
	// fmt.Println(MONAN)

	// fmt.Print("Hello Gia Huy")
	// fmt.Print("Gia Huy")

	// fmt.Println("Hello Gia Huy")
	// fmt.Println("Gia Huy")

	// var fullName = "Pham Thanh Gia Huy"
	// var age = 22
	// fmt.Printf("Hello %s and your age is %d. \n", fullName, age)

	// var firstName, lastName string
	// fmt.Print("Please enter your name: ")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Println("My name is", firstName, lastName)

	// var firstName, lastName string
	// fmt.Print("Please enter your first name: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Please enter your last name: ")
	// fmt.Scanln(&lastName)
	// fmt.Println("My name is", firstName, lastName)

	var name string
	var age int
	fmt.Println("Please enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Please enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("My name is %s and age is %d \n", name, age)

}
