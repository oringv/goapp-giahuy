package main // 1. Phải để là package main mới chạy được

import "fmt" // 2. Thêm import fmt

func main() {
	// --- Ví dụ 1: Map Student (Key: int, Value: string) ---
	student := map[int]string{
		10: "Tuan",
		15: "Teo",
		18: "Ti",
		20: "Henry",
	}

	// 3. FIX LỖI TÊN BIẾN: Sửa 'ke' thành 'k' cho đồng bộ
	for k, v := range student {
		fmt.Printf("Student[%d] co gia tri la: %s\n", k, v)
	}

	fmt.Println("-----------------------")

	// --- Ví dụ 2: Khởi tạo bằng hàm make ---
	monan := make(map[string]int)
	monan["chao"] = 500
	monan["com"] = 300

	// 4. Kiểm tra sự tồn tại của một phần tử trong Map
	value, exists := monan["chao"]
	if exists {
		fmt.Printf("Gia cua mon chao la: %d\n", value)
	}

	// 5. Cách in nhanh toàn bộ Map để debug
	fmt.Printf("Toan bo map mon an: %+v\n", monan)
}
