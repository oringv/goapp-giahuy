package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

// 1. RunMonitor: Thu thập chỉ số hệ thống
func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStats, m monitors.Monitor) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			val, alert := m.Check(ctx)
			stat := models.SystemStats{
				Label:   m.Name(),
				Value:   val,
				IsAlert: alert,
			}
			statCh <- stat

			if alert {
				LogAlert(stat)
			}
		}
	}
}

// ✅ Hàm LogAlert duy nhất - Ghi log vào file chuẩn như video
func LogAlert(stat models.SystemStats) {
	f, err := os.OpenFile("alert.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339)
	line := fmt.Sprintf("[%s] ALERT: %s = %s\n", timestamp, stat.Label, stat.Value)
	f.WriteString(line)
}

// 2. GetTopProcesses: Lấy danh sách tiến trình
func GetTopProcesses(ctx context.Context) ([]models.ProStat, []models.ProStat) {
	allProcesses, err := process.ProcessesWithContext(ctx)
	if err != nil {
		fmt.Println("❌ Lỗi lấy danh sách tiến trình")
		return nil, nil
	}

	var procList []models.ProStat
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

	cpuSorted := make([]models.ProStat, len(procList))
	copy(cpuSorted, procList)
	sort.Slice(cpuSorted, func(i, j int) bool { return cpuSorted[i].CPU > cpuSorted[j].CPU })

	memSorted := make([]models.ProStat, len(procList))
	copy(memSorted, procList)
	sort.Slice(memSorted, func(i, j int) bool { return memSorted[i].Memory > memSorted[j].Memory })

	fmt.Println()
	PrintTop5(cpuSorted, "Top 5 CPU consuming processes")
	fmt.Println()
	PrintTop5(memSorted, "Top 5 RAM consuming processes")

	return cpuSorted, memSorted
}

// 3. PrintTop5: Helper in bảng
func PrintTop5(list []models.ProStat, label string) {
	fmt.Printf("== %s ==\n", label)
	for i := 0; i < 5 && i < len(list); i++ {
		p := list[i]
		fmt.Printf("[%d] %-15s - CPU: %6.2f%% - RAM: %7.2f MB (%5.2f%%) - Running: %s\n",
			p.PID, p.Name, p.CPU, float64(p.Memory)/1024/1024, p.RamPercent, p.RunningTime)
	}
}

// 4. ExportToCSV: Xuất file CSV chuẩn
func ExportToCSV(cpuList, memList []models.ProStat) {
	f, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
		f.WriteString("Timestamp,PID,Name,CPU (%),RAM (MB),RAM (%),Running Time\n")
	}

	ts := time.Now().Format(time.RFC3339)
	save := func(list []models.ProStat) {
		for i := 0; i < 5 && i < len(list); i++ {
			p := list[i]
			f.WriteString(fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
				ts, p.PID, p.Name, p.CPU, float64(p.Memory)/1024/1024, p.RamPercent, p.RunningTime))
		}
	}
	save(cpuList)
	save(memList)
	fmt.Println("✅ Đã ghi log vào file process_stats.csv")
}
