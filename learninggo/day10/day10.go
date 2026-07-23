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

//** Bài 33 **//

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Định nghĩa Struct
type Hinhchunhat struct {
	Chieudai  float32
	Chieurong float32
}

// 2. Các Method tính toán
func (hcn *Hinhchunhat) chuvi() float32 {
	return (hcn.Chieudai + hcn.Chieurong) * 2
}

func (hcn *Hinhchunhat) dientich() float32 {
	return hcn.Chieudai * hcn.Chieurong
}

// 3. Hàm bổ trợ: Đọc dữ liệu từ bàn phím và chuyển sang số thực
func readFloat(prompt string) (float32, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Loại bỏ dấu xuống dòng và khoảng trắng

	value, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

// 4. Hàm xử lý nhập liệu cho Hình Chữ Nhật (Có vòng lặp kiểm tra)
func inputHinhChuNhat() Hinhchunhat {
	var hcn Hinhchunhat

	// Vòng lặp cho Chiều Rộng
	for {
		val, err := readFloat("Vui long nhap kich thuoc chieu rong hinh chu nhat: ")
		if err != nil || val <= 0 {
			fmt.Println("⛔ Kich thuoc chieu rong phai lon hon 0")
			continue
		}
		hcn.Chieurong = val
		break
	}

	// Vòng lặp cho Chiều Dài
	for {
		val, err := readFloat("Vui long nhap kich thuoc chieu dai hinh chu nhat: ")
		if err != nil || val <= 0 {
			fmt.Println("⛔ Kich thuoc chieu dai phai lon hon 0")
			continue
		}
		hcn.Chieudai = val
		break
	}

	return hcn
}

func main() {
	fmt.Println("=== CHUONG TRINH QUAN LY HINH CHU NHAT ===")

	// Gọi hàm nhập liệu
	hcn := inputHinhChuNhat()

	// Hiển thị kết quả
	fmt.Println("\n--- KET QUA ---")
	fmt.Printf("Kich thuoc chu vi hinh chu nhat: %.2f \n", hcn.chuvi())
	fmt.Printf("Kich thuoc dien tich hinh chu nhat: %.2f \n", hcn.dientich())
}
