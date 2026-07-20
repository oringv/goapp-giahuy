package main

import "fmt"

func tong(n int) int {
	// Dieu kien dung
	if n == 1 {
		return 1
	}
	// Goi de quy
	return n + tong(n-1)
}

func fibo(n int) {
	term1, term2 := 0, 1

	if n == 1 {
		fmt.Printf("Day so fibonacci: %d \n", term1)
		fmt.Printf("Tong day so fibonacci: %d \n", term1)
	} else if n == 2 {
		fmt.Printf("Day so fibonacci: %d %d  \n", term1, term2)
		fmt.Printf("Tong day so fibonacci: %d \n", term1+term2)
	} else {

		total := term1 + term2

		fmt.Printf("Day sp fibonacci: %d %d", term1, term2)

		for i := 3; i <= n; i++ {
			term3 := term1 + term2

			fmt.Printf("%d", term3)
			total += term3
			term1, term2 = term2, term3
		}

		fmt.Println()

		fmt.Printf("Tong day so fibonacci: %d \n", total)
	}
}
func main() {
	// Dùng vòng lặp for vô tận để giữ Menu luôn hiện ra
	for {
		fmt.Println("\n-=-=-=-= Menu Chuc Nang-=-=-=-= ")
		fmt.Println("[1] Tinh tong day so 1 den N")
		fmt.Println("[2] Hien thi va tinh tong day so Fibonacci")
		fmt.Println("[0] Thoat chuong trinh")
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

		fmt.Print("Vui long nhap so chuc nang: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var numb int
			fmt.Print("Nhập số nguyên dương N: ")
			_, err := fmt.Scan(&numb)

			if err != nil || numb <= 0 {
				fmt.Println("❌ Lỗi: Vui lòng nhập số nguyên dương!")
				continue // Quay lại đầu vòng lặp Menu
			}

			// GỌI HÀM TÍNH TỔNG (Bạn đã thiếu phần này)
			result := tong(numb)
			fmt.Printf("⭐ Kết quả: Tổng từ 1 đến %d là: %d\n", numb, result)

		case 2:
			var numb int
			fmt.Print("Nhập số lượng số Fibonacci: ")
			_, err := fmt.Scan(&numb)

			if err != nil || numb <= 0 {
				fmt.Println("❌ Lỗi: Số lượng phải lớn hơn 0!")
				continue
			}
			fibo(numb)

		case 0:
			fmt.Println("👋 Cam on da su dung chuong trinh. Bye !")
			return // Thoát hẳn chương trình

		default:
			fmt.Println("⚠️ Vui lòng chọn đúng số chức năng (1, 2 hoặc 0)")
		}
	}
}
