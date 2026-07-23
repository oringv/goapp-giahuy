package main

// import (
// 	"fmt"
// 	"goapp-giahuy/learninggo/day11/cat"
// 	"goapp-giahuy/learninggo/day11/dog"
// 	"goapp-giahuy/learninggo/day11/mouse"
// 	"goapp-giahuy/learninggo/day11/service"
// )

// // --- CÁC HÀM CŨ CỦA BẠN ---
// func MakeSound(a service.Animal) {
// 	fmt.Printf("Day la tieng cua %s \n", a.GetName())
// 	fmt.Println(a.Speak())
// }

// func MakeSoundPlus(p service.AnimalPlus) {
// 	fmt.Printf("Day la tieng cua %s \n", p.GetName())
// 	fmt.Println(p.Speak())
// 	fmt.Println(p.Eat())
// }

// // --- PHẦN MỚI: TYPE ASSERTION (Kiểm tra kiểu ép buộc) ---
// func CheckTypeAssertion(v interface{}) {
// 	s, ok := v.(string) // Ép kiểu v về string
// 	if ok {
// 		fmt.Println("✅ Day la chuoi (string):", s)
// 	} else {
// 		fmt.Println("❌ Loi: Vui long gui vao mot chuoi (Please send a string)")
// 	}
// }

// // --- PHẦN MỚI: TYPE SWITCH (Kiểm tra nhiều kiểu cùng lúc) ---
// func CheckTypeSwitch(v interface{}) {
// 	switch val := v.(type) {
// 	case string:
// 		fmt.Println("👉 Day la kieu STRING:", val)
// 	case int:
// 		fmt.Println("👉 Day la kieu INT:", val)
// 	case *dog.Dog:
// 		fmt.Println("👉 Day la kieu doi tuong DOG, ten la:", val.GetName())
// 	default:
// 		fmt.Printf("❓ Khong xac dinh duoc kieu %T \n", val)
// 	}
// }

// func main() {
// 	// 1. Chạy phần cũ (Dog, Cat, Mouse)
// 	mydog, _ := dog.New("Bully")
// 	MakeSound(mydog)

// 	fmt.Println("-=-=-=-=-=-=-=-=-=-")

// 	mycat, _ := cat.New("Pon")
// 	MakeSoundPlus(mycat)

// 	fmt.Println("-=-=-=-=-=-=-=-=-=-")

// 	mymouse, _ := mouse.New("Mr Ti")
// 	fmt.Println(mymouse.GetName())

// 	fmt.Println("\n======= TEST TYPE ASSERTION & SWITCH =======")

// 	// 2. Thử nghiệm Type Assertion
// 	CheckTypeAssertion("Xin chao") // Đúng kiểu
// 	CheckTypeAssertion(100)        // Sai kiểu (sẽ báo Please send a string)

// 	fmt.Println("-----------------")

// 	// 3. Thử nghiệm Type Switch
// 	CheckTypeSwitch("Gia Huy")
// 	CheckTypeSwitch(2024)
// 	CheckTypeSwitch(mydog) // Truyền hẳn con chó vào
// }

// import "fmt"

// //** Bài 39 **/
// func PrintValue[T any](value T) {
// 	fmt.Print(value)
// }

// func main() {

// 	PrintValue("Tuan")
// 	PrintValue(100)
// 	PrintValue(false)

// }

//** Bài 40**//
import (
	"cmp" // Thư viện hỗ trợ so sánh (từ Go 1.21+)
	"fmt"
)

// 1. Sử dụng 'comparable': Chỉ dùng được phép so sánh == và !=
// Hàm này kiểm tra hai giá trị có bằng nhau hay không
func IsEqual[T comparable](a, b T) bool {
	return a == b
}

// 2. Sử dụng 'cmp.Ordered': Dùng được các phép so sánh >, <, >=, <=
// Hàm này tìm số lớn nhất trong hai số (hoặc chuỗi)
func GetMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("--- Kết quả chạy chương trình ---")

	// Lần chạy 1: Với số nguyên (int)
	fmt.Println(GetMax(10, 9)) // In ra 10

	// Lần chạy 2: Với số thực (float64)
	fmt.Println(GetMax(10.0, 9.2)) // In ra 10

	// Lần chạy 3: Với chuỗi (string) - So sánh theo thứ tự bảng chữ cái
	fmt.Println(GetMax("Tuan", "An")) // In ra Tuan

	// Test thử comparable
	fmt.Println("10 có bằng 10 không?", IsEqual(10, 10))
	fmt.Println("Chuỗi có bằng nhau không?", IsEqual("Gia Huy", "Gia Huy"))

	// CHỈ CÓ 1 HÀM MAIN DUY NHẤT

	fmt.Println("--- Kết quả chạy chương trình ---")

	// In kết quả tìm số lớn nhất (tự động xuống hàng nhờ Println)
	fmt.Println(GetMax(10, 9))
	fmt.Println(GetMax(10.0, 9.2))
	fmt.Println(GetMax("Tuan", "An"))

	fmt.Println("\n--- Test thử so sánh bằng (comparable) ---")
	fmt.Println("10 có bằng 10 không?", IsEqual(10, 10))
	fmt.Println("Chuỗi có giống nhau không?", IsEqual("Gia Huy", "Gia Huy"))
}
