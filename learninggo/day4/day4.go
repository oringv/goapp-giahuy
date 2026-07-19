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

	// diem := -3

	// if diem > 8 {
	// 	fmt.Println("Hoc sinh gioi")
	// } else if diem <= 8 && diem >= 6 { // Sửa ≤ thành <=
	// 	fmt.Println("Hoc sinh kha")
	// } else if diem <= 5 && diem >= 3 { // Sửa ≤ thành <=
	// 	fmt.Println("Hoc sinh trung binh")
	// } else {
	// 	fmt.Println("Hoc sinh yeu")
	// }

	// monan := "pho"
	// switch monan {
	// case "com":
	// 	fmt.Println("Mon an la com")
	// case "pho":
	// 	fmt.Println("Mon an la pho")
	// case "bun":
	// 	fmt.Println("Mon an la bun")
	// default:
	// 	fmt.Println("Khong an gi ca")

	// }

	// Một số điểm 10
	i := 1
	for i <= 10 {
		fmt.Printf("Number %d \n", i)
		i++
	}

	// Vòng lặp chạy từ 1 đến 100
	// for i := 1; i <= 100; i++ {

	// 	// Kiểm tra nếu i chia hết cho 9 (số dư bằng 0)
	// 	if i%9 == 0 {
	// 		fmt.Printf("Number %d \n", i)
	// 	}
	// }
	// 1. Ví dụ về BREAK
	// fmt.Println("--- Test BREAK ---")
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		break // Gặp số chẵn đầu tiên (số 2) là thoát hẳn vòng lặp
	// 	}
	// 	fmt.Printf("Number %d \n", i)
	// }
	// fmt.Println("End code")

	// fmt.Println("\n--- Test CONTINUE ---")
	// 2. Ví dụ về CONTINUE
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Gặp số chẵn thì bỏ qua, nhảy sang số tiếp theo
		}
		fmt.Printf("Number %d \n", i)
	}
}
