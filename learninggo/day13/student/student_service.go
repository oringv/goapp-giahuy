package student

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/utils"
	"strings"
)

var List []Student

// ✅ THÊM HÀM NÀY: Kiểm tra xem ID đã tồn tại chưa
func IsIdUnique(id int) bool {
	for _, s := range List {
		if s.Id == id {
			return false // ID đã có người dùng
		}
	}
	return true // ID chưa có ai dùng, hợp lệ
}

func SubMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==== QUAN LY SINH VIEN ====")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPositiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			handleThemSinhVien()
		case 2:
			handleXoa()
		case 3:
			fmt.Println("Chức năng Sửa đang phát triển...")
		case 4:
			handleHienThiDanhSach()
		case 5:
			handleTimKiem()
		case 6:
			return
		default:
			fmt.Println("⚠️ Lựa chọn không hợp lệ!")
		}
		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}

func handleThemSinhVien() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Them Sinh Vien -=-=-=-=-=-")

	var id int
	// LOGIC KIỂM TRA ID TRÙNG (Vòng lặp cho đến khi nhập ID mới)
	for {
		id = utils.GetPositiveInt("- Nhap id: ")
		if IsIdUnique(id) {
			break // ID hợp lệ thì thoát vòng lặp để nhập tiếp tên
		}
		fmt.Println("❌ Id da ton tai! Vui long nhap Id khac.")
	}

	name := utils.ReadInput("- Nhap ten: ")
	class := utils.ReadInput("- Nhap lop: ")

	// Nhập nhiều điểm
	numScores := utils.GetPositiveInt("- Nhap so luong diem: ")
	var scores []float64
	for i := 1; i <= numScores; i++ {
		prompt := fmt.Sprintf("- Nhap diem %d: ", i)
		val := utils.GetPositiveFloat(prompt)
		scores = append(scores, val)
	}

	List = append(List, Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	})

	fmt.Println("✅ Them sinh vien thanh cong!")
}

func handleHienThiDanhSach() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Danh Sach Sinh Vien -=-=-=-=-=-")
	if len(List) == 0 {
		fmt.Println("📭 Danh sách trống.")
		return
	}

	fmt.Printf("%-5s | %-15s | %-10s | %-7s\n", "ID", "Họ Tên", "Lớp", "ĐTB")
	fmt.Println(strings.Repeat("-", 45))
	for _, s := range List {
		fmt.Printf("%-5d | %-15s | %-10s | %-7.2f\n", s.Id, s.Name, s.Class, s.GetAverage())
	}
}

func handleXoa() {
	id := utils.GetPositiveInt("Nhập ID cần xóa: ")
	for i, s := range List {
		if s.Id == id {
			List = append(List[:i], List[i+1:]...)
			fmt.Println("✅ Đã xóa thành công!")
			return
		}
	}
	fmt.Println("❌ Không tìm thấy sinh viên.")
}

func handleTimKiem() {
	name := utils.ReadInput("Nhập tên cần tìm: ")
	fmt.Println("\n--- Kết quả tìm kiếm ---")
	for _, s := range List {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(name)) {
			fmt.Printf("ID: %d | Tên: %s | Lớp: %s\n", s.Id, s.Name, s.Class)
		}
	}
}
