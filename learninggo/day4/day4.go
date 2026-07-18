package main

import "fmt"

func main() {
	// diem := 8

	// if diem >= 8 { // Thêm dấu bằng ở đây
	// 	fmt.Println("Diem lon hon hoac bang 8")
	// } else {
	// 	fmt.Println("Diem khong lon hon 8")
	// }
	// fmt.Println("Out !")

	diem := -3

	if diem > 8 {
		fmt.Println("Hoc sinh gioi")
	} else if diem <= 8 && diem >= 6 { // Sửa ≤ thành <=
		fmt.Println("Hoc sinh kha")
	} else if diem <= 5 && diem >= 3 { // Sửa ≤ thành <=
		fmt.Println("Hoc sinh trung binh")
	} else {
		fmt.Println("Hoc sinh yeu")
	}
}
