package library

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/models"
	"goapp-giahuy/learninggo/day16/utils"
	"strings"
)

// Chức năng Thêm Sách
func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNonEmptyString("- Nhập tiêu đề: ")
	author := utils.GetNonEmptyString("- Nhập tác giả: ")

	// Tạo đối tượng sách
	book := models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	// Gán vào map (Dùng dấu = thay vì := và đúng tên trường Books)
	lib.Books[id] = book

	fmt.Println("✅ Thêm sách thành công!")
	// In toàn bộ kho sách giống như ảnh Terminal bạn gửi
	fmt.Printf("Current Library Map: %+v \n", lib.Books)

	return nil
}

// Chức năng Xem Danh Sách
func ListBooks(lib *Library) error {
	// 2. Kiểm tra thư viện rỗng
	if len(lib.Books) == 0 {
		fmt.Println("📭 Thư viện hiện tại chưa có sách nào.")
		return nil
	}

	fmt.Println("\n--- DANH SÁCH SÁCH TRONG THƯ VIỆN ---")
	// Định dạng bảng
	fmt.Printf("%-36s | %-15s | %-15s | %-10s\n", "Mã ID (UUID)", "Tiêu Đề", "Tác Giả", "Trạng Thái")
	fmt.Println(strings.Repeat("-", 85))

	// 3. Duyệt Map bằng for range để in thông tin
	for _, b := range lib.Books {
		status := "🟢 Còn sách"
		if b.IsBorrowed {
			status = "🔴 Đã mượn"
		}

		fmt.Printf("%-36s | %-15s | %-15s | %-10s\n", b.Id, b.Title, b.Author, status)
	}
	return nil
}

// Các chức năng khác trả về nil để không bị lỗi build
func AddBorrower(lib *Library) error       { return nil }
func ListBorrowers(lib *Library) error     { return nil }
func BorrowBook(lib *Library) error        { return nil }
func ListBorrowHistory(lib *Library) error { return nil }
func ReturnBook(lib *Library) error        { return nil }
func SearchBooks(lib *Library) error       { return nil }
