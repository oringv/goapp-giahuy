package main

//** Bài 31 **//
// import (
// 	"encoding/json" // Thêm thư viện để dùng JSON
// 	"fmt"
// 	"os" // Thêm thư viện để dùng os.Exit
// )

// type Giangvien struct {
// 	Name   string `json:"ho_ten"` // Thêm tag để JSON trông chuyên nghiệp hơn
// 	Gender int    `json:"gioi_tinh"`
// 	Phone  string `json:"so_dien_thoai"`
// }

// // Method hiển thị thông tin
// func (gv *Giangvien) hienThiThongTin() {
// 	fmt.Printf("Ho ten: %s \n", gv.Name)
// 	fmt.Printf("Gioi tinh: %d \n", gv.Gender)
// 	fmt.Printf("Dien thoai: %s \n", gv.Phone)
// }

// // Method xóa sạch dữ liệu
// func (gv *Giangvien) clear() {
// 	gv.Name = ""
// 	gv.Gender = 0
// 	gv.Phone = ""
// }

// func main() {
// 	// 1. Khởi tạo dữ liệu
// 	giahuy := Giangvien{
// 		Name:   "Pham Thanh Gia Huy",
// 		Gender: 1,
// 		Phone:  "0116448999",
// 	}

// 	fmt.Println("--- Thong tin goc ---")
// 	giahuy.hienThiThongTin()

// 	// 2. Chuyển đổi sang JSON
// 	// Sửa lỗi chính tả: Marsal -> Marshal và giahuuy -> giahuy
// 	output, err := json.MarshalIndent(giahuy, "", "  ") // Dùng MarshalIndent để in ra đẹp hơn
// 	if err != nil {
// 		fmt.Println("Loi JSON:", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("\n--- Ket qua sau khi chuyen sang JSON ---")
// 	fmt.Println(string(output))

// 	// 3. Test thử hàm clear
// 	fmt.Println("\n--- Goi ham Clear ---")
// 	giahuy.clear()
// 	giahuy.hienThiThongTin() // Sẽ in ra trắng trơn
// }

import "fmt"

// 1. Định nghĩa Struct
type Hinhchunhat struct {
	// LƯU Ý: Phải dùng dấu huyền ` ` cho Struct Tag
	Chieudai  float32 `desc:"Chieu dai hinh chu nhat"`
	Chieurong float32 `desc:"Chieu rong hinh chu nhat"`
}

// 2. Phương thức tính Chu vi
// Công thức: (dài + rộng) * 2
func (hcn *Hinhchunhat) chuvi() float32 {
	return (hcn.Chieudai + hcn.Chieurong) * 2
}

// 3. Phương thức tính Diện tích
// Công thức: dài * rộng
func (hcn *Hinhchunhat) dientich() float32 {
	return hcn.Chieudai * hcn.Chieurong
}

func main() {
	// Khởi tạo đối tượng
	// 	hinhchunhat := Hinhchunhat{
	// 		Chieudai:  20,
	// 		Chieurong: 5,
	// 	}

	// 	// In kết quả với định dạng 2 chữ số thập phân (%.2f)
	// 	fmt.Printf("Kich thuoc chu vi hinh chu nhat: %.2f \n", hinhchunhat.chuvi())
	// 	fmt.Printf("Kich thuoc dien tich hinh chu nhat: %.2f \n", hinhchunhat.dientich())
	// }

	var hinhchunhat Hinhchunhat

	// --- NHẬP CHIỀU RỘNG ---
	for {
		fmt.Print("Vui long nhap kich thuoc chieu rong: ")
		_, err := fmt.Scan(&hinhchunhat.Chieurong) // Dùng Scan cho ổn định

		// SỬA LOGIC: err == nil (không lỗi) VÀ giá trị > 0 thì mới thoát (break)
		if err == nil && hinhchunhat.Chieurong > 0 {
			break
		}

		// Nếu sai thì thông báo và bắt nhập lại
		fmt.Println("❌ Loi: Kich thuoc chieu rong phai la so và lon hon 0!")

		// Xóa bộ nhớ đệm đề phòng người dùng nhập chữ
		var discard string
		fmt.Scanln(&discard)
	}

	// --- NHẬP CHIỀU DÀI ---
	for {
		fmt.Print("Vui long nhap kich thuoc chieu dai: ")
		_, err := fmt.Scan(&hinhchunhat.Chieudai)

		// SỬA LOGIC: err == nil VÀ dài > 0 thì mới thoát
		if err == nil && hinhchunhat.Chieudai > 0 {
			break
		}

		fmt.Println("❌ Loi: Kich thuoc chieu dai phai la so và lon hon 0!")
		var discard string
		fmt.Scanln(&discard)
	}

	fmt.Println("\n--- KẾT QUẢ ---")
	fmt.Printf("Kich thuoc chu vi hinh chu nhat: %.2f \n", hinhchunhat.chuvi())
	fmt.Printf("Kich thuoc dien tich hinh chu nhat: %.2f \n", hinhchunhat.dientich())
}
