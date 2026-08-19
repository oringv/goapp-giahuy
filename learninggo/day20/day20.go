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
	// 1. Thiết lập thời gian chạy 20 giây
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStats)

	// Danh sách các bộ giám sát
	monitorList := []monitors.Monitor{
		&monitors.MemoryMonitor{},
		&monitors.CPUMonitor{},
		&monitors.NetMonitor{},
		&monitors.DiskMonitor{},
	}

	fmt.Println("🚀 Hệ thống giám sát bắt đầu hoạt động...")

	// 2. Kích hoạt các Goroutine làm việc (Producers)
	for _, m := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, m)
	}

	go func() {
		for stat := range statCh {
			models.StatsMutex.Lock()
			models.Stats[stat.Label] = stat
			models.StatsMutex.Unlock()

			fmt.Printf("📥 Log: %+v\n", stat)
		}
	}()

	go func() {
		wg.Wait()
		close(statCh)
	}()

	fmt.Println("⏳ Đang thu thập dữ liệu trong 20 giây...")
	<-ctx.Done() // Đợi cho đến khi Timeout (20s) xảy ra

	fmt.Println("\n📊 TỔNG KẾT DỮ LIỆU TRONG MAP:")
	models.StatsMutex.Lock()
	for label, data := range models.Stats {
		fmt.Printf("{%s: %s}\n", label, data.Value)
	}
	models.StatsMutex.Unlock()

	fmt.Println("🏁 Chương trình kết thúc thành công.")
}
