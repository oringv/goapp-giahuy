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
	// 1. Thiết lập thời gian chạy tổng cộng là 30 giây
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

	// 2. Chạy các Goroutine thu thập dữ liệu (Producers)
	for _, m := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, m)
	}

	// 3. Goroutine thu thập dữ liệu từ channel và lưu vào Map tập trung
	go func() {
		for stat := range statCh {
			models.StatsMutex.Lock()
			models.Stats[stat.Label] = stat
			models.StatsMutex.Unlock()
			// Bạn có thể mở log dưới đây nếu muốn xem dữ liệu thô chảy về
			// fmt.Printf("📥 Log: %+v\n", stat)
		}
	}()

	// 4. Goroutine đợi các task hoàn thành để đóng channel
	go func() {
		wg.Wait()
		close(statCh)
	}()

	// 5. Cài đặt Ticker để in báo cáo "Live" mỗi 5 giây
	printTicker := time.NewTicker(5 * time.Second)
	defer printTicker.Stop()

	// 6. Goroutine in báo cáo trực tiếp (Live Status)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-printTicker.C:
				fmt.Println("\n===== 📊 System Status (Live) =====")
				models.StatsMutex.Lock()
				// Sửa lỗi 'for_,' và 'stat.Name'
				for _, stat := range models.Stats {
					fmt.Printf("[%s] %s \n", stat.Label, stat.Value)
				}
				models.StatsMutex.Unlock()

				// Gọi hàm lấy thông tin tiến trình từ package processor
				topInfo := processor.GetTopProcesses(ctx)
				fmt.Println(topInfo)
			}
		}
	}()

	// 7. QUAN TRỌNG: Hàm main đứng đợi ở đây cho đến khi hết 30 giây
	fmt.Println("⏳ Đang thu thập dữ liệu và cập nhật trực tiếp mỗi 5 giây...")
	<-ctx.Done()

	// 8. TỔNG KẾT CUỐI CÙNG sau khi hết 30 giây
	fmt.Println("\n📊 --- TỔNG KẾT DỮ LIỆU TRONG MAP ---")
	models.StatsMutex.Lock()
	for label, data := range models.Stats {
		fmt.Printf("{%s: %s}\n", label, data.Value)
	}
	models.StatsMutex.Unlock()

	fmt.Println("🎉 Chương trình kết thúc thành công.")
}
