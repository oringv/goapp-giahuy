package teacher

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/utils" // Nhập gói utils dùng chung
	"strings"                             // Cần thiết để tạo đường kẻ ngang
)

// GiangVien đại diện cho một giảng viên
type GiangVien struct {
	ID          int
	Ten         string
	MonGiangDay string
	LuongCoBan  int
	Thuong      int
}

// DanhSachGiangVien là danh sách giảng viên toàn cục
var DanhSachGiangVien []GiangVien

// MenuQuanLyGiangVien chứa vòng lặp cho menu quản lý giảng viên
func MenuQuanLyGiangVien() {
	for {
		utils.ClearScreen()
		fmt.Println("==== QUẢN LÝ GIẢNG VIÊN ====")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Quay lai")

		choice := utils.ReadInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			handleThemGiangVien()
		case 2:
			fmt.Println("Chức năng Xóa giảng viên chưa được triển khai.")
		case 3:
			fmt.Println("Chức năng Sửa giảng viên chưa được triển khai.")
		case 4:
			handleHienThiDanhSach()
		case 5:
			fmt.Println("Chức năng Tìm kiếm giảng viên chưa được triển khai.")
		case 6:
			return // Thoát khỏi menu
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
		utils.ReadString("\nNhan phim Enter de tiep tuc...")
	}
}

func handleThemGiangVien() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Them Giang Vien -=-=-=-=-=-")

	id := utils.ReadInt("- Nhap ID: ")
	ten := utils.ReadString("- Nhap ten: ")
	monGiangDay := utils.ReadString("- Nhap mon giang day: ")
	luongCoBan := utils.ReadInt("- Nhap luong co ban: ")
	thuong := utils.ReadInt("- Nhap thuong: ")

	gvMoi := GiangVien{
		ID:          id,
		Ten:         ten,
		MonGiangDay: monGiangDay,
		LuongCoBan:  luongCoBan,
		Thuong:      thuong,
	}

	DanhSachGiangVien = append(DanhSachGiangVien, gvMoi)
	fmt.Println("✅ Them giang vien thanh cong!")
}

func handleHienThiDanhSach() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Danh Sach Giang Vien -=-=-=-=-=-")
	if len(DanhSachGiangVien) == 0 {
		fmt.Println("📭 Danh sách giảng viên trống.")
	} else {
		// Định dạng hiển thị cho đẹp bằng Printf giống như trong hình
		headerFmt := "%-5s | %-20s | %-15s | %-12s | %-12s\n"
		rowFmt := "%-5d | %-20s | %-15s | %-12d | %-12d\n"

		fmt.Printf(headerFmt, "ID", "Ten", "Mon Giang Day", "Luong CB", "Thuong")
		fmt.Println(strings.Repeat("-", 75))

		for _, gv := range DanhSachGiangVien {
			fmt.Printf(rowFmt, gv.ID, gv.Ten, gv.MonGiangDay, gv.LuongCoBan, gv.Thuong)
		}
	}
}
