package library

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/models"
	"time"
)

type Library struct {
	Books     map[string]models.Book
	Borrowers map[string]models.Borrower
	Transaction map[string]models.Transaction
	Records []models.BorrowRecord
}

func NewLibrary() *Library {
	return &Library{
		Books:     make(map[string]models.Book),
		Borrowers: make(map[string]models.Borrower),
		Transaction: make(map[string]models.Transaction)
		Records: []models.BorrowRecord{},
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.Books[id]; exists {
		return fmt.Errorf("sach voi Id %s da ton tai", id)
	}

	lib.Books[id] = models.Book{
		Id:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false, // Mặc định là chưa mượn
	}
	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		books = append(books, book)
	}
	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.Borrowers[id]; exists {
		return fmt.Errorf("nguoi muon voi Id %s da ton tai", id)
	}
	lib.Borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}
	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	list := make([]models.Borrower, 0, len(lib.Borrowers))
	for _, b := range lib.Borrowers {
		list = append(list, b)
	}
	return list
}

// ✅ HÀM XỬ LÝ MƯỢN SÁCH (Đã hết lỗi đỏ)
func (lib *Library) BorrowBookStore(bookId, borrowerId string) error {
	// 1. Kiểm tra sách
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("không tìm thấy mã sách này!")
	}

	// 2. Kiểm tra trạng thái
	if book.IsBorrowed {
		return fmt.Errorf("sách này đang được người khác mượn!")
	}

	// 3. Kiểm tra người mượn
	if _, ok := lib.Borrowers[borrowerId]; !ok {
		return fmt.Errorf("không tìm thấy mã người mượn!")
	}

	// 4. Cập nhật trạng thái sách trong Map
	book.IsBorrowed = true
	lib.Books[bookId] = book

	// 5. Lưu vào lịch sử (Đã có trường Records ở trên nên sẽ chạy được)
	lib.Records = append(lib.Records, models.BorrowRecord{
		BookId:     bookId,
		BorrowerId: borrowerId,
		BorrowDate: time.Now(),
	})

	return nil
}
