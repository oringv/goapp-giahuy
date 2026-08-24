# Video 85-

1. Xem video, thực hành theo và lưu vào day20.go 
2. Ghi chú lại tất cả những kiến thức gì trong file day20.md
3. Phải đảm bảo code mẫu trong day20.go giống với video và chạy được không có lỗi 
4. Viết các lỗi thường gặp, nguyên nhân, cách giải quyết lỗi
5. Câu hỏi chưa hiểu cần hỏi để làm rõ 

## Video 85
- todo 
Tìm hiểu GOMAXPROCS và giới hạn CPU trong Golang
Mặc định Goroutione sẽ tự động chạy Parallelism kết hợp với Concurrent nếu ứng dụng chúng ta có nhiều Concurrent, cũng như sẽ tự động phân phối lượng CPU cho phù hợp để đạt được hiệu năng tốt nhất cho ứng dụng. Tuy nhiên, trong một số trường hợp điều này làm chiếm dụng hết toàn bộ CPU trên hệ thống khiến các chương trình khác bị treo. Vì thế chúng ta cần quản lý số lượng CPU phù hợp cho ứng dụng.

## Video 86
- todo 
Demo bài tập ứng dụng giúp quản lý hệ thống Server

## Video 87
- todo 
Hiển thị lượng phần trăm CPU đang sử dụng trên hệ thống 

## Video 88
- todo 
Bài tập tích hợp Goroutines vào dự án 

## Video 89
- todo 
Tối ưu mã nguồn và tìm hiểu rõ hơn về context.WithCancel và context.WithTimeout

## Video 90
- todo 
Hiển thị phần trăm dung dượng của Memory

## Video 91
- todo 
Sử dụng Channel để truyển dữ liệu giữa các Goroutine 

## Video 92
- todo 
Sử dụng Interface và Slice để tránh duplicate và tối ưu 

## Video 93
- todo 
Hiển thị toàn bộ thông tin trạng thái hoạt động của hệ thống trong Golang

## Video 94
- todo 
Hiển thị danh sách các chương trình đang hoạt động trên hệ thống
Tạo file mem.go trong package monitor  
Tái sử dụng code từ cpu.go để xử lý memory  
Dùng hàm VirtualMemoryWithContext để lấy thông tin RAM  
Trích xuất UsedPercent từ virtual memory  
In thông tin phần trăm sử dụng của RAM và CPU  
Xuất thông tin sang main.go để xử lý bên ngoài 

- PID Là số tiến trình chạy trên hệ thống 
- Mutex dùng để cho an toàn
  
  PHẦN 1: ĐẶT VẤN ĐỀ & MỤC TIÊU 
