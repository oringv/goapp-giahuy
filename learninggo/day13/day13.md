# Video 48

1. Xem video, thực hành theo và lưu vào day3.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day3.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 48
- todo 
Những Cách Khởi Tạo Slice Trong Golang
Cách tạo slice từ array
Chỉ mục trong slice (start, end)
Cách kiểm tra slice và array
Dùng make để khởi tạo slice

## Video 49
- todo 
Cách Sử Dụng append() Với Slice Hiệu Quả
Cách duyệt slice trong Golang
Sử dụng vòng lặp for và range với slice
Làm việc với slice lồng nhau (nested slices)
Thêm phần tử vào slice với append

## Video 50
- todo 
 Subslice – Hiểu Rõ Cách Hoạt Động Của cap() Và len() 
 Tạo sub slice là tạo ra 1 slice con từ 1 slice cha (slice ban đầu)
 Subslice là gì?
    Cách tạo subslice từ một slice gốc
    Cách hoạt động của chỉ số trong subslice
    Ảnh hưởng của len() và cap() khi tạo subslice

## Video 51
- todo 
 Những Hàm Xử Lý Với Slice Quan Trọng Trong Golang    
  Giới thiệu về Slice và sự khác biệt với Array
   Cách sử dụng package slices (Go 1.21+)
   Các hàm quan trọng: Clone(), Equal(), Index(), Contains(), Insert(), Delete(), Reverse(), Sort()
   Thực hành viết code xử lý Slice trong Golang

## Video 52
- todo
   Demo Bài tập Quản lý sinh viên và giảng viên
   Ôn tập kiến thức về Array và Slice trong Golang
   Demo xây dựng chương trình quản lý sinh viên & giảng viên
   Thêm, xóa, sửa, tìm kiếm và hiển thị danh sách sinh viên/giảng viên
   Ứng dụng thực tế của Array và Slice trong lập trình Golang

## Video 53
- todo
   Bài tập xây dựng Main Menu cho dự án
   Tạo project quản lý sinh viên & giảng viên
   Thiết kế menu chính với các chức năng
   Xử lý input từ bàn phím với Bufio và kiểm tra dữ liệu hợp lệ
   Clear màn hình trên macOS và Windows
   Gợi ý xây dựng package cho sinh viên và giảng viên

## Video 54
- todo
   Bài tập xây dựng menu cho sinh viên và giảng viên
   Tạo menu quản lý sinh viên với các chức năng: thêm, xóa, sửa, danh sách, tìm kiếm
   Tạo menu quản lý giảng viên tương tự sinh viên
   Tích hợp menu sinh viên và giảng viên vào menu chính
   Xử lý input và quay lại menu chính với clear screen

## Video 55
- todo
   Bài tập Xây dựng chức năng Thêm sinh viên
   Tạo struct Student với các trường ID, tên, lớp, điểm
   Xây dựng chức năng thêm sinh viên với kiểm tra input hợp lệ
   Sử dụng slice để lưu trữ nhiều sinh viên
   Hiển thị danh sách sinh viên với vòng lặp và kiểm tra rỗng
   Ôn tập append slice và duyệt slice trong Golang

 ## Video 56
- todo  
Tạo hàm GetInfo với receiver để hiển thị thông tin sinh viên
 Tính điểm trung bình cho sinh viên với hàm CalculateAverageScore
 Xây dựng struct và chức năng thêm giảng viên tương tự sinh viên
 Tạo hàm GetInfo và CalculateSalary cho giảng viên
 Hiển thị danh sách sinh viên và giảng viên với định dạng dễ đọc


## Video 57
- todo
Bài tập Xây dựng chức năng xử lý tránh duplicate id
Kiểm tra trùng ID khi thêm sinh viên và giảng viên
 Tạo hàm GetID và IsIDUnique cho struct Student
 Áp dụng tương tự cho struct Teacher
 Tối ưu code bằng generic function để tránh lặp code
 Xử lý lỗi vòng lặp package trong Golang

## Video 57
- todo
Bài tập Xây dựng chức năng sửa sinh viên
Tìm sinh viên theo ID trong student list
Cập nhật tên, lớp với tùy chọn giữ nguyên giá trị cũ
Xử lý cập nhật điểm bằng cách duyệt slice score
Tạo hàm GetOptionalString và GetOptionalPositiveFloat
Kiểm tra trường hợp ID không tồn tại