package main

import (
	"fmt"
	"strconv" // Thêm thư viện này để đổi số sang chữ
)

func main() {
	/** Bai 1 **/
	xhmt := ""      // Phải khai báo biến chứa chuỗi kết quả trước
	isFirst := true // Dùng biến này để xử lý dấu phẩy đẹp hơn

	for i := 1; i <= 100; i++ {
		// Bỏ qua các số theo yêu cầu
		if i == 6 || i == 48 || i == 75 || i == 89 {
			continue
		}

		// Nếu không phải số đầu tiên, thêm dấu phẩy đằng trước
		if !isFirst {
			xhmt += ", "
		}

		xhmt += strconv.Itoa(i) // Đổi số i sang chuỗi và cộng vào xhmt
		isFirst = false         // Đã qua số đầu tiên rồi
	}

	// In toàn bộ chuỗi kết quả ra màn hình
	fmt.Println(xhmt)
}

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