Giới thiệu: Đây là dự án "Hệ thống Giám sát Tài nguyên Thời gian thực". Ý tưởng bắt nguồn từ việc cần theo dõi sức khỏe máy tính một cách liên tục mà không gây treo máy.
Vấn đề: Các thông số như CPU, RAM, Mạng lấy dữ liệu với tốc độ khác nhau. Nếu chạy tuần tự (hết việc này mới đến việc kia), giao diện sẽ bị đứng và số liệu không chính xác.
Mục tiêu:
Ứng dụng đa luồng để thu thập dữ liệu song song.
Tổ chức code theo chuẩn Package chuyên nghiệp.
Đảm bảo an toàn dữ liệu khi nhiều luồng cùng hoạt động.
PHẦN 2: KIẾN TRÚC HỆ THỐNG & INTERFACE (3 - 7 PHÚT)
Tổ chức Package: Dự án chia làm 3 tầng:
Tầng Dữ liệu (models): Định nghĩa các "thùng chứa" dữ liệu (Struct) và kho lưu trữ tập trung (Map).
Tầng Cảm biến (monitors): Chứa các thợ đo đạc. Đây là nơi ứng dụng Interface.
Tầng Xử lý (processor): Bộ não điều phối các thợ đo và xử lý logic chi tiết từng ứng dụng.
Điểm nhấn - Interface Monitor:
Đây là "bản hợp đồng" chung cho mọi bộ đo.
Lợi ích: Tính mở rộng cực cao. Nếu muốn đo thêm nhiệt độ hay tốc độ quạt, chỉ cần tạo bộ đo mới tuân thủ Interface mà không cần sửa code ở hàm main. Đây là tư duy thiết kế hệ thống hiện đại.
PHẦN 3: ĐỘNG CƠ ĐA LUỒNG - CONCURRENCY (7 - 11 PHÚT)
Đây là phần "xịn" nhất của dự án, bạn nên tập trung giải thích kỹ:
Goroutines & Channels:
Mình tạo ra các Producers (Người sản xuất): Mỗi bộ đo chạy trên một Goroutine riêng, không ông nào phải đợi ông nào.
Dùng Channel làm đường ống vận chuyển: Đảm bảo dữ liệu chảy từ các bộ đo về trung tâm một cách trơn tru, không bị thất lạc.
Quản lý an toàn với Mutex:
Khi 4-5 ông thợ cùng chạy về ghi tên lên một cái bảng (Map), sẽ xảy ra xung đột (Race Condition).
Mình dùng sync.Mutex làm "ổ khóa". Ai muốn ghi vào bảng phải cầm khóa, ghi xong mới đưa khóa cho người tiếp theo. Điều này giúp dữ liệu luôn chính xác và chương trình không bao giờ bị "crash".
PHẦN 4: VÒNG ĐỜI & HIỂN THỊ (11 - 14 PHÚT)
Kiểm soát với Context: Ứng dụng sử dụng context.WithTimeout. Sau 60 giây (hoặc thời gian định sẵn), ứng dụng sẽ tự động "thu quân", đóng toàn bộ các luồng để trả lại tài nguyên cho máy tính.
Nhịp tim với Ticker: Thay vì in chữ chạy liên tục làm lóa mắt, mình dùng time.Ticker để cứ đúng 5 giây mới xuất bản báo cáo một lần. Điều này tạo ra sự ổn định cho giao diện người dùng.
Giải thích kết quả:
Dấu [...]: Thông số tổng thể (Sức khỏe toàn diện của máy).
Dấu {...}: Thông số chi tiết (Từng ứng dụng đang làm gì, chạy bao lâu).
PHẦN 5: TỔNG KẾT & MỞ RỘNG (14 - 15 PHÚT)
Kết quả: Dự án đã vận hành thành công một bộ máy thu thập dữ liệu phức tạp, an toàn và linh hoạt.

## Video 95
- todo 
Hiển thị danh sách các ứng dụng sử dụng CPU
 Tổng quan về channel trong Golang và cách tạo channel có buffer
 Sử dụng goroutine để nhận và gửi dữ liệu qua channel
 Cách đóng channel đúng cách và tránh deadlock
 Lọc và truyền dữ liệu theo điều kiện CPU lớn hơn 5% hoặc RAM lớn hơn 5%
 Sắp xếp danh sách các tiến trình tiêu tốn nhiều CPU và RAM nhất
 Hiển thị danh sách top 5 tiến trình tiêu tốn CPU và RAM nhiều nhất

 
## Video 96
- todo 
Tối ưu ứng dụng và cập nhật code
Giải thích cách sử dụng channel có buffer trong Golang và khi nào có thể bỏ buffer để tối ưu.
Loại bỏ biến không cần thiết (mch) để code gọn và hiệu quả hơn.
Cách xử lý return và continue trong vòng lặp để đảm bảo an toàn và logic rõ ràng.
Sử dụng vòng lặp for range để duyệt map và slice hiệu quả hơn, thay thế cho vòng for i.
So sánh kết quả giữa code cũ và code tối ưu, đảm bảo tính chính xác và hiệu suất.
Giải thích lý do tại sao CPU sử dụng lại có thể tăng đột biến do việc tạo nhiều goroutine.
Một số gợi ý tối ưu thêm khi xử lý nhiều goroutine để giảm tải hệ thống.

## Video 97
- todo 
Export toàn bộ danh sách ứng dụng vào CSV file
Hiển thị danh sách các ứng dụng sử dụng CPU và RAM cao nhất
Xây dựng hàm xuất dữ liệu sang file CSV trong Golang
Giải thích chi tiết các thao tác mở file, quyền truy cập, và định dạng CSV
Cách xử lý lỗi khi thao tác file trong Golang
Giải thích quyền truy cập file (chmod 0644) trong hệ điều hành Unix/Linux
Mã nguồn chi tiết và dễ hiểu

