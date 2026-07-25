package main

import (
	"fmt"
	// CẦN THAY ĐỔI ĐƯỜNG DẪN NÀY CHO ĐÚNG VỚI DỰ ÁN CỦA BẠN
	"goapp-giahuy/learninggo/day13/student"
	"goapp-giahuy/learninggo/day13/teacher"
	"goapp-giahuy/learninggo/day13/utils"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("==== HỆ THỐNG QUẢN LÝ ====")
		fmt.Println("1. Quản lý Sinh viên")
		fmt.Println("2. Quản lý Giảng viên")
		fmt.Println("3. Thoát")

		choice := utils.ReadInt("👉 Chọn chức năng: ")

		switch choice {
		case 1:
			student.SubMenu() // Gọi menu con của sinh viên
		case 2:
			teacher.MenuQuanLyGiangVien() // Gọi menu con của giảng viên
		case 3:
			fmt.Println("👋 Tạm biệt!")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ.")
			utils.ReadString("\nNhấn Enter để tiếp tục...")
		}
	}
}
