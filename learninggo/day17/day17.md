# Video 73-77

1. Xem video, thực hành theo và lưu vào day17.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day17.md
3. Phải đảm bảo code mẫu trong day17.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 73
- todo 
Tìm hiểu khái niệm Concurrent, Parallelism, Sequential
Goroutine là một "luồng nhẹ"(lightweight thread) mà Go tạo ra để chạy một tác vụ nào đó đồng thời(cuncurrent) với chương trình chính

## Video 74
- todo 
Cách thức hoạt động của Goroutines trong Golang
Cách hoạt động của Goroutine trong Go
 So sánh giữa chạy tuần tự và chạy đồng thời
 Đo thời gian thực tế khi chạy tuần tự và khi dùng Goroutine
 Tạo và sử dụng Goroutine đơn giản chỉ với từ khóa go
 Ứng dụng Goroutine trong xử lý tác vụ đồng thời

 Cách hoạt động của goroutine
 Nếu hàm main () kết thúc, tất cả Goroutine cũng bị dừng, ngay cả khi chúng chưa hoàn thành.
 Trong ví dụ trên, time.Sleep là cách tạm thời để chờ, nhưng không phải giải pháp tốt.

## Video 75
- todo 
 Đồng bộ hóa với sync.Waitgroup trong Goroutines
 Vấn đề khi dùng time.Sleep để chờ Goroutine kết thúc
 Cách tự viết biến đếm thủ công để theo dõi số lượng Goroutine
 Giới thiệu sync.WaitGroup để quản lý đồng bộ goroutine hiệu quả hơn
 Sử dụng defer để đảm bảo hàm chạy cuối cùng trong Goroutine

## Video 76
- todo 
Tìm hiểu Unbuffered Channel khi sử dụng Goroutines
Channel trong Golang là một ống dẫn (pipeline) cho phép các Goroutine giao tiếp và truyền dữ liệu với nhau một cách an toàn. Bạn có thể hình dung Channel như một băng chuyền trong nhà máy: một Goroutine đặt dữ liệu lên băng chuyền (gửi), và một Goroutine khác lấy dữ liệu từ băng chuyền (nhận)

Các loại channel 
- Channel không buffer (Unbuffered Channel): Chỉ chứa được 1 giá trị tại một thời điểm.
+ Gửi(ch <- value) sẽ bị chặn cho đến khi có Goroutine nhận (<-ch)
+ Nhận (<-ch) sẽ bị chặn cho đến khi có dữ liệu được gửi
+ Ví dụ: Hộp thư chỉ chứa được 1 lá thư. Người gửi phải đợi người nhận lấy thư trước khi gửi lá thư tiếp theo.

- Channel có buffer (Buffered Channel): Có thể chứa nhiều giá trị (kích thước buffer do bạn định nghĩa).
+ Gửi (<- value) chỉ bị chặn khi buffer đầy.
+ Nhận (<-ch) chỉ bị chặn khi buffer trống
+ Ví dụ: Hộp thư chưa được 5 lá thư. Người gửi có thể gửi 5 thư trước khi phải đợi.

## Video 77
- todo 
Tìm hiểu Buffered Channel khi sử dụng Goroutines
Khái niệm channel có buffer là gì
Sự khác biệt giữa channel có và không có buffer
Cách khai báo, truyền dữ liệu và xử lý block khi buffer đầy hoặc rỗng
Demo trực tiếp với ví dụ minh hoạ và xử lý lỗi thường gặp