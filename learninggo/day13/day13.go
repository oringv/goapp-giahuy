package main

import (
	"fmt"
	"reflect"
)

// "reflect"

func main() {

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
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice cha")
	fmt.Println(slice)
	fmt.Println("slice co phai la Slice khong ? ", reflect.TypeOf(slice).Kind() == reflect.Slice)

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Slice con")
	subSlice := slice[1:4]
	fmt.Println(subSlice)
	fmt.Println("subSlice co phai la Slice khong ? ", reflect.TypeOf(subSlice).Kind() == reflect.Slice)

}
