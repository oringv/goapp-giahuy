package library

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/models"
	"strings"
	"time"
)

type Library struct {
	Books        map[string]models.Book
	Borrowers    map[string]models.Borrower
	Transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		Books:        make(map[string]models.Book),
		Borrowers:    make(map[string]models.Borrower),
		Transactions: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.Books[id]; exists {
		return fmt.Errorf("sách với Id %s đã tồn tại", id)
	}
	lib.Books[id] = models.Book{Id: id, Title: title, Author: author, IsBorrowed: false}
	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	var books []models.Book
	for _, b := range lib.Books {
		books = append(books, b)
	}
	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.Borrowers[id]; exists {
		return fmt.Errorf("người mượn với Id %s đã tồn tại", id)
	}
	lib.Borrowers[id] = models.Borrower{Id: id, Name: name, Email: email}
	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	var list []models.Borrower
	for _, b := range lib.Borrowers {
		list = append(list, b)
	}
	return list
}

func (lib *Library) BorrowBookStore(transId, bookId, borrowerId string) error {
	book, ok := lib.Books[bookId]
	if !ok {
		return fmt.Errorf("không tìm thấy mã sách!")
	}
	if book.IsBorrowed {
		return fmt.Errorf("sách này đang được mượn!")
	}
	if _, ok := lib.Borrowers[borrowerId]; !ok {
		return fmt.Errorf("không tìm thấy người mượn!")
	}

	book.IsBorrowed = true
	lib.Books[bookId] = book

	lib.Transactions[transId] = models.Transaction{
		Id:         transId,
		BookId:     bookId,
		BorrowerId: borrowerId,
		BorrowDate: time.Now(),
	}
	return nil
}

func (lib *Library) SearchBooksStore(keyword string) []models.Book {
	var results []models.Book

	// Chuyển từ khóa về chữ thường để tìm kiếm chính xác hơn
	keyword = strings.ToLower(strings.TrimSpace(keyword))

	for _, book := range lib.Books {
		// Kiểm tra nếu tiêu đề HOẶC tác giả chứa từ khóa
		if strings.Contains(strings.ToLower(book.Title), keyword) ||
			strings.Contains(strings.ToLower(book.Author), keyword) {
			results = append(results, book)
		}
	}
	return results
}

// Chức năng: Trả sách bằng cách nhập mã ID của chính cuốn sách đó
func (lib *Library) ReturnBookStore(bookId string) error {
	// 1. Kiểm tra xem mã sách có tồn tại không
	book, exists := lib.Books[bookId]
	if !exists {
		return fmt.Errorf("sách với ID %s không tồn tại", bookId)
	}

	// 2. Kiểm tra xem sách có thực sự đang bị mượn không
	if !book.IsBorrowed {
		return fmt.Errorf("sách '%s' hiện đang ở trong kho, không cần trả", book.Title)
	}

	// 3. Cập nhật trạng thái sách thành CÒN SÁCH
	book.IsBorrowed = false
	lib.Books[bookId] = book

	// 4. Tìm giao dịch tương ứng của cuốn sách này để ghi nhận ngày trả
	foundTrans := false
	for id, trans := range lib.Transactions {
		// Nếu đúng mã sách và giao dịch này chưa được trả (ReturnDate trống)
		if trans.BookId == bookId && trans.ReturnDate.IsZero() {
			trans.ReturnDate = time.Now()
			lib.Transactions[id] = trans // Lưu lại thay đổi
			foundTrans = true
			break
		}
	}

	if !foundTrans {
		return fmt.Errorf("không tìm thấy lịch sử mượn cho cuốn sách này")
	}

	return nil
}
