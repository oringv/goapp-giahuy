# Video 64-

1. Xem video, thực hành theo và lưu vào day3.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day3.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 64
- todo 
Khởi tạo project quản lý thư viện với nhiều file Golang
 Xây dựng menu chọn chức năng với vòng lặp for và switch case
 Áp dụng lại kiến thức từ các bài trước như GetPositiveInt()
 Tách chức năng thêm/sửa/xem sách và người mượn thành các hàm
 Cách xử lý error và trả về lỗi khi thực hiện chức năng

## Video 65
- todo 
 Bài tập Xây dựng Chức năng thêm Sách cho thư viện 
 Xây dựng chức năng thêm sách cơ bản
 Sử dụng uuid để tạo ID tự động
 Lưu dữ liệu sách bằng struct và map
 Hiển thị danh sách sách đã thêm và tránh lỗi ghi đè

## Video 66
- todo 
 Bài tập Tối ưu chức năng thêm sách giúp code hiệu quả
 Tách riêng xử lý dữ liệu (store) và xử lý logic (service)
 Khai báo struct Library quản lý sách với map
 Sử dụng con trỏ trong Golang để truyền dữ liệu hiệu quả
 Tạo constructor NewLibrary() để khởi tạo thư viện
 Gọi hàm thêm sách (AddBook) thông qua receiver hoặc truyền tham số
 Quản lý dự án rõ ràng bằng cách chia file theo chức năng

## Video 67
- todo 
 Bài tập Hiển thị danh sách của sách trong thư viện
  Demo chức năng hiển thị danh sách sách
    Kiểm tra thư viện rỗng và xử lý thông báo
    Duyệt slice với for range để in thông tin sách
    Thêm trạng thái sách "Còn / Đã mượn" vào struct

## Video 68
- todo 
Bài tập Thêm người mượn sách và danh sách người mượn
Tạo struct để lưu thông tin người mượn (ID, Name, Email)
 Lưu người mượn vào map[string]Borrower
 Hiển thị danh sách người mượn từ map
 Kiểm tra dữ liệu rỗng và xử lý hợp lý

## Video 69
- todo 
Bài tập Xây dựng chức năng mượn sách
Thiết kế mô hình giao dịch mượn sách trong thư viện
Kiểm tra điều kiện đầu vào (sách tồn tại, người mượn tồn tại, sách đã được mượn hay chưa)
Cập nhật trạng thái sách và lưu lịch sử mượn

## Video 70
- todo 
Bài tập Hiển thị lịch sử mượn sách theo người mượn
Hiểu cách kiểm tra và lấy thông tin lịch sử mượn sách của người dùng
Cách viết hàm để lấy toàn bộ giao dịch mượn sách của người dùng từ session
Cách xử lý và hiển thị thông tin về sách đã mượn, ngày mượn và ngày trả một cách rõ ràng

## Video 71
- todo 
Bài tập Xây dựng chúc năng trả sách cho thư viện
Hiểu cách kiểm tra trạng thái trả sách qua ID giao dịch
Cập nhật thông tin sách và ngày trả khi thực hiện trả sách

## Video 72
- todo 
 Bài tập Chức năng tìm kiếm sách với tiêu đề và tác giả