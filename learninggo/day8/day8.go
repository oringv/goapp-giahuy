package main

import "fmt"

func main() {
	//** Bài 24-25**//

	// // Hàm tính tổng đệ quy
	// func tong(n int) int {
	// 	if n == 1 {
	// 		return 1
	// 	}
	// 	return n + tong(n-1)
	// }

	// func main() {
	// 	for {
	// 		// 1. Hiển thị Menu
	// 		fmt.Println("\n-=-=-=-= Menu Chuc Nang -=-=-=-=")
	// 		fmt.Println("[1] Tinh tong day so N phan tu")
	// 		fmt.Println("[2] Hien thi va tinh tong day so Fibonacci")
	// 		fmt.Println("[0] Thoat chuong trinh")
	// 		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	// 		var choice int
	// 		// VÒNG LẶP CHO LỰA CHỌN MENU
	// 		for {
	// 			fmt.Print("Vui long nhap so chuc nang muon thuc hien: ")
	// 			_, err := fmt.Scan(&choice)

	// 			if err != nil {
	// 				fmt.Println("⛔ Vui long nhap mot so nguyen")
	// 				// Xóa bộ nhớ đệm khi nhập sai chữ
	// 				var discard string
	// 				fmt.Scanln(&discard)
	// 				continue
	// 			}
	// 			break
	// 		}

	// 		switch choice {
	// 		case 1:
	// 			var numb int
	// 			// VÒNG LẶP CHO VIỆC NHẬP SỐ N (Giống hệt ảnh)
	// 			for {
	// 				fmt.Print("Vui long nhap 1 so nguyen duong N: ")
	// 				_, err := fmt.Scan(&numb)

	// 				if err != nil || numb <= 0 {
	// 					fmt.Println("⛔ Vui long nhap vao la 1 so nguyen duong N")
	// 					var discard string
	// 					fmt.Scanln(&discard)
	// 					continue
	// 				}
	// 				break
	// 			}
	// 			// Tính và in kết quả
	// 			result := tong(numb)
	// 			fmt.Printf("Tong tu 1 den %d co ket qua la: %d\n", numb, result)

	// 		case 2:
	// 			fmt.Println("Chức năng Fibonacci đang phát triển...")

	// 		case 0:
	// 			fmt.Println("👋 Cam on da su dung chuong trinh. Bye !")
	// 			return

	// 		default:
	// 			fmt.Println("⚠️ Vui long chon dung so chuc nang (1, 2 hoac 0)")
	// 		}
	// 	}
	// }

	//** Bài 26**//
	name := "Pham Thanh Gia Huy"
	fmt.Println("-=-=-=-=-=Information name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", name)

	fmt.Println()

	// Tạo con tro (pointer)
	ptrName := &name
	fmt.Println("-=-=-=-=-=Information name variable -=-=-=-=-=")
	fmt.Printf("Data type: %T \n", ptrName)
	fmt.Printf("Value: %v \n", ptrName)
	fmt.Printf("Address: %v \n", ptrName)
}
