package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
)

// 1. Hàm chạy thu thập dữ liệu từ các Monitor
func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStats, m monitors.Monitor) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			val := m.Check(ctx)
			statCh <- models.SystemStats{Label: m.Name(), Value: val}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	fmt.Println("\n--- Top 3 Processes Detail ---")

	// 1. Lấy % CPU tổng quát (Đo trong 500ms để không bị chậm)
	cpuPercentAll, _ := cpu.PercentWithContext(ctx, 500*time.Millisecond, false)
	if len(cpuPercentAll) > 0 {
		fmt.Printf("🔥 Global CPU Usage: %.2f%%\n", cpuPercentAll[0])
	}

	// 2. Lấy danh sách tiến trình
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return "Error"
	}

	limit := 3
	count := 0
	for _, p := range processes {
		if count >= limit {
			break
		}

		name, _ := p.NameWithContext(ctx)

		cpuP, _ := p.CPUPercentWithContext(ctx)
		mPercent, _ := p.MemoryPercentWithContext(ctx)
		mInfo, err := p.MemoryInfoWithContext(ctx)

		if err != nil || mInfo == nil {
			continue
		}

		createT, _ := p.CreateTimeWithContext(ctx)
		runTime := "N/A"
		if createT > 0 {
			runTime = time.Since(time.Unix(createT/1000, 0)).Truncate(time.Second).String()
		}

		s := models.ProStat{
			PID:         p.Pid,
			Name:        name,
			CPU:         cpuP,
			Memory:      mInfo.RSS,
			RamPercent:  float64(mPercent),
			RunningTime: runTime,
		}
		fmt.Printf("%+v\n", s)
		count++
	}

	return "-----------------------------------"
}
