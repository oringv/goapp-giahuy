package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"sort" // ✅ Cần thư viện này để sắp xếp Top 5
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

// 1. Giữ nguyên hàm RunMonitor cũ của bạn
func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStats, m monitors.Monitor) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			statCh <- models.SystemStats{Label: m.Name(), Value: m.Check(ctx)}
		}
	}
}

// 2. ✅ HÀM NÂNG CẤP: Lấy chi tiết và phân loại Top 5 (Giống 100% ảnh mẫu)
func GetTopProcesses(ctx context.Context) {
	allProcesses, err := process.ProcessesWithContext(ctx)
	if err != nil {
		fmt.Println("❌ Lỗi lấy danh sách tiến trình")
		return
	}

	var procList []models.ProStat

	// Thu thập dữ liệu thô từ tất cả tiến trình (Giới hạn 50 thằng đầu tiên để tránh lag máy)
	for i, p := range allProcesses {
		if i > 50 {
			break
		}

		name, _ := p.NameWithContext(ctx)
		cpuP, _ := p.CPUPercentWithContext(ctx)
		memP, _ := p.MemoryPercentWithContext(ctx)
		mInfo, err := p.MemoryInfoWithContext(ctx)
		if err != nil || mInfo == nil {
			continue
		}

		createT, _ := p.CreateTimeWithContext(ctx)
		runTime := "N/A"
		if createT > 0 {
			runTime = time.Since(time.Unix(createT/1000, 0)).Truncate(time.Second).String()
		}

		procList = append(procList, models.ProStat{
			PID: p.Pid, Name: name, CPU: cpuP,
			Memory: mInfo.RSS, RamPercent: float64(memP), RunningTime: runTime,
		})
	}

	// --- PHẦN 1: TOP 5 CPU ---
	fmt.Println("\n== Top 5 CPU consuming processes ==")
	sort.Slice(procList, func(i, j int) bool {
		return procList[i].CPU > procList[j].CPU
	})
	printTop5(procList)

	// --- PHẦN 2: TOP 5 RAM ---
	fmt.Println("\n== Top 5 RAM consuming processes ==")
	sort.Slice(procList, func(i, j int) bool {
		return procList[i].Memory > procList[j].Memory
	})
	printTop5(procList)
}

// Hàm hỗ trợ in định dạng đẹp như ảnh mẫu
func printTop5(list []models.ProStat) {
	for i := 0; i < 5 && i < len(list); i++ {
		p := list[i]
		// Định dạng: [PID] Name - CPU: % - RAM: MB (%) - Running: Time
		fmt.Printf("[%d] %-15s - CPU: %6.2f%% - RAM: %7.2f MB (%5.2f%%) - Running: %s\n",
			p.PID, p.Name, p.CPU, float64(p.Memory)/1024/1024, p.RamPercent, p.RunningTime)
	}
}
