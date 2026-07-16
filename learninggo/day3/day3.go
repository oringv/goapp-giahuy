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

	// message := fmt.Sprint("My name is ", "Huy")
	// fmt.Println(message)

	// message := fmt.Sprintln("My name is ", "Huy")
	// fmt.Println(message)

	// name := "Huy"
	// age := 22
	// message := fmt.Sprintf("My name is %s and age is %d", name, age)
	// fmt.Println(message)

	// ten := "Pham Thanh Gia Huy"
	// tuoi := 22
	// chieucao := 1.62
	// daTotNghiep := true
	// phanTram := 10

	// fmt.Printf("Kieu du lieu cua bien ten:  %T \n", ten)
	// fmt.Printf("Kieu du lieu cua bien tuoi:  %T \n", tuoi)
	// fmt.Printf("Kieu du lieu cua bien chieu cao:  %T \n", chieucao)
	// fmt.Printf("Kieu du lieu cua bien daTotNghiep:  %T \n", daTotNghiep)
	// fmt.Printf("Kieu du lieu cua bien phanTram:  %T \n", phanTram)

	// fmt.Printf("Toi ten la: %v \n", ten)

	// fmt.Printf("Toi ten la %s va toi %d tuoi \n", ten, tuoi)

	// fmt.Printf("Chieu cao cua toi la %.2fm \n", chieucao)
	// fmt.Printf("Chieu cao cua toi la %.5fm \n", chieucao)
	// fmt.Printf("Chieu cao cua toi la %.1fm \n", chieucao)

	// fmt.Printf("Toi da tot nghiep: %t \n", daTotNghiep)

	// fmt.Printf("Toi da hoc duoc %d%% \n", phanTram)

	s1 := 15
	s2 := 15

	phepsosanh := s1 <= s2
	fmt.Printf("Ket qua phep so sanh: %t \n", phepsosanh)

	phepluanly := !((5 > 6) && (5 > 3) && (5 > 4))
	fmt.Printf("Ket qua phep luan ly: %t \n", phepluanly)

	// tong := s1 + s2
	// hieu := s1 - s2
	// tich := s1 * s2
	// thuong := float32(s1) / float32(s2)
	// chialaydu := s1 % s2
	// s1 += 5

	// fmt.Printf("Tong cua %d va %d la: %d \n", s1, s2, tong)
	// fmt.Printf("Hieu cua %d va %d la: %d \n", s1, s2, hieu)
	// fmt.Printf("Tich cua %d va %d la: %d \n", s1, s2, tich)
	// fmt.Printf("Thuong cua %d va %d la: %.2f \n", s1, s2, thuong)
	// fmt.Printf("Chia lay du cua %d va %d la: %d \n", s1, s2, chialaydu)
	// fmt.Printf("S1 %d \n", s1)

	// s3 := 15
	// s3 %= 4
	// fmt.Printf("S3: %d \n", s3)
}
