# Video 8-11

1. Xem video, thực hành theo và lưu vào day3.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day1.md
3. Phải đảm bảo code mẫu trong day3.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 8
- todo 
Tìm hiểu Package fmt– In, Định Dạng, Nhập/Xuất Dữ Liệu |
Package fmt trong Go được xây dựng để nhằm mục đích tương tác trực tiếp với dữ liệu như hiển thị dữ liệu, lấy dữ liệu từ bàn phím. Một số hàm trong package fmt hay sử dụng:
1. Print, Println, Printf: Hiển thị dữ liệu trực tiếp ra màn hình 
2. Scan, Scanln, Scanf: Nhập dữ liệu từ bàn phím, yêu cầu con trỏ để gán giá trị 
3. Sprint, Sprintln, Sprint: Xử lý chuỗi và trả về kết quả (không in ra màn hình, thường dùng để lưu trữ và xử lý tiếp)

CHẠY APP go run .

## Video 9
- todo 
Cách Dùng Formatting Verbs Trong fmt - Định Dạng Dữ Liệu
Go Formatting Verbs (còn được gọi là Printf Verbs hay Format Specifiers) được sử dụng trong các hàm như fmt.Printf, fmt.Sprintf.. để kiểm soát cách giá trị được định dạng và in ra dưới dạng chuỗi. Chúng hoạt động như những placeholder cho các giá trị mà bạn muốn chèn vào chuỗi kết quả, và đồng thời chỉ định kiểu dữ liệu và định dạng cụ thể cho các giá trị đó 

Một số Formatting Verbs phổ biến:
%v: Dùng để lấy giá trị của 1 biến.
%T: Dùng để hiển thị kiểu dữ liệu 1 biến.
%s: String (Chuỗi).
%d: Integer (decimal - Số nguyên).
%.f: Floating-point number (Số thực).
%t: Boolean (Kiểu đúng/sai).
%c: Character (rune - Ký tự).
%p: Pointer (Con trỏ địa chỉ).
%%: Trường hợp khi cần hiển thị dấu %.

## Video 10
- todo 
Toán Tử & Phép Gán Trong Go 
Toán tử	Mô tả	                      Ví dụ (a=10, b=3)
+	Cộng	                            a + b = 13
-	Trừ	                                a - b = 7
*	Nhân	                            a * b = 30
/	Chia (Lấy phần nguyên nếu là int)	a / b = 3
%	Chia lấy dư (Modulus)	            a % b = 1
++	Tăng 1 đơn vị	                    a++ (a thành 11)
--	Giảm 1 đơn vị	                    a-- (a thành 9)

## Video 11
- todo 
Phép So Sánh & Phép Luận Lý 
Bảng 1: Phép so sánh
Phép toán	  Tên	                      Ví dụ
==	        Phép bằng	                  x == y
!=	        Không bằng	                  x != y
>	        Lớn hơn	                      x > y
<	        Bé hơn	                      x < y
>=	        Lớn hơn hoặc bằng	          x >= y
<=	        Bé hơn hoặc bằng	          x <= y

Bảng 2: Phép luận lý (Logical)
Phép toán	    Tên	                        Mô tả
&&	            Và                  Trả về TRUE khi toàn bộ mệnh đề là TRUE
||		        Hoặc                Trả về TRUE khi chỉ cần 1 mệnh đề là TRUE
!	            Phủ định            Phủ định mệnh đề (Đúng thành Sai, Sai thành Đúng)


