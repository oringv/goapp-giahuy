package main

// import (
// 	"cmp"
// 	"fmt"

// 	// Lưu ý: Thư viện này thường phải chạy lệnh 'go get' mới có
// 	// Nếu không muốn cài thêm, bạn có thể comment dòng này lại
// 	"golang.org/x/exp/constraints"
// )

// // Khai báo Struct với Generics
// type Box[T any] struct {
// 	Content     T
// 	Description T
// }

// // 1. Hàm in bất kỳ kiểu gì
// func PrintValue[T any](v T) {
// 	fmt.Println(v)
// }

// // 2. So sánh bằng
// func IsEqual[T comparable](a, b T) bool {
// 	return a == b
// }

// func NotIsEqual[T comparable](a, b T) bool {
// 	return a != b
// }

// // 3. Định nghĩa Interface cho số
// type Number interface {
// 	constraints.Integer | constraints.Float
// }

// func Sum[T Number](a, b T) T {
// 	return a + b
// }

// // 4. Hàm tìm Max dùng thư viện cmp chuẩn của Go 1.21+
// func Max[T cmp.Ordered](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 5. Hàm so sánh độ dài chuỗi
// func MaxLengthString(a, b string) string {
// 	if len(a) > len(b) {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	// fmt.Println("--- Kết quả chạy bài 40 ---")

// 	// fmt.Print("Số lớn nhất giữa 10 và 7 là: ")
// 	// PrintValue(Max(10, 7))

// 	// fmt.Print("Số lớn nhất giữa 5.5 và 9 là: ")
// 	// PrintValue(Max(5.5, 9))

// 	// fmt.Print("Chuỗi lớn hơn (theo bảng chữ cái): ")
// 	// PrintValue(Max("Henry", "Henry Pham"))

// 	// fmt.Print("Tên có độ dài lớn hơn là: ")
// 	// PrintValue(MaxLengthString("Tung", "Vy"))

// 	// fmt.Println("\n--- Test Box với Generics ---")

// 	// stringBox := Box[string]{Content: "Hoc Golang Generic", Description: "Mo ta String Box"}
// 	// intBox := Box[int]{Content: 99, Description: 100}

// 	// fmt.Printf("String Box: %v - %v\n", stringBox.Content, stringBox.Description)
// 	// fmt.Printf("Int Box: %v - %v\n", intBox.Content, intBox.Description)

// 	PrintValue(Sum(5.5, 10))
// 	PrintValue(Sum(9, 10))
// 	PrintValue(Sum(3.6, 6.5))
// }

//** Bài 43 **//

import "fmt"

func main() {

	// var number int
	// fmt.Println(number)

	// var numbers [5]int
	// fmt.Println(numbers)

	// var character string
	// fmt.Println(character)

	// var character [3]string
	// fmt.Println(characters)

	var number [5]int
	number[2] = 10 // 0 0 10 0 0
	number[4] = 5  // 0 0 10 0 5
	fmt.Println(number)

	var numbers03 = [...]int{5, 8, 9, 10, 11}

	// 2. In kiểu dữ liệu (%T) để xem Go đếm đúng không
	fmt.Printf("Total array: %T \n", numbers03)

	// 3. In toàn bộ mảng
	fmt.Println(numbers03)
}
