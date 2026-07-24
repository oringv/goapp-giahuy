package main

import "fmt"

func main() { // Sửa từ mainday12 thành main

	// 1. Khởi tạo slice rỗng
	var numbers []int
	fmt.Println(numbers) // Kết quả in ra: []

	// 2. Khởi tạo slice có giá trị
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // Kết quả in ra: [1 2 3 4 5]

	// 3. Khởi tạo array có giá trị
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array) // Kết quả in ra: [1 2 3 4 5]
}
