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
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"goapp-giahuy/learninggo/day20/processor"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStats)

	// ✅ FIX: Khai báo danh sách các bộ giám sát
	monitorList := []monitors.Monitor{
		&monitors.MemoryMonitor{},
		&monitors.CPUMonitor{},
	}

	fmt.Println("🚀 Đang khởi tạo hệ thống giám sát...")

	// Duyệt qua danh sách và chạy mỗi con vật trong 1 Goroutine riêng
	for _, m := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, m)
	}

	// Goroutine đóng channel
	go func() {
		wg.Wait()
		close(statCh)
	}()

	// ✅ LOGIC IN RA GIỐNG ẢNH TERMINAL CỦA BẠN:
	for stat := range statCh {
		// Chỉ dùng dòng in này để ra kết quả có dấu { }
		fmt.Printf("%v\n", stat)
	}

	fmt.Println("🏁 Chương trình kết thúc.")
}
