package teacher

// Định nghĩa cấu trúc Giảng viên
type Teacher struct {
	Id         int
	Name       string
	Subject    string
	BaseSalary int
	Bonus      int
}

// ✅ ĐẶT HÀM KIỂM TRA TẠI ĐÂY (Và xóa ở file service đi)
func IsIdUniqueTeacher(id int, list []Teacher) bool {
	for _, gv := range list {
		if gv.Id == id {
			return false
		}
	}
	return true
}
