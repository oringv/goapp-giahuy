package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 	/** Bai 1 **/
	// 	xhmt := ""      // Phải khai báo biến chứa chuỗi kết quả trước
	// 	isFirst := true // Dùng biến này để xử lý dấu phẩy đẹp hơn

	// 	for i := 1; i <= 100; i++ {
	// 		// Bỏ qua các số theo yêu cầu
	// 		if i == 6 || i == 48 || i == 75 || i == 89 {
	// 			continue
	// 		}

	// 		// Nếu không phải số đầu tiên, thêm dấu phẩy đằng trước
	// 		if !isFirst {
	// 			xhmt += ", "
	// 		}

	// 		xhmt += strconv.Itoa(i) // Đổi số i sang chuỗi và cộng vào xhmt
	// 		isFirst = false         // Đã qua số đầu tiên rồi
	// 	}

	// 	// In toàn bộ chuỗi kết quả ra màn hình
	// 	fmt.Println(xhmt)
	// }

	/** Bai 2 **/

	// count := 0 // Biến đếm để biết khi nào đủ 3 số
	// for i := 1; i <= 100; i++ {
	// 	if i%2 != 0 {
	// 		count++
	// 		fmt.Print(i)

	// 		// Nếu là số thứ 3 trong dòng thì xuống dòng
	// 		if count%3 == 0 {
	// 			fmt.Println()
	// 		} else {
	// 			// Nếu chưa phải số cuối cùng (99) thì mới in dấu phẩy
	// 			if i < 99 {
	// 				fmt.Print(", ")
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println()

	/** Bai 3 **/

	var batdau, ketthuc int
	var xhtml = "" // Biến này dùng để chứa toàn bộ nội dung

	fmt.Print("Vui long nhap so bat dau: ")
	fmt.Scan(&batdau)

	fmt.Print("Vui long nhap so ket thuc: ")
	fmt.Scan(&ketthuc)

	if batdau == 0 || ketthuc == 0 {
		fmt.Println("⛔ So bat dau va so ket thuc phai lon hon 0")
	} else if ketthuc < batdau {
		fmt.Println("⛔ So ket thuc phai lon hon so bat dau")
	} else {
		for k := batdau; k <= ketthuc; k++ {
			// Tiêu đề bảng cửu chương
			xhtml += "Bang cuu chuong " + strconv.Itoa(k) + "\n"

			for l := 1; l <= 10; l++ {
				// Chỉ in phép tính (k x l = result), không lặp lại chữ "Bang cuu chuong"
				xhtml += strconv.Itoa(k) + " x " + strconv.Itoa(l) + " = " + strconv.Itoa(k*l) + "\n"
			}
			xhtml += "\n" // Thêm một dòng trống giữa các bảng cho đẹp
		}

		// QUAN TRỌNG NHẤT: Phải có lệnh này thì kết quả mới hiện ra!
		fmt.Println(xhtml)
	}
}
