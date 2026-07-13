# Video 1-10

1. Xem video, thực hành theo và lưu vào day1.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day1.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 1
- todo 
- Golang Cơ Bản #01: Module & Package Là Gì?
1.Tìm hiểu module và package trong golang
1.	Module: Là nơi sẽ lưu trữ những package (ví dụ package util, Package Main, package database) nằm trong module
2.	Package: Là các thư mục chức năng bên trong Module (một module sẽ chứa 1 hoặc nhiều package trong 1 package sẽ chứa 1 hoặc nhiều file go).
3.	Package main + func main(): Là điểm khởi đầu bắt buộc để tạo ra một ứng dụng có thể chạy được.
4.	Compiler (Biên dịch): Gom tất cả các Package lại, nén thành 1 file duy nhất (.exe hoặc binary) để mang đi chạy ở bất cứ đâu mà không cần cài lại Go.

Quy trình chuẩn để làm việc với Go:
B1: go mod init ở thư mục gốc (làm 1 lần).
go mod init goapp-giahuy
B2: Viết code vào file .go (nhớ có package main và func main).
B3: Mở terminal vào đúng thư mục chứa file đó.
B4: go run tên_file.go.

Trong một module gồm nhiều Package ví dụ package main thì nhiều go gồm func main() để thực thi ra Compiler sau đó hướng đến 1 go để chạ y

## Video 2
- todo 
Tìm Hiểu Các Cách Compile Code Go 
Những cách để compile ứng dụng golang
Lệnh	        Ý nghĩa trực quan	    Sử dụng như thế nào 
go mod init	    Xây móng nhà	        Chỉ làm 1 lần lúc mới tạo dự án.
go run .	    Chạy thử	            Tìm hàm main trên toàn bộ file và chạy
go build .	    Chạy ứng dụng	        Chạy với file thực thi được tạo
go run day1.go  Chạy ứng dụng           Tìm hàm main trên 1 file cụ thể để chạy 

## Video 3
- todo 
Biến ở trong golang
Biến là một vùng nhớ để lưu trữ dữ liệu
Giá trị của biến có thể bị đè

Kiểu dữ liệu trong golang
bool: Đúng hoặc Sai (true, false).
string: Văn bản/Chữ (Để trong ngoặc kép: "Gia Huy").
int: Số nguyên (1, 100, -5).
float64: Số có dấu phẩy (3.14, 0.5).
rune: Dùng cho 1 ký tự duy nhất, nhất là chữ có dấu hoặc Emoji.
byte: Dùng cho dữ liệu máy tính (như file, hình ảnh).

## Video 4
- todo 
Cách Khai Báo Biến và Hằng Số Hiệu Quả Trong Go
Tên biến nên được viết theo định dạng camelCase.
Tên biến không được phép bắt đầu bằng số, không dấu, không khoảng trắng, không ký tự đặc biệt không trùng các từ khóa trong golang
Có thể khai báo nhiều biến cùng 1 lúc
Cách 1
var varialbleName type = value có thể không khai báo kiểu dữ liệu hoặc giá trị 

## Video 5
- todo 
Hiểu Rõ & Dùng Short Variable Declaration (:=) Hiệu Quả
Kiểu dữ liệu của biến được suy ra từ giá trị (có nghĩa là trình biên dịch quyết định Kiểu dữ liệu của biến dựa trên giá trị)
Luôn phải khai báo giá trị
Cách 2
varialbleName := value 

## Video 6
- todo 
Nhập Dữ Liệu Từ Bàn Phím Trong Go – fmt.Scanln & Ứng Dụng 

## Video 7
- todo 

## Video 8
- todo 

## Video 9
- todo 

## Video 10
- todo 