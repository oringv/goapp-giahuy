package main

import "fmt"

// func main() {
/** BAI 19 **/
// func pheptoan(numb1, numb2 int) int {
// 	if 	numb2 == 0 {
// 		numb2 = 5
// 	}

// 	tong := numb1 + numb2

// 	return tong
// }

// func main() {
// 	ketqua := pheptoan(10, 0)

// 	fmt.Println("Day la ket qua cua phep toan", ketqua)

// 	if ketqua >= 10 {
// 		fmt.Println("Ket qua lon hon hoac bang 10")
// 	}

// }

/** BAI 20 **/

/** --- ĐỊNH NGHĨA CÁC HÀM (Nằm ngoài hàm main) --- **/

// 1. Hàm tính toán cho bài 20 (Trả về 4 giá trị)
func pheptoan(numb1, numb2 int) (int, int, int, float32) {
	// Tránh lỗi chia cho 0
	if numb2 == 0 {
		numb2 = 1
	}

	tong := numb1 + numb2
	hieu := numb1 - numb2
	tich := numb1 * numb2
	thuong := float32(numb1) / float32(numb2)

	return tong, hieu, tich, thuong
}

// 2. Hàm đếm ngược (Đệ quy)
func countDown(number int) {
	fmt.Println(number)

	if number > 0 {
		countDown(number - 1)
	}
}

func main() {
	// 	cong, tru, nhan, chia := pheptoan(10, 0)

	// 	fmt.Println("Day la ket qua cua phep cong", cong)
	// 	fmt.Println("Day la ket qua cua phep tru", tru)
	// 	fmt.Println("Day la ket qua cua phep nhan", nhan)
	// 	fmt.Println("Day la ket qua cua phep chia", chia)

	// 	cong, tru, nhan, chia = pheptoan(20, 6)

	// 	fmt.Println("Day la ket qua cua phep cong", cong)
	// 	fmt.Println("Day la ket qua cua phep tru", tru)
	// 	fmt.Println("Day la ket qua cua phep nhan", nhan)
	// 	fmt.Println("Day la ket qua cua phep chia", chia)
	countDown(20)

}
