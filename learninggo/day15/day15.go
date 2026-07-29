package main // 1. Phải để là package main mới chạy được

import "fmt"

// 1. PHẢI CÓ ĐOẠN NÀY: Định nghĩa cấu trúc nhân viên
type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	// 2. Khởi tạo Map chứa các đối tượng Employee
	// employees := map[string]Employee{
	// 	"e1": {Name: "Tuan", Age: 31, Role: "Developer"},
	// 	"e2": {Name: "Teo", Age: 35, Role: "Manager"},
	// }

	// // 3. Truy cập và in thông tin chi tiết
	// fmt.Println("--- Thong tin nhan vien e1 ---")

	// // Cách lấy đúng: employees["e1"].Name
	// fmt.Printf("Name: %s \n", employees["e1"].Name)
	// fmt.Printf("Age:  %d \n", employees["e1"].Age) // %d dành cho số nguyên
	// fmt.Printf("Role: %s \n", employees["e1"].Role)

	// fmt.Println("\n--- Danh sach tat ca nhan vien ---")
	// for id, emp := range employees {
	// 	fmt.Printf("ID: %s | Ten: %-10s | Chuc vu: %s\n", id, emp.Name, emp.Role)

	studentSubject := map[string][]string{
		"Tuan": {"Toan", "Khoa hoc"},
		"Teo":  {"CNTT", "khoa hoc"},
	}

	// fmt.Printf("%+v \n", studentSubject)

	// fmt.Printf("Mon hoc cua Tuan la: %s \n", studentSubject["Tuan"][0])
	// fmt.Printf("Mon toan cua Tuan la: %s \n ", studentSubject["Tuan"][1])

	for key, value := range studentSubject {
		for _, subject := range value {
			fmt.Printf("Mon hoc cua %s la: %s \n", key, subject)
		}
	}

}

// student := map[int]string{
// 	10: "Tuan",
// 	15: "Teo",
// 	18: "Ti",
// 	20: "Henry",
// }

// // 3. FIX LỖI TÊN BIẾN: Sửa 'ke' thành 'k' cho đồng bộ
// for k, v := range student {
// 	fmt.Printf("Student[%d] co gia tri la: %s\n", k, v)
// }

// fmt.Println("-----------------------")

// // --- Ví dụ 2: Khởi tạo bằng hàm make ---
// monan := make(map[string]int)
// monan["chao"] = 500
// monan["com"] = 300

// // 4. Kiểm tra sự tồn tại của một phần tử trong Map
// value, exists := monan["chao"]
// if exists {
// 	fmt.Printf("Gia cua mon chao la: %d\n", value)
// }

// // 5. Cách in nhanh toàn bộ Map để debug
// fmt.Printf("Toan bo map mon an: %+v\n", monan)
