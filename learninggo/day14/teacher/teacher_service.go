package teacher

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/utils"
	"strings"
)

var DanhSachGiangVien []Teacher // Lưu ý: Dùng đúng kiểu dữ liệu Teacher

func MenuQuanLyGiangVien() {
	for {
		utils.ClearScreen()
		fmt.Println("==== 👨‍🏫 QUẢN LÝ GIẢNG VIÊN ====")
		fmt.Println("1. Thêm giảng viên")
		fmt.Println("2. Xóa giảng viên")
		fmt.Println("3. Sửa giảng viên")
		fmt.Println("4. Danh sách giảng viên")
		fmt.Println("5. Tìm kiếm giảng viên")
		fmt.Println("6. Quay lại Menu chính")

		choice := utils.GetPositiveInt("👉 Chọn chức năng: ")

		switch choice {
		case 1:
			handleThemGiangVien()
		case 2:
			handleXoaGiangVien()
		case 3:
			handleSuaGiangVien()
		case 4:
			handleHienThiDanhSach()
		case 5:
			handleTimKiemGiangVien()
		case 6:
			return
		}
		utils.ReadInput("\nNhấn Enter để tiếp tục...")
	}
}

func handleThemGiangVien() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Thêm Giảng Viên -=-=-=-=-=-")

	var id int
	for {
		id = utils.GetPositiveInt("- Nhập ID: ")
		if IsIdUniqueTeacher(id, DanhSachGiangVien) {
			break
		}
		fmt.Println("❌ ID đã tồn tại! Vui lòng nhập ID khác.")
	}

	// ✅ ĐÃ FIX: Toàn bộ phần này phải nằm TRONG hàm handleThemGiangVien
	name := utils.ReadInput("- Nhập tên: ")
	sub := utils.ReadInput("- Nhập môn giảng dạy: ")
	luong := utils.GetPositiveInt("- Nhập lương cơ bản: ")
	thuong := utils.GetPositiveInt("- Nhập thưởng: ")

	DanhSachGiangVien = append(DanhSachGiangVien, Teacher{
		Id: id, Name: name, Subject: sub, BaseSalary: luong, Bonus: thuong,
	})
	fmt.Println("✅ Thêm giảng viên thành công!")
}

func handleHienThiDanhSach() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Danh Sách Giảng Viên -=-=-=-=-=-")
	if len(DanhSachGiangVien) == 0 {
		fmt.Println("📭 Danh sách đang trống.")
		return
	}
	fmt.Printf("%-5s | %-20s | %-15s | %-10s\n", "ID", "Tên", "Môn Dạy", "Lương")
	fmt.Println(strings.Repeat("-", 60))
	for _, gv := range DanhSachGiangVien {
		fmt.Printf("%-5d | %-20s | %-15s | %-10d\n", gv.Id, gv.Name, gv.Subject, gv.BaseSalary+gv.Bonus)
	}
}

func handleXoaGiangVien() {
	id := utils.GetPositiveInt("Nhập ID cần xóa: ")
	for i, gv := range DanhSachGiangVien {
		if gv.Id == id {
			DanhSachGiangVien = append(DanhSachGiangVien[:i], DanhSachGiangVien[i+1:]...)
			fmt.Println("✅ Đã xóa!")
			return
		}
	}
	fmt.Println("❌ Không tìm thấy ID.")
}

func handleSuaGiangVien() {
	id := utils.GetPositiveInt("Nhập ID cần sửa: ")
	for i := range DanhSachGiangVien {
		if DanhSachGiangVien[i].Id == id {
			DanhSachGiangVien[i].Name = utils.ReadInput("Tên mới: ")
			DanhSachGiangVien[i].Subject = utils.ReadInput("Môn mới: ")
			fmt.Println("✅ Đã cập nhật!")
			return
		}
	}
}

func handleTimKiemGiangVien() {
	ten := utils.ReadInput("Nhập tên cần tìm: ")
	for _, gv := range DanhSachGiangVien {
		if strings.Contains(strings.ToLower(gv.Name), strings.ToLower(ten)) {
			fmt.Printf("ID: %d | Tên: %s | Môn: %s\n", gv.Id, gv.Name, gv.Subject)
		}
	}
}
