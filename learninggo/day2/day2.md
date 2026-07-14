# Video 11-20

1. Xem video, thực hành theo và lưu vào day2.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day2.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 


## Video 6
- todo 
Nhập Dữ Liệu Từ Bàn Phím Trong Go – fmt.Scanln & Ứng Dụng 
1. Nhập số hoặc từ đơn
var age int
fmt.Scanln(&age) 
2. Nhập cả câu (có khoảng trắng)
reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')

## Video 7
- todo 
Hằng số trong Golang 
Giống như biến, hằng số giúp lưu trữ dữ liệu
Giá trị của hằng số KHÔNG được phép ghi đè trong cùng 1 phạm vi
Tên hằng số được viết hoa(Ví dụ HOTEN). Điều này sẽ giúp phân biệt HẰNG và BIẾN
Có thể khai báo trong hàm hoặc ngoài hàm
Khai báo hằng số
const CONSTNAME type = value 

## Video 8
- todo 
Tìm hiểu Package fmt
Package fmt trong Go được xây dựng để nhằm mục đích tương tác trực tiếp với dữ liệu như hiển thị dữ liệu, lấy dữ liệu từ bàn phím. Một số hàm trong package fmt hay sử dụng:
1. Print, Println, Printf: Hiển thị dữ liệu trực tiếp ra màn hình 
2. Scan, Scanln, Scanf: Nhập dữ liệu từ bàn phím, yêu cầu con trỏ để gán giá trị 
3. Sprint, Sprintln, Sprint: Xử lý chuỗi và trả về kết quả (không in ra màn hình, thường dùng để lưu trữ và xử lý tiếp)

CHẠY APP go run .
## Video 9
- todo 
Gin ?
Gin Framework là một web framework nhẹ và nhanh trong Go (Golang), dùng để xây dựng RESTful API và ứng dụng web. Gin dựa trên net/http, cung cấp cú pháp đơn giản, hiệu năng cao, và các tính năng như:
Routing: Dễ dàng định nghĩa endpoint (ví dụ: Get/users)
Middleware: Hỗ trợ xử lý xác thực, logging, CORS, v.v.
JSON handling: Tự động mã hóa/ gửi JSON(như c.JSON)
Hiệu nănng: Tối ưu tốc độ, phù hợp chp API quy mô lớn
## Video 10
- todo 

## Video 16
- todo 

## Video 17
- todo 

## Video 18
- todo 

## Video 19
- todo 

## Video 20
- todo 