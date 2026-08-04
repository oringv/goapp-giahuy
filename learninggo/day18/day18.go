package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // 1. Hàm thực hiện công việc
// func task(id int, ch chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Task %d bat dau \n", id)

// 	// Giả lập công việc tốn 1 giây
// 	time.Sleep(1 * time.Second)

// 	// Gửi nhiều dòng thông báo vào channel (giống trong ảnh video)
// 	ch <- fmt.Sprintf("Task %d ket thuc", id)
// 	ch <- fmt.Sprintf("Ending task %d", id)
// 	ch <- fmt.Sprintf("Ending task %d", id)
// 	ch <- fmt.Sprintf("Ending task %d", id)
// 	ch <- fmt.Sprintf("Ending task %d", id)
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	// ✅ QUAN TRỌNG: Khai báo bộ đệm (ví dụ 20) để chứa được hết các tin nhắn
// 	// Nếu không có số 20, chương trình sẽ bị treo (Deadlock)
// 	ch := make(chan string, 20)

// 	// 2. Chạy 4 Goroutine song song
// 	for i := 1; i <= 4; i++ {
// 		wg.Add(1)
// 		go task(i, ch, &wg)
// 	}

// 	// 3. Đợi các task làm xong rồi mới đóng channel
// 	wg.Wait()
// 	close(ch)

// 	// 4. In toàn bộ kết quả từ channel ra màn hình
// 	for value := range ch {
// 		fmt.Println(value)
// 	}

// 	// In tổng thời gian (Sẽ ra ~1 giây vì chạy song song)
// 	fmt.Println("Tong thoi gian hoan thanh: ", time.Since(start))
// }

//** Bài 79 **//

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		ch1 <- "Data from channel 1"
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "Data from channel 2"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println(msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println(msg2)
// 		}

// 	}
// }

//** Bài 81 **//

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// // Hàm mô phỏng một nhân viên đang làm việc
// func employee(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// Khi hết thời gian (Timeout), case này sẽ được kích hoạt
// 			fmt.Println("Cong viec da bi huy: ", ctx.Err())
// 			return
// 		default:
// 			// Lấy giá trị từ context bằng Key "priority"
// 			priority := ctx.Value("priority")
// 			fmt.Println("Dang lam viec cua task uu tien voi muc do: ", priority)

// 			// Nghỉ 100ms trước khi lặp lại
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}
// }

// func main() {
// 	// 1. Tạo một context với thời gian chờ (Timeout) là 2 giây
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

// 	// Đảm bảo giải phóng tài nguyên sau khi xong
// 	defer cancel()

// 	// 2. Gán thêm một giá trị vào context (Key: "priority", Value: "hight")
// 	// Lưu ý: "hight" là lỗi chính tả trong video, đúng phải là "high"
// 	ctx = context.WithValue(ctx, "priority", "hight")

// 	// 3. Chạy hàm nhân viên trong một Goroutine riêng (chạy ngầm)
// 	go employee(ctx)

// 	// 4. Hàm main đợi 3 giây để chúng ta quan sát kết quả
// 	// Vì Timeout chỉ có 2 giây nên sau 2 giây nhân viên sẽ nghỉ việc
// 	time.Sleep(3 * time.Second)
// }

//** Bài 82 **//

import (
	"context"
	"fmt"
	"time"
)

// Hàm nấu Phở - Tốn 1 giây
func cookPho(ctx context.Context, ch chan<- string) {
	fmt.Println("Bat dau nau Pho")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Huy nau Pho")
	case ch <- "Pho da nau xong":
	}
}

// Hàm nấu Cháo - Tốn 3 giây (Món này sẽ bị Timeout vì quá lâu)
func cookChao(ctx context.Context, ch chan<- string) {
	fmt.Println("Bat dau nau Chao")
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Huy nau Chao")
	case ch <- "Chao da nau xong":
	}
}

// Hàm nấu Cơm - Tốn 1.5 giây
func cookCom(ctx context.Context, ch chan<- string) {
	fmt.Println("Bat dau nau Com")
	time.Sleep(1500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("Huy nau Com")
	case ch <- "Com da nau xong":
	}
}

func main() {
	// 1. Thiết lập thời gian chờ tối đa là 2 giây
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 2. Tạo các kênh (channels) để nhận món ăn
	chPho := make(chan string)
	chChao := make(chan string)
	chCom := make(chan string)

	// 3. Bắt đầu nấu 3 món ăn cùng lúc (Goroutines)
	go cookPho(ctx, chPho)
	go cookChao(ctx, chChao)
	go cookCom(ctx, chCom)

	// 4. Vòng lặp đợi nhận 3 kết quả
	for i := 1; i <= 3; i++ {
		select {
		case res := <-chPho:
			fmt.Println("Nhan duoc: ", res)
		case res := <-chChao:
			fmt.Println("Nhan duoc: ", res)
		case res := <-chCom:
			fmt.Println("Nhan duoc: ", res)
		case <-ctx.Done():
			// Nếu đã quá 2 giây mà chưa nhận đủ món
			fmt.Println("Timeout, khong nhan mon")
			return
		}
	}
}
