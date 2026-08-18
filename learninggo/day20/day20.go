package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func heavyTask(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 0; i < 100e8; i++ {
// 		sum += i
// 	}

// }

// func main() {
// 	numCPU := runtime.NumCPU()
// 	fmt.Println("So luong CPU:", numCPU)

// 	runtime.GOMAXPROCS(numCPU)

// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go heavyTask(&wg)
// 	wg.Wait()
// 	fmt.Println("Tong thoi gian:", time.Since(start))
// }

//**BÀI 87**//

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/processor"
	"sync"
	"time"
)

func main() {
	// Chạy trong 10 giây rồi tự nghỉ
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("🚀 Đang khởi tạo hệ thống giám sát...")

	// Gọi hàm từ package processor
	go processor.RunMonitor(ctx, &wg)

	wg.Wait()
	fmt.Println("🏁 Chương trình kết thúc thành công.")
}
