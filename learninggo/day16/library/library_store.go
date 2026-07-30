package library

import "goapp-giahuy/learninggo/day16/models"

// 1. Định nghĩa Struct duy nhất cho toàn bộ package
type Library struct {
	Books map[string]models.Book
}

// 2. Hàm khởi tạo Library mới
func NewLibrary() *Library {
	return &Library{
		// Cấp phát bộ nhớ cho Map để tránh lỗi panic
		Books: make(map[string]models.Book),
	}
}