- Cài đặt bảng CSV để xem các cột Pid, Name, CPU... thẳng hàng và chuyên nghiệp hơn 

## Video 98
- todo 
Thông báo vào CSV file nếu CPU, RAM, Disk vượt ngưỡng cho phép 
Giới thiệu chức năng giám sát tài nguyên hệ thống bằng Golang
Lưu trữ thông tin sử dụng CPU, Memory, Disk vào file CSV
Thiết lập ngưỡng cảnh báo tài nguyên (CPU lớn hơn 60%, Memory lớn hơn 80%, Disk lớn hơn 60%)
Gửi cảnh báo khi vượt ngưỡng qua nhiều kênh (Telegram, Email, CSV)
Xử lý đồng thời với goroutines, channel giúp tăng hiệu suất
Viết log cảnh báo chuẩn và tối ưu ghi file an toàn
Giải thích chi tiết từng đoạn code, cách debug và xử lý lỗi
Hướng dẫn chạy thử và kiểm tra kết quả thực tế

KỊCH BẢN CHI TIẾT: QUẢN TRỊ DỮ LIỆU ĐA LUỒNG VỚI MAP & MUTEX TRONG GOLANG
PHẦN 1: DẪN NHẬP - BÀI TOÁN QUẢN LÝ DỮ LIỆU THỜI GIAN THỰC 
1. Lời chào và bối cảnh:
"Chào mọi người, trong kỷ nguyên số, dữ liệu không chỉ cần lớn (Big Data) mà còn phải nhanh (Fast Data). Hãy tưởng tượng bạn đang quản lý một hệ thống ngân hàng hoặc một sàn chứng khoán, thông tin thay đổi từng mili giây. Làm sao để lưu trữ và truy xuất nó tức thời?"
2. Tại sao lại là Map? (Phân tích cấu trúc dữ liệu):
So sánh Slice vs Map:
Với Slice, bạn phải duyệt từ đầu đến cuối (O(n)) để tìm xem "CPU" đang ở đâu.
Với Map, bạn tìm thông tin theo từ khóa (Key). Độ phức tạp là O(1) - tức là tìm 1 nhãn hay 1 triệu nhãn thì tốc độ vẫn nhanh như nhau.
Thiết kế Struct SystemStats: Giải thích về việc đóng gói dữ liệu gồm: Tên nhãn, Giá trị đo được, và Trạng thái báo động. Đây là cách tiếp cận hướng đối tượng trong Go.
PHẦN 2: "TỬ HUYỆT" CỦA MAP TRONG MÔI TRƯỜNG ĐA LUỒNG 
1. Bản chất của Concurrency trong Go:
"Go cho phép chúng ta tạo ra hàng ngàn Goroutine cực nhẹ. Nhưng sức mạnh này mang theo một hiểm họa: Race Condition (Xung đột tài nguyên)."
2. Tại sao Map lại sập? (Deep Dive):
Giải thích kỹ thuật: Map trong Go không được thiết kế để chịu tải ghi đồng thời (Not thread-safe). Khi 2 luồng cùng ghi vào 1 ô nhớ, cấu trúc nội bộ của Map sẽ bị phá vỡ. Go chọn cách "tự sát" (Panic) thay vì cho phép dữ liệu bị sai lệch.
Đây là triết lý của Go: "Sai thà chết còn hơn sai mà không biết".
3. Giải pháp Mutex - "Người bảo vệ thầm lặng":
Phân tích cơ chế sync.Mutex:
Lock(): Chiếm quyền kiểm soát tuyệt đối.
Unlock(): Trả tự do cho tài nguyên.
Nhấn mạnh: "Mutex không làm cho Map chạy nhanh hơn, nhưng nó làm cho Map chạy SỐNG SÓT trong môi trường đa luồng."
PHẦN 3: DÒNG CHẢY DỮ LIỆU - TỪ CẢM BIẾN ĐẾN KHO LƯU TRỮ 
1. Sự phối hợp giữa Channel và Map:
Phân tích dòng code: for stat := range statCh { ... }
"Channel đóng vai trò là Đường vận chuyển. Map đóng vai trò là Nhà kho. Dữ liệu từ các bộ cảm biến (Monitors) được đóng gói và ném vào đường ống. Một Goroutine duy nhất ở đầu kia sẽ nhặt từng món và xếp vào kho."
2. Tại sao lại dùng 1 Goroutine duy nhất để cập nhật Map?
Đây là một kỹ thuật tối ưu. Thay vì để 4-5 ông monitor cùng tranh nhau cái khóa Mutex, ta cho một ông "thủ kho" duy nhất làm nhiệm vụ ghi. Điều này giảm thiểu thời gian chờ (Contention) và giúp hệ thống mượt mà hơn.
3. Tối ưu hóa Channel có Buffer (Video 96):
Giải thích về việc cấp "kho đệm" cho channel. Nếu nhà kho (Map) bận xử lý, đường ống (Channel) vẫn có thể chứa thêm vài kiện hàng, giúp các bộ đo không bị đứng hình.
PHẦN 4: KHAI THÁC DỮ LIỆU TỪ MAP - TOP 5 & EXPORT
1. Xuất bản báo cáo với Ticker:
"Hệ thống không in dữ liệu vô tội vạ. Nhịp đập time.Ticker 5 giây giúp dữ liệu trong Map có thời gian 'lắng đọng' trước khi được xuất bản."
2. Bài toán sắp xếp (Sorting) - (Video 95):
"Dữ liệu trong Map là lộn xộn. Để tìm được Top 5, ta phải đổ dữ liệu từ Map ra một Slice trung gian, sau đó dùng sort.Slice để đưa những kẻ ngốn tài nguyên nhất lên đầu."
3. Persistence (Lưu trữ bền vững) - (Video 97 & 98):
Lấy dữ liệu từ Map để ghi vào CSV.
Cơ chế Alert: "Khi duyệt Map, nếu IsAlert == true, hệ thống sẽ kích hoạt hàm LogAlert. Đây là sự kết hợp hoàn hảo giữa Giám sát (Monitoring) và Phản ứng (Alerting)."
PHẦN 5: TỔNG KẾT & BÀI HỌC KINH NGHIỆM 
1. Những gì chúng ta đã đạt được:
Một hệ thống đa luồng an toàn tuyệt đối nhờ Mutex.
Khả năng truy xuất dữ liệu nhờ Map.
Hệ thống lưu trữ lịch sử và cảnh báo chuyên nghiệp.

