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
	// 1. Chạy trong 30 giây rồi tự đóng
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStats)

	monitorList := []monitors.Monitor{
		&monitors.MemoryMonitor{},
		&monitors.CPUMonitor{},
		&monitors.NetMonitor{},
		&monitors.DiskMonitor{},
	}

	fmt.Println("🚀 Hệ thống giám sát bắt đầu hoạt động...")

	for _, m := range monitorList {
		wg.Add(1)
		go func(val monitors.Monitor) {
			processor.RunMonitor(ctx, &wg, statCh, val)
		}(m) // Truyền m vào đây để tạo bản sao riêng cho mỗi Goroutine
	}

	// 3. Consumer: Lưu dữ liệu vào Map
	go func() {
		for stat := range statCh {
			models.StatsMutex.Lock()
			models.Stats[stat.Label] = stat
			models.StatsMutex.Unlock()
		}
	}()

	// 4. Đóng channel khi các producers xong việc
	go func() {
		wg.Wait()
		close(statCh)
	}()

	// 5. In báo cáo mỗi 5 giây
	printTicker := time.NewTicker(5 * time.Second)
	defer printTicker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-printTicker.C:
				fmt.Println("\n===== 📊 System Status (Live) =====")
				models.StatsMutex.Lock()
				for _, stat := range models.Stats {
					fmt.Printf("[%s]: %s \n", stat.Label, stat.Value)
				}
				models.StatsMutex.Unlock()

				// Lấy thông tin chi tiết tiến trình
				processor.GetTopProcesses(ctx)
			}
		}
	}()

	// Đợi đến khi hết timeout
	<-ctx.Done()
	fmt.Println("\n🎉 Chương trình kết thúc.")
}
