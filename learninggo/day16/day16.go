package main

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/library"
	"goapp-giahuy/learninggo/day16/utils"
)

func main() {
	lib := library.NewLibrary()

	for {
		utils.ClearScreen()
		fmt.Println("\n📚 CHƯƠNG TRÌNH QUẢN LÝ THƯ VIỆN")
		fmt.Println("1. Thêm sách")
		fmt.Println("2. Xem danh sách sách")
		fmt.Println("3. Thêm người mượn")
		fmt.Println("4. Xem danh sách người mượn")
		fmt.Println("5. Mượn sách")
		fmt.Println("6. Xem lịch sử mượn")
		fmt.Println("7. Trả sách")
		fmt.Println("8. Tìm kiếm sách")
		fmt.Println("9. Thoát")

		choice := utils.GetPositiveInt("👉 Chọn chức năng: ")

		switch choice {
		case 1:
			fmt.Println("-=-=-=-=- Thêm Sách -=-=-=-=-")
			if err := library.AddBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm sách: %v\n", err)
			}
		case 2:
			fmt.Println("-=-=-=-=- Xem Danh Sách Sách -=-=-=-=-")
			if err := library.ListBooks(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xem danh sách: %v\n", err)
			}
		case 3:
			fmt.Println("-=-=-=-=- Thêm Người Mượn -=-=-=-=-")
			if err := library.AddBorrower(lib); err != nil {
				fmt.Printf("❌ Lỗi khi thêm người mượn: %v\n", err)
			}
		case 4:
			fmt.Println("-=-=-=-=- Xem Danh Sách Người Mượn -=-=-=-=-")
			if err := library.ListBorrowers(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xem danh sách: %v\n", err)
			}
		case 5:
			fmt.Println("-=-=-=-=- Mượn Sách -=-=-=-=-")
			if err := library.BorrowBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi mượn sách: %v\n", err)
			}
		case 6:
			fmt.Println("-=-=-=-= Xem Lịch Sử Mượn Sách -=-=-=-=-")
			if err := library.ListBorrowHistory(lib); err != nil {
				fmt.Printf("❌ Lỗi khi xem lịch sử: %v\n", err)
			}
		case 7:
			fmt.Println("-=-=-=-=- Trả Sách -=-=-=-=-")
			if err := library.ReturnBook(lib); err != nil {
				fmt.Printf("❌ Lỗi khi trả sách: %v\n", err)
			}
		case 8:
			fmt.Println("-=-=-=-=- Tìm Kiếm Sách -=-=-=-=-")
			if err := library.SearchBooks(lib); err != nil {
				fmt.Printf("❌ Lỗi khi tìm kiếm: %v\n", err)
			}
		case 9:
			fmt.Println("👋 Tạm biệt!")
			return
		default:
			fmt.Println("⚠️ Lựa chọn không hợp lệ!")
		}
		utils.ReadInput("\nNhấn Enter để tiếp tục...")
	}
}
