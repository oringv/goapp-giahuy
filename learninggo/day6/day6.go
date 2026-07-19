package main

import (
	"bufio"
	"fmt"
	"os"
)

var address = "Ho Chi Minh City"

var (
	monhoc   string
	thongtin string
)

func main() {
	//mt.Println("Hello Gia Huy")
	//randomUser()

	// var fullName = "Pham Thanh Gia Huy"
	// fullName = "Tony Teo"
	// fmt.Println(fullName)

	// var age int
	// fmt.Println(age

	// phone := "0123456789"
	// fmt.Println(phone)

	//var toan, tiengviet, tunhien int
	//toan = 9
	//tiengviet = 8
	//tunhien = 7
	//fmt.Println(toan)
	//fmt.Println(tiengviet)
	//fmt.Println(tunhien)

	// toan, tiengviet, tunhien := 9, 8, 7
	// fmt.Println(toan)
	// fmt.Println(tiengviet)
	// fmt.Println(tunhien)

	// monhoc = "Lap trinh golang"
	// fmt.Println(monhoc)

	// thongtin = "Sinh vien: Gia Huy"
	// fmt.Println(thongtin)

	// var hoten string
	// fmt.Println("Nhap ho ten: ")
	// fmt.Scanln(&hoten)
	// fmt.Println("Ho ten cua ban la: ", hoten)

	var hoten string
	fmt.Print("Nhap ho ten:  ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		hoten = scanner.Text()
	}
	fmt.Println("Ho ten cua ban la: ", hoten)
}
