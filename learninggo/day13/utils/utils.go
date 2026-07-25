package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ClearScreen xóa màn hình console (hoạt động trên Linux/macOS và Windows terminal hiện đại)
func ClearScreen() {
	// Đây là mã ANSI escape sequence để xóa màn hình.
	fmt.Print("\033[H\033[2J")
}

// ReadString đọc một chuỗi từ đầu vào chuẩn, cắt bỏ ký tự xuống dòng thừa
func ReadString(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// ReadInt đọc một số nguyên từ đầu vào chuẩn, xử lý lỗi nhập không phải số
func ReadInt(prompt string) int {
	for {
		s := ReadString(prompt)
		i, err := strconv.Atoi(s)
		if err == nil {
			return i
		}
		fmt.Println("❌ Giá trị không hợp lệ, vui lòng nhập số nguyên.")
	}
}
