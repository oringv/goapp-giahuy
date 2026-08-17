# Video 83-

1. Xem video, thực hành theo và lưu vào day19.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day19.md
3. Phải đảm bảo code mẫu trong day19.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 83
- todo 
Tìm hiểu lỗi Race Condition và sử dụng sync.Mutex để khắc phục trong Golang
Hiểu lỗi Race Condition là gì và khi nào xảy ra
    Ví dụ thực tế tạo 1000 goroutine gây lỗi
    Cách sử dụng sync.WaitGroup để chờ tất cả goroutine
    Sửa lỗi Race Condition với sync.Mutex
    So sánh kết quả chạy khi có và không có khóa đồng bộ

Mutex (viết tắt là Mutual Exclution) là một cơ chế đồng bộ hóa trong lập trình đa luồng (multithreading), dùng để đảm bảo rằng chỉ có một luồng (thread) tại một thời điểm có thể truy cập vào vùng tài nguyên dùng chung (shared resource), chẳng hạn như chế biến toàn cục, bộ nhớ hoặc tệp tin.

## Video 84
- todo 
So sánh giữa Sequential, Concurrent và Parallelism
1. Sequential (tuần tự)
- Mô tả:Một lõi CPU xử lý từng tác vụ một cách tuần tự, ví dụ Task 1 xong rooif mới đến Task 2.
- Chạy một chương trình rồi chờ xong, mới chạy chương trình khác.

2. Concurrent (Đồng thời)
- Mô tả: Một lõi CPU xử lý các tác vụ đan xen nhau (ví dụ xử lý một phần của Task 1, rồi chuyển sang xử lý một phần của Task 2, rồi quay lại Task 1,...) nên có tính đồng thời
- Các tác vụ có thể hoàn thành gần cùng lúc, dù vẫn chỉ chạy trên một lõi CPU.
- Đa nhiệm trên một lõi CPU (multitasking sử dụng kỹ thuật như time slicing).

3. Parallel (song song)
- Mô tả: Hai lõi CPU, mỗi lõi xử lý một tác vụ riêng biệt từ đầu đến cuối (Các tác vụ này thường là các tác vụ nặng). Các tác vụ này không đan xen, nhưng được xử lý cùng lúc.
- Đây là tính song song vì mỗi tác vụ chỉ dùng một luồng và không bị ngắt quãng.
- Hai tiến trình độc lập chạy trên hai lõi khác nhau.

4. Concurrent, Parallel (Đồng thời và song song)
- Mô tả: Nhiều lõi CPU xử lý các tác vụ song song nhau vừa đồng thời, tức là xử lý nhiều tác vụ cùng lúc và mỗi tác vụ có thể được chia nhỏ để chạ đồng thời.
- Đây là hình thức mạnh nhất, tận dụng cả đa lõi CPU và kỹ thuật xử lý đồng thời.
- Hệ thống đa luồng (multithreading) trên nhiều lõi CPU.

## Video 85
- todo 
Tìm hiểu GOMAXPROCS và giới hạn CPU trong Golang
Mặc định Goroutione sẽ tự động chạy Parallelism kết hợp với Concurrent nếu ứng dụng chúng ta có nhiều Concurrent, cũng như sẽ tự động phân phối lượng CPU cho phù hợp để đạt được hiệu năng tốt nhất cho ứng dụng. Tuy nhiên, trong một số trường hợp điều này làm chiếm dụng hết toàn bộ CPU trên hệ thống khiến các chương trình khác bị treo. Vì thế chúng ta cần quản lý số lượng CPU phù hợp cho ứng dụng.