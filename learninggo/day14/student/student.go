package student

import (
	"fmt"
	"goapp-giahuy/learninggo/day13/utils"
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

func SubMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==== 🎓 QUẢN LÝ SINH VIÊN ====")
		fmt.Println("1. Thêm mới | 2. Xóa | 3. Sửa thông tin | 4. Danh sách | 5. Tìm kiếm | 6. Quay lại")
		choice := utils.GetPositiveInt("👉 Chọn chức năng: ")

		switch choice {
		case 1:
			handleThemSinhVien()
		case 2:
			handleXoa()
		case 3:
			handleSua() // ✅ CHỨC NĂNG MỚI
		case 4:
			handleHienThiDanhSach()
		case 5:
			handleTimKiem()
		case 6:
			return
		}
		utils.ReadInput("\nNhấn Enter để tiếp tục...")
	}
}

// ✅ CHỨC NĂNG SỬA SINH VIÊN (Logic bài tập)
func handleSua() {
	fmt.Println("\n-=-=-=-=- Sửa Thông Tin Sinh Viên -=-=-=-=-")
	id := utils.GetPositiveInt("Nhập ID sinh viên cần sửa: ")

	// 1. Tìm sinh viên theo ID
	index := -1
	for i, s := range List {
		if s.Id == id {
			index = i
			break
		}
	}

	// 2. Kiểm tra nếu ID không tồn tại
	if index == -1 {
		fmt.Println("❌ Lỗi: Không tìm thấy sinh viên có ID này!")
		return
	}

	fmt.Println("🔍 Đã tìm thấy sinh viên. Nhấn Enter nếu không muốn thay đổi.")

	// 3. Cập nhật Tên và Lớp (Có tùy chọn giữ nguyên)
	List[index].Name = utils.GetOptionalString("- Nhập tên mới", List[index].Name)
	List[index].Class = utils.GetOptionalString("- Nhập lớp mới", List[index].Class)

	// 4. Cập nhật điểm (Duyệt qua slice scores)
	if len(List[index].Scores) > 0 {
		fmt.Println("- Cập nhật danh sách điểm:")
		for i, oldScore := range List[index].Scores {
			prompt := fmt.Sprintf("  + Điểm số %d", i+1)
			List[index].Scores[i] = utils.GetOptionalPositiveFloat(prompt, oldScore)
		}
	}

	fmt.Println("✅ Cập nhật thông tin thành công!")
}

// ... Giữ các hàm handleThemSinhVien, handleXoa, handleHienThiDanhSach, handleTimKiem như cũ ...

func handleThemSinhVien() {
	fmt.Println("\n-=-=-=-=- Thêm Mới -=-=-=-=-")
	var id int
	for {
		id = utils.GetPositiveInt("- ID: ")
		if IsIdUnique(id) {
			break
		}
		fmt.Println("❌ ID đã tồn tại!")
	}
	name := utils.ReadInput("- Tên: ")
	class := utils.ReadInput("- Lớp: ")
	n := utils.GetPositiveInt("- Số lượng điểm: ")
	var sc []float64
	for i := 1; i <= n; i++ {
		sc = append(sc, utils.GetPositiveFloat(fmt.Sprintf("  + Điểm %d: ", i)))
	}
	List = append(List, Student{Id: id, Name: name, Class: class, Scores: sc})
	fmt.Println("✅ Thành công!")
}

func handleHienThiDanhSach() {
	if len(List) == 0 {
		fmt.Println("📭 Trống.")
		return
	}
	fmt.Printf("%-5s | %-15s | %-8s | %-5s\n", "ID", "Họ Tên", "Lớp", "ĐTB")
	fmt.Println(strings.Repeat("-", 45))
	for _, s := range List {
		fmt.Printf("%-5d | %-15s | %-8s | %-5.2f\n", s.Id, s.Name, s.Class, s.GetAverage())
	}
}

func handleXoa() {
	id := utils.GetPositiveInt("ID cần xóa: ")
	for i, s := range List {
		if s.Id == id {
			List = append(List[:i], List[i+1:]...)
			fmt.Println("✅ Đã xóa!")
			return
		}
	}
	fmt.Println("❌ Không tìm thấy.")
}

func handleTimKiem() {
	name := utils.ReadInput("Tên cần tìm: ")
	for _, s := range List {
		if strings.Contains(strings.ToLower(s.Name), strings.ToLower(name)) {
			fmt.Printf("ID: %d | Tên: %s | Lớp: %s\n", s.Id, s.Name, s.Class)
		}
	}
}
