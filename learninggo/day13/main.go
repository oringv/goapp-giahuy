package main

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/student"
	"goapp-giahuy/learninggo/day13/teacher"
	"goapp-giahuy/learninggo/day13/utils"
	"os"
)

func main() {
	for {
		utils.ClearScreen()
		fmt.Println("======== 📚 HỆ THỐNG QUẢN LÝ ========")
		fmt.Println("1. Quản lý Sinh viên")
		fmt.Println("2. Quản lý Giảng viên")
		fmt.Println("3. Thoát chương trình")
		fmt.Println("====================================")

		choice := utils.GetPositiveInt("👉 Chọn chức năng: ")

		switch choice {
		case 1:
			student.SubMenu()
		case 2:
			teacher.MenuQuanLyGiangVien()
		case 3:
			fmt.Println("👋 Tạm biệt! Tác giả: Gia Huy")
			os.Exit(0)
		default:
			fmt.Println("⚠️ Lựa chọn không hợp lệ!")
			utils.ReadInput("\nNhấn Enter để chọn lại...")
		}
	}
}
