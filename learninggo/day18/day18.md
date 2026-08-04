# Video 78-82

1. Xem video, thực hành theo và lưu vào day78.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day18.md
3. Phải đảm bảo code mẫu trong day78.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 78
- todo 
 Kết hợp sync.WaitGroup và Channel khi làm Goroutines
  Ôn lại kiến thức Channel & WaitGroup
    Cách truyền dữ liệu bằng Channel
    Sử dụng Channel có buffer vs không buffer
    Áp dụng for-range để đọc dữ liệu từ Channel
    Kết hợp Channel và WaitGroup chuẩn mô hình
    Phân tích lợi - hại khi dùng Channel có buffer

## Video 79
- todo 
Sử dụng Select để tối ưu hiệu suất khi làm việc với Channel trong Goroutines
Select là gì trong Golang?
 Cách xử lý dữ liệu từ nhiều channel một cách linh hoạt
 So sánh giữa dùng select và cách lấy dữ liệu tuần tự
 Tối ưu hiệu suất chương trình với select và goroutine

Sử dụng select để xử lý nhiều channel
Select giúp xử lý nhiều channel một cách linh hoạt và hiệu quả. Select sẽ ưu tiên channel nào sẵn sàng trước, tránh việc chờ tuần tự

## Video 80
- todo 
Những điểm cần lưu ý khi sử dụng Channel trong Golang
Khi nào nên đóng channel và cách đóng đúng cách
 Tránh lỗi deadlock khi truyền dữ liệu vào channel
 Không được tạo channel sai cú pháp
 Cách dùng make(chan Type) để khai báo channel đúng chuẩn

 Những lưu ý khi sử dụng channel
 Đóng Channel đúng cách:
 -SHOULD: Đóng channel (close(ch)) khi không còn dữ liệu để gửi.
 -SHOULD NOT: Đừng đóng channel nếu vẫn có Goroutine đang gửi dữ liệu (gây panic)

## Video 81
- todo 
 Sử dụng Context để quản lý thời gian sống trong Golang
Context là gì? Tại sao nên dùng?
 Các loại Context: WithTimeout, WithCancel, WithValue
 Hướng dẫn sử dụng context để giới hạn thời gian xử lý
 Cách truyền dữ liệu metadata giữa các goroutine
 Tránh leak bộ nhớ khi sử dụng context
 Ví dụ thực tế: Mô phỏng người sếp giao task với thời hạn

 Context trong Golang là một cơ chế giúp quản lý và kiểm soát thời gian sống cũng như hành vi của các tác vụ trong chương trình, đặc biệt khi làm việc với Goroutine. Chúng ta có thể hình dung Context như một bộ điểu khiển từ xa.
 Hủy(cancel) các Goroutine khi không cần thiết nữa.
 Đặt thời gian sống (tinmeout)
 Truyền dữ liệu (metadata) giữa các Goroutine mà không cần biết toàn cục 

## Video 82
- todo 
 Áp dụng thực tế Context và Channel khi làm Goroutines
  Ôn lại kiến thức về context trong Golang
    Thực hành ví dụ thực tế: nấu phở và cháo với timeout
    Sử dụng select để xử lý channel và timeout
    Cách tạo và huỷ context đúng chuẩn
    Hiểu cơ chế hoạt động của goroutine khi bị timeout