package main

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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStats)
	monitorList := []monitors.Monitor{&monitors.CPUMonitor{}, &monitors.MemoryMonitor{}, &monitors.NetMonitor{}, &monitors.DiskMonitor{}}

	for _, m := range monitorList {
		wg.Add(1)
		go func(val monitors.Monitor) { processor.RunMonitor(ctx, &wg, statCh, val) }(m)
	}

	go func() {
		for stat := range statCh {
			models.StatsMutex.Lock()
			models.Stats[stat.Label] = stat
			models.StatsMutex.Unlock()
		}
	}()

	go func() { wg.Wait(); close(statCh) }()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n🏁 Chương trình kết thúc.")
			return
		case <-ticker.C:
			// Làm sạch màn hình Terminal
			fmt.Print("\033[H\033[2J")

			// 1. Lấy dữ liệu và tự động in bảng Top 5 ra Terminal
			cpuList, memList := processor.GetTopProcesses(ctx)

			// 2. Xuất dữ liệu vào CSV
			processor.ExportToCSV(cpuList, memList)

			// 3. In trạng thái hệ thống tổng quát
			fmt.Println("\n=== System Status ===")
			models.StatsMutex.Lock()
			for _, s := range models.Stats {
				fmt.Printf("[%s] %s \n", s.Label, s.Value)
			}
			models.StatsMutex.Unlock()
		}
	}
}
