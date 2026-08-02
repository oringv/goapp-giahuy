package library

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/utils"
	"strings"
)

// 1. Chức năng Thêm Sách
func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNonEmptyString("- Nhập tiêu đề sách: ")
	author := utils.GetNonEmptyString("- Nhập tác giả: ")

	return lib.AddBookStore(id, title, author)
}

// 2. Chức năng Xem Danh Sách Sách
func ListBooks(lib *Library) error {
	if len(lib.Books) == 0 {
		fmt.Println("📭 Thư viện hiện tại chưa có sách nào.")
		return nil
	}
	fmt.Println("\n--- DANH SÁCH SÁCH TRONG THƯ VIỆN ---")
	fmt.Printf("%-36s | %-20s | %-15s | %-10s\n", "Mã ID", "Tiêu Đề", "Tác Giả", "Trạng Thái")
	fmt.Println(strings.Repeat("-", 90))
	for _, b := range lib.Books {
		status := "🟢 Còn sách"
		if b.IsBorrowed {
			status = "🔴 Đã mượn"
		}
		fmt.Printf("%-36s | %-20s | %-15s | %-10s\n", b.Id, b.Title, b.Author, status)
	}
	return nil
}

// 3. Chức năng Thêm Người Mượn
func AddBorrower(lib *Library) error {
	id := utils.GenerateId()
	name := utils.GetNonEmptyString("- Nhập tên người mượn: ")
	email := utils.GetNonEmptyString("- Nhập email: ")
	return lib.AddBorrowerStore(id, name, email)
}

// 4. Chức năng Xem Danh Sách Người Mượn
func ListBorrowers(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("📭 Danh sách người mượn trống.")
		return nil
	}
	fmt.Println("\n--- DANH SÁCH NGƯỜI MƯỢN ---")
	for _, b := range borrowers {
		fmt.Printf("ID: %s | Tên: %-15s | Email: %s\n", b.Id, b.Name, b.Email)
	}
	return nil
}

// 5. Chức năng Mượn Sách (In thông báo ✅ đẹp như ảnh)
func BorrowBook(lib *Library) error {
	fmt.Println("-=-=-=-=- Mượn Sách -=-=-=-=-")
	bId := utils.ReadInput("Nhập ID Sách: ")
	uId := utils.ReadInput("Nhập ID Người mượn: ")
	tId := utils.GenerateId()

	err := lib.BorrowBookStore(tId, bId, uId)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Mượn sách thành công! ID giao dịch: %s\n", tId)
	return nil
}

// 6. Chức năng Xem Lịch Sử Giao Dịch
func ListBorrowHistory(lib *Library) error {
	if len(lib.Transactions) == 0 {
		fmt.Println("📭 Chưa có lịch sử giao dịch nào.")
		return nil
	}
	fmt.Println("\n--- LỊCH SỬ GIAO DỊCH MƯỢN SÁCH ---")
	for _, t := range lib.Transactions {
		bookTitle := lib.Books[t.BookId].Title
		borrowerName := lib.Borrowers[t.BorrowerId].Name
		fmt.Printf("Giao dịch: %s | Sách: %s | Người mượn: %s | Ngày: %v\n",
			t.Id, bookTitle, borrowerName, t.BorrowDate.Format("02/01/2006 15:04"))
	}
	return nil
}

// 8. Chức năng Tìm Kiếm Sách
// ✅ CHỨC NĂNG 8: Tìm kiếm sách đầy đủ
func SearchBooks(lib *Library) error {
	fmt.Println("-=-=-=-=- Tim Kiem Sach -=-=-=-=-")

	// 1. Nhập từ khóa tìm kiếm
	keyword := utils.ReadInput("Nhap tieu de hoac tac gia de tim kiem: ")

	// 2. Gọi xuống kho để lấy danh sách kết quả
	results := lib.SearchBooksStore(keyword)

	// 3. Kiểm tra nếu không tìm thấy gì
	if len(results) == 0 {
		fmt.Printf("🔍 Khong tim thay ket qua nao phu hop voi tu khoa: '%s'\n", keyword)
		return nil
	}

	// 4. In kết quả ĐẦY ĐỦ THÔNG TIN giống hệt ảnh mẫu
	for _, b := range results {
		status := "Con"
		if b.IsBorrowed {
			status = "Da muon"
		}

		fmt.Printf("Id: %s, Tieu de: %s, Tac gia: %s, Trang thai: %s\n",
			b.Id, b.Title, b.Author, status)
	}

	return nil
}

func ReturnBook(lib *Library) error {
	fmt.Println("-=-=-=-=- Trả Sách -=-=-=-=-")

	// ✅ SỬA DÒNG NÀY: Bảo người dùng nhập ID Sách
	id := utils.ReadInput("Nhập ID Sách muốn trả: ")

	err := lib.ReturnBookStore(id)
	if err != nil {
		fmt.Printf("❌ Lỗi khi trả sách: %v\n", err)
		return nil
	}

	fmt.Println("✅ Đã trả sách thành công! Trạng thái đã cập nhật thành CÒN SÁCH.")
	return nil
}
