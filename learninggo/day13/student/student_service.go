package student

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/utils" // Giả sử đây là đường dẫn import của bạn
)

// Danh sách sinh viên toàn cục (Chỉ khai báo ở đây)
var List []Student

// Menu con quản lý sinh viên (Chỉ khai báo ở đây)
func SubMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==== QUẢN LÝ SINH VIÊN ====")
		fmt.Println("1. Thêm")
		fmt.Println("2. Danh sách")
		fmt.Println("3. Quay lại")
		c := utils.ReadInt("👉 Chọn chức năng: ")

		if c == 1 {
			handleThemSinhVien()
		} else if c == 2 {
			handleHienThiDanhSach()
		} else if c == 3 {
			return // Thoát ra menu chính
		} else {
			fmt.Println("❌ Lựa chọn không hợp lệ.")
		}
		utils.ReadString("\nNhấn Enter để tiếp tục...")
	}
}

func handleThemSinhVien() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Them Sinh Vien -=-=-=-=-=-")
	id := utils.ReadInt("- Nhập ID: ")
	name := utils.ReadString("- Nhập Tên: ")

	// Thêm vào slice List
	List = append(List, Student{ID: id, Name: name})

	fmt.Println("✅ Thêm sinh viên thành công!")
}

func handleHienThiDanhSach() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Danh Sach Sinh Vien -=-=-=-=-=-")
	if len(List) == 0 {
		fmt.Println("📭 Danh sách sinh viên trống.")
	} else {
		// Hiển thị danh sách
		for _, s := range List {
			fmt.Printf("ID: %d, Tên: %s\n", s.ID, s.Name)
		}
	}
}