1. Tầng Monitors (Cảm biến): Tính mở rộng tuyệt vời
Tại sao chuẩn:tách riêng cpu.go, disk.go, mem.go, net.go. Đây là cách thiết kế theo Single Responsibility Principle (Nguyên tắc đơn trách nhiệm).
 "Nếu sau này tôi muốn đo thêm nhiệt độ GPU, tôi chỉ cần tạo thêm file gpu.go trong package monitors mà không cần chạm vào bất kỳ dòng code nào của CPU hay RAM. Điều này giúp hệ thống cực kỳ an toàn và dễ bảo trì."
2. Tầng Models (Dữ liệu): Sự tập trung và bảo mật
Toàn bộ Struct và Map/Mutex vào đây. Đây là "nguồn sự thật duy nhất" (Single Source of Truth).
 "Mọi dữ liệu trong hệ thống đều phải tuân thủ khuôn mẫu từ tầng models. Việc tập trung Mutex tại đây giúp tôi kiểm soát được mọi luồng truy cập dữ liệu, đảm bảo không bao giờ xảy ra xung đột (Race Condition) dù có hàng nghìn Goroutine chạy cùng lúc."
3. Tầng Processor (Xử lý): Tách biệt Logic và Giao diện
Tại sao chuẩn: tách riêng logic sắp xếp Top 5 và ghi file CSV/Log ra khỏi hàm main.
 "Hàm main của tôi chỉ đóng vai trò là 'Người nhạc trưởng' để khởi động hệ thống. Mọi logic tính toán phức tạp hay thao tác với ổ đĩa đều được tầng processor đảm nhận. Nếu tôi muốn thay đổi cách ghi file từ CSV sang Database, tôi chỉ cần sửa ở tầng này."