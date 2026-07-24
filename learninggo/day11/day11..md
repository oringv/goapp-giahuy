# Video 34-

1. Xem video, thực hành theo và lưu vào day3.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day3.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 34
- todo 
Khai Báo Interface – Tại Sao Cần Dùng Interface?
Interface là một tập hợp các phương thức (methods) mà một kiểu dữ liệu phải triển khai.

## Video 35
- todo 
Ứng Dụng Interface - Tích Hợp Interface Vào Code Go

## Video 36
- todo 
Cách Quản Lý Interface Hiệu Quả Trong Golang

## Video 37
- todo 
 Tìm Hiểu Empty Interface – Khi Nào Nên Dùng?
 - Empty Interface (interface{}) là một Interface không có phương thức nào.
 - Nó có thể chứa bất kỳ giá trị nào, tương tự như kiểu Object trong các ngôn ngữ khác.

 ## Video 38
- todo 
Phân biệt Type Assertion & Type Switch
- Type Assertion được sử dụng để truy xuất giá trị cụ thể từ một Empty Interface (interface{})
- Nó giúp chúng ta chuyển đổi một giá trị từ kiểu interface{} sang kiểu dữ liệu cụ thể.
- Type switch 
Là một cấu trúc điều khiển cho phép kiểm tra kiểu dữ liệu của một Empty Interface (interface{})
- Nó giống như switch thông thường, nhưng thay vì so sánh giá trị, nó so sánh kiểu dữ liệu

## Video 39
- todo 
Generics Là Gì? – Cách Dùng Generics Trong Golang
- Generic là một tính năng cho phép viết code linh hoạt, có thể làm việc với nhiều kiểu dữ liệu khác nhau mà không cần viết lại code cho từng kiểu dữ liệu cụ thể.

Golang sử dụng cú pháp đơn giản để định nghĩa Generic:
Type Parameters: Định nghĩa các kiểu dữ liệu tổng quát.
Type Constraints: Giới hạn các kiểu dữ liệu có thể sử dụng (any, comparable, Interface và constraints)

## Video 40
Cách Dùng Generics với Comparable & CmpOrdered
