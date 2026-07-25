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

// import "fmt"

// func main() {

// 	// var number int
// 	// fmt.Println(number)

// 	// var numbers [5]int
// 	// fmt.Println(numbers)

// 	// var character string
// 	// fmt.Println(character)

// 	// var character [3]string
// 	// fmt.Println(characters)

// 	var number [5]int
// 	number[2] = 10 // 0 0 10 0 0
// 	number[4] = 5  // 0 0 10 0 5
// 	fmt.Println(number)

// 	var numbers03 = [...]int{5, 8, 9, 10, 11}

// 	// 2. In kiểu dữ liệu (%T) để xem Go đếm đúng không
// 	fmt.Printf("Total array: %T \n", numbers03)

// 	// 3. In toàn bộ mảng
// 	fmt.Println(numbers03)
// }

//** Bài 44 **//
// import "fmt"

// func main() {
// 	var matrix = [2][3]int{

// 		{1, 2, 3},
// 		{4, 5, 6},
// 	}

// 	fmt.Println(matrix)

// 	matrix[1][1] = 9
// 	matrix[0][2] = 10
// 	fmt.Println(matrix)
// }

//** Bài 45 **//
// import "fmt"

// func main() {

// numbers := [5]int{6, 7, 8, 9, 10}
// fmt.Println(numbers[0])
// fmt.Println(numbers[1])
// fmt.Println(numbers[2])
// fmt.Println(numbers[3])
// fmt.Println(numbers[4])

// fmt.Println(len(numbers))
// for i := 0; i < len(numbers); i++ {
// 	fmt.Println(numbers[i])
// }

// numbers := [3][4]int{
// 	{1, 2, 3, 4},
// 	{5, 6, 7, 8},
// 	{9, 10, 11, 12},
// }

// for k := 0; k < len(numbers); k++ {
// 	for j := 0; j < len(numbers[k]); j++ {
// 		fmt.Println(numbers[k][j])
// 	}
// }

// numbers := [5] int {6,7,8,9,10}
// for key, val := range numbers {
// 	fmt.Printf("Array numbers [%d] = %d \n", key, val)
// }

// 	numbers := [3][4]int{
// 		{1, 2, 3, 4},
// 		{5, 6, 7, 8},
// 		{9, 10, 11, 12},
// 	}
// 	for key1, val := range numbers {
// 		for key2, val2 := range val {
// 			fmt.Printf("Array numbers[%d][%d]= %d \n", key1, key2, val2)
// 		}
// 	}
// }

//** Bài 46 **//
// import (
// 	"fmt"
// )

// type Nhanvien struct {
// 	Id   int
// 	Name string
// 	Age  int
// }

// func main() {
// 	// Khai báo mảng nhân viên
// 	employees := [...]Nhanvien{
// 		{Id: 1, Name: "Tuan", Age: 10},
// 		{Id: 2, Name: "An", Age: 15},
// 		{Id: 3, Name: "Binh", Age: 18},
// 		{Id: 4, Name: "Hoa", Age: 12},
// 	}

// 	// SỬ DỤNG VÒNG LẶP FOR RANGE ĐỂ IN GIỐNG TRONG ẢNH
// 	for _, nv := range employees {
// 		// %s cho chuỗi (Name), %d cho số nguyên (Age)
// 		fmt.Printf("Name: %s and Age: %d \n", nv.Name, nv.Age)
// 	}
// }
