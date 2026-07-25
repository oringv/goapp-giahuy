package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Khởi tạo slice rỗng
	var numbers []int
	fmt.Println("Numbers (slice rỗng):", numbers)

	// 2. Khởi tạo slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// 3. Khởi tạo array
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array)

	// --- PHẦN KIỂM TRA BẢN CHẤT (SỬA LỖI Ở ĐÂY) ---
	// Dùng reflect.TypeOf(biến).Kind()
	fmt.Println("slice có phải là Slice không?", reflect.TypeOf(slice).Kind() == reflect.Slice)
	fmt.Println("array có phải là Array không?", reflect.TypeOf(array).Kind() == reflect.Array)
}
