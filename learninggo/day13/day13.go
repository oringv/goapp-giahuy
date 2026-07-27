 package main

// "fmt"
// "slices"
// "reflect"

// "reflect"

// func main() {

// arr := [5]int{10, 20, 30, 40, 50}
// fmt.Println("arr co phai la Array khong ?", reflect.TypeOf(arr).Kind() == reflect.Array)
// fmt.Println(arr)

// slice := arr[1:4]
// fmt.Println(slice)
// fmt.Println("slice co phai la Slice khong ? ", reflect.TypeOf(slice).Kind() == reflect.Slice)

// slice := make([]int, 2, 5)
// slice[0] = 1
// slice[1] = 2
// slice = append(slice, 3, 4, 5, 6, 7, 8, 9, 10)

// fmt.Println(slice)
// fmt.Println("Chieu dai cua slice: ", len(slice))
// fmt.Println("Dung luong toi da cua slice: ", cap(slice))

// slice := []int{1, 2, 3, 4, 5}
// fmt.Println(slice)
// fmt.Println("Chieu dai (length) cua slice:", len(slice))
// fmt.Println("Dung luong (capacity) toi da cua slice:", cap(slice))

// apple := []string{"Apple 1", "Apple 2", "Apple 3", "Apple 4", "Apple 5"}
// fmt.Println(apple[0])
// fmt.Println(apple[1])
// fmt.Println(apple[2])
// fmt.Println(apple[3])
// fmt.Println(apple[4])

// for i := 0; i < len(apple); i++ {
// 	fmt.Println(apple[i])
// }

// for key, value := range apple {
// 	fmt.Printf("apple[%d] = %s \n", key, value)
// }

// school := [][]string{
// 	{"Tuan", "Binh", "Dai"},
// 	{"Duy", "Hoa", "Huynh"},
// }

// for _, class := range school {
// 	for _, student := range class {

// 		fmt.Println(student)
// 	}
// }

// apple := []string{"Apple1", "Apple2", "Apple3", "Apple4", "Apple5"}
// apple = append(apple, "Apple 6", "Apple7")
// fmt.Println(apple)

// apple1 := []string{"Apple 1", "Apple 2", "Apple 3"}
// apple2 := []string{"Apple 4", "Apple 5", "Apple 6"}
// apple3 := []string{"Apple 7", "Apple 8", "Apple 9"}
// apple1 = append(apple1, apple2...)
// apple1 = append(apple1, apple3...)
// fmt.Println(apple1)

//** Bài 50 **//
// slice := []int{10, 20, 30, 40, 50}
// fmt.Println("Slice cha")
// fmt.Println(slice)
// fmt.Println("Length slice:", len(slice))
// fmt.Println("Capacity slice:", cap(slice))

// fmt.Println("-=-=-=-=-=-=-=-=-=-=-")
// fmt.Println("Slice con")
// subSlice := slice[2:4]
// subSlice = append(subSlice, 90)
// fmt.Println(subSlice)
// fmt.Println("Length slice:", len(subSlice))
// fmt.Println("Capacity slice:", cap(subSlice))

//** Bài 51 **//

// func main() {
// 	fmt.Println("-=-=-=-=- TÌM HIỂU HÀM SLICES -=-=-=-=-")

// 	/** 1. Clone: Tạo bản sao của slice **/
// 	copied := slices.Clone([]int{1, 2, 3})
// 	fmt.Println("Clone result:", copied)

// 	/** 2. Equal: So sánh 2 slice có giống hệt nhau không **/
// 	compareSlice := slices.Equal([]int{1, 2}, []int{1, 2})
// 	fmt.Println("Equal result:", compareSlice) // Trả về true/false

// 	/** 3. Index: Tìm vị trí đầu tiên của phần tử trong mảng **/
// 	// Tìm số 2 trong dãy {1, 2, 3, 4}
// 	findFirstPosition := slices.Index([]int{1, 2, 3, 4}, 2)
// 	fmt.Println("Index of 2 is:", findFirstPosition) // Kết quả: 1 (vị trí thứ 2)

// 	/** 4. Contains: Kiểm tra phần tử có tồn tại trong slice không **/
// 	itemExist := slices.Contains([]int{1, 2, 3}, 3)
// 	fmt.Println("Does 3 exist?", itemExist) // Trả về true

// 	/** 5. Insert: Chèn phần tử vào vị trí i bất kỳ **/
// 	// Chèn số 2 vào vị trí index 1 của slice {1, 3}
// 	insertValueAnyPostion := slices.Insert([]int{1, 3}, 1, 2)
// 	fmt.Println("After Insert:", insertValueAnyPostion) // Kết quả: [1 2 3]

// 	/** 6. Delete: Xóa phần tử từ vị trí i đến j-1 **/
// 	// Xóa từ index 1 đến trước index 3 (xóa phần tử tại vị trí 1 và 2)
// 	deleteValueAnyPostion := slices.Delete([]int{1, 2, 3, 4}, 1, 3)
// 	fmt.Println("After Delete:", deleteValueAnyPostion) // Kết quả: [1 4]

// 	/** 7. Reverse: Đảo ngược slice **/
// 	s := []int{1, 2, 3}
// 	slices.Reverse(s)                // Lưu ý: Hàm này thay đổi trực tiếp trên biến s
// 	fmt.Println("After Reverse:", s) // Kết quả: [3 2 1]

// 	/** 8. Sort: Sắp xếp slice tăng dần **/
// 	sSort := []int{3, 1, 2}
// 	slices.Sort(sSort)
// 	fmt.Println("After Sort:", sSort) // Kết quả: [1 2 3]

// 	/** 9. SortFunc: Sắp xếp theo điều kiện tùy chỉnh **/
// 	// Ví dụ sắp xếp giảm dần (b - a)
// 	sFunc := []int{3, 1, 2}
// 	slices.SortFunc(sFunc, func(a, b int) int {
// 		return b - a
// 	})
// 	fmt.Println("After SortFunc (desc):", sFunc) // Kết quả: [3 2 1]

}