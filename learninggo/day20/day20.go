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
	// Chương trình chạy trong 60 giây
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStats)

	monitorList := []monitors.Monitor{
		&monitors.CPUMonitor{}, &monitors.MemoryMonitor{},
		&monitors.NetMonitor{}, &monitors.DiskMonitor{},
	}

	fmt.Println("🚀 Đang khởi tạo hệ thống giám sát...")

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

	printTicker := time.NewTicker(5 * time.Second)
	defer printTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("🏁 Chương trình kết thúc.")
			return
		case <-printTicker.C:
			// 1. In chi tiết trước
			processor.GetTopProcesses(ctx)

			// 2. In tổng quát sau (Giống thứ tự trong video)
			fmt.Println("\n=== System Status ===")
			models.StatsMutex.Lock()
			for _, s := range models.Stats {
				fmt.Printf("[%s] %s \n", s.Label, s.Value)
			}
			models.StatsMutex.Unlock()
		}
	}
}
