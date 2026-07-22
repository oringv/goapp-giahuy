package main

import "fmt"

// func LythuyetPointer() {

// 	name := "Pham Thanh Gia Huy"

// 	fmt.Println("-=-=-=-=-= Biến name thông thường -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", name)
// 	fmt.Printf("Value:     %v \n", name)
// 	fmt.Printf("Address:   %v \n", &name)

// 	fmt.Println()

// 	// ptrName là con trỏ bậc 1 (trỏ đến biến name)
// 	ptrName := &name
// 	fmt.Println("-=-=-=-=-= Con trỏ ptrName (Cấp 1) -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", ptrName)  // Kiểu *string
// 	fmt.Printf("Value:     %v \n", ptrName)  // Lưu địa chỉ của name
// 	fmt.Printf("Address:   %v \n", &ptrName) // Địa chỉ của chính nó

// 	fmt.Println()

// 	// ptrName2 là con trỏ bậc 2 (trỏ đến con trỏ ptrName)
// 	ptrName2 := &ptrName
// 	fmt.Println("-=-=-=-=-= Con trỏ ptrName2 (Cấp 2) -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", ptrName2)  // Kiểu **string
// 	fmt.Printf("Value:     %v \n", ptrName2)  // Lưu địa chỉ của ptrName
// 	fmt.Printf("Address:   %v \n", &ptrName2) // Địa chỉ của chính nó

// 	fmt.Println("\n---------------- Lấy giá trị ----------------")

// 	// Dùng 1 dấu * để lấy giá trị ptrName đang giữ (là địa chỉ của name)
// 	fmt.Printf("Lấy giá trị ptrName từ ptrName2 (*ptrName2): %v\n", *ptrName2)

// 	// Dùng 2 dấu ** để lấy giá trị cuối cùng (là chuỗi tên của bạn)
// 	fmt.Printf("Lấy giá trị name từ ptrName2 (**ptrName2): %v\n", **ptrName2)
// }

// func updateName(name *string) {
// 	fmt.Println("Show", name)
// 	*name = "Nguyen Van A"
// 	fmt.Println("-=-=-=-=-= Information name variable after update value -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", *name)

// }

// func main() {

// 	name := "Pham Thanh Gia Huy"
// 	fmt.Println("-=-=-=-=-= Information name variable -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", name)
// 	fmt.Printf("Value:     %v \n", name)
// 	fmt.Printf("Address:   %v \n", &name)

// 	fmt.Println()
// 	name = "Tony Teo"
// 	fmt.Println("-=-=-=-=-= Information name variable after update value -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", name)
// 	fmt.Printf("Value:     %v \n", name)
// 	fmt.Printf("Address:   %v \n", &name)

// 	fmt.Println()

// 	updateName(&name)

// 	fmt.Println("-=-=-=-=-= Information name variable after run updateName -=-=-=-=-=")
// 	fmt.Printf("Data type: %T \n", name)
// 	fmt.Printf("Value:     %v \n", name)
// 	fmt.Printf("Address:   %v \n", &name)

type Giangvien struct {
	name   string
	gender int
	phone  string
}

func (gv Giangvien) hienThiThongTin() {
	fmt.Printf("Ho ten: %s \n", gv.name)
	fmt.Printf("Gioi tinh: %d \n", gv.gender)
	fmt.Printf("Dien thoai: %s \n", gv.phone)

}

func (gv Giangvien) clear() {
	gv.name = ""
	gv.gender = 0
	gv.phone = ""
}

func main() {

	giahuy := Giangvien{
		name:   "Pham Thanh Gia Huy",
		gender: 1,
		phone:  "096579653",
	}

	giahuy.hienThiThongTin()

	fmt.Println()

	tonyteo := Giangvien{
		name:   "Tony Teo",
		gender: 1,
		phone:  "096579885",
	}

	tonyteo.hienThiThongTin()

	fmt.Println()

	giahuy.clear()

	fmt.Println()

	giahuy.hienThiThongTin()
}
