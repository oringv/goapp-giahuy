package student

import (
	"fmt"
	"goapp-giahuy/learninggo/day14/utils" // Đảm bảo đường dẫn này đúng dự án của bạn
	"strings"
)

var List []Student

// Kiểm tra ID duy nhất
func IsIdUnique(id int) bool {
	for _, s := range List {
		if s.Id == id {
			return false
		}
	}
	return true
}

// ✅ SỬA GIAO DIỆN TẠI ĐÂY
func SubMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==== 🎓 QUẢN LÝ SINH VIÊN ====")
		fmt.Println("1. Thêm sinh viên mới")
		fmt.Println("2. Xóa sinh viên")
		fmt.Println("3. Sửa thông tin sinh viên")
		fmt.Println("4. Danh sách sinh viên")
		fmt.Println("5. Tìm kiếm sinh viên")
		fmt.Println("6. Quay lại Menu chính")

		choice := utils.GetPositiveInt("\n👉 Chọn chức năng: ")

		switch choice {
		case 1:
			handleThemSinhVien()
		case 2:
			handleXoa()
		case 3:
			handleSua()
		case 4:
			handleHienThiDanhSach()
		case 5:
			handleTimKiem()
		case 6:
			return // Thoát ra menu chính
		default:
			fmt.Println("⚠️ Lựa chọn không hợp lệ!")
		}
		utils.ReadInput("\nNhấn phím Enter để tiếp tục...")
	}
}

// Hàm thêm sinh viên
func handleThemSinhVien() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Thêm Sinh Viên -=-=-=-=-=-")
	var id int
	for {
		id = utils.GetPositiveInt("- Nhập ID: ")
		if IsIdUnique(id) {
			break
		}
		fmt.Println("❌ Id đã tồn tại! Vui lòng nhập Id khác.")
	}
	name := utils.ReadInput("- Nhập tên: ")
	class := utils.ReadInput("- Nhập lớp: ")

	numScores := utils.GetPositiveInt("- Nhập số lượng đầu điểm: ")
	var scores []float64
	for i := 1; i <= numScores; i++ {
		prompt := fmt.Sprintf("  + Nhập điểm số %d: ", i)
		val := utils.GetPositiveFloat(prompt)
		scores = append(scores, val)
	}

	List = append(List, Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	})
	fmt.Println("✅ Thêm sinh viên thành công!")
}

// Hàm hiển thị danh sách
func handleHienThiDanhSach() {
	utils.ClearScreen()
	fmt.Println("-=-=-=-=-=- Danh Sách Sinh Viên -=-=-=-=-=-")
	if len(List) == 0 {
		fmt.Println("📭 Danh sách đang trống.")
		return
	}
	fmt.Printf("%-5s | %-15s | %-8s | %-5s\n", "ID", "Họ Tên", "Lớp", "ĐTB")
	fmt.Println(strings.Repeat("-", 45))
	for _, s := range List {
		fmt.Printf("%-5d | %-15s | %-8s | %-5.2f\n", s.Id, s.Name, s.Class, s.GetAverage())
	}
}

// Hàm xóa
func handleXoa() {
	id := utils.GetPositiveInt("Nhập ID sinh viên cần xóa: ")
	for i, s := range List {
		if s.Id == id {
			List = append(List[:i], List[i+1:]...)
			fmt.Println("✅ Đã xóa sinh viên thành công!")
			return
		}
	}
	fmt.Println("❌ Không tìm thấy sinh viên có ID này.")
}

// Hàm sửa
func handleSua() {
	fmt.Println("\n-=-=-=-=- Sửa Thông Tin Sinh Viên -=-=-=-=-")
	id := utils.GetPositiveInt("Nhập ID sinh viên cần sửa: ")
	index := -1
	for i, s := range List {
		if s.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("❌ Lỗi: Không tìm thấy sinh viên!")
		return
	}
	List[index].Name = utils.GetOptionalString("- Nhập tên mới", List[index].Name)
	List[index].Class = utils.GetOptionalString("- Nhập lớp mới", List[index].Class)
	if len(List[index].Scores) > 0 {
		for i, oldScore := range List[index].Scores {
			prompt := fmt.Sprintf("  + Điểm số %d", i+1)
			List[index].Scores[i] = utils.GetOptionalPositiveFloat(prompt, oldScore)
		}
	}
	fmt.Println("✅ Cập nhật thông tin thành công!")
}

// Hàm tìm kiếm
func handleTimKiem() {
	name := utils.ReadInput("Nhập tên sinh viên cần tìm: ")
	fmt.Println("\n--- Kết quả tìm kiếm ---")
	found := false
	for _, s := range List {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(name)) {
			fmt.Printf("ID: %d | Tên: %s | Lớp: %s | ĐTB: %.2f\n", s.Id, s.Name, s.Class, s.GetAverage())
			found = true
		}
	}
	if !found {
		fmt.Println("Không tìm thấy kết quả nào.")
	}
}
