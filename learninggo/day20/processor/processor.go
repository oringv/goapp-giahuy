package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

// HÀM BỊ THIẾU: RunMonitor thu thập dữ liệu từ các bộ giám sát
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
			statCh <- models.SystemStats{
				Label: m.Name(),
				Value: val,
			}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	fmt.Println("\n==== 🖥️  System Status Detail ====")

	// Memory tổng quát
	vmStat, _ := mem.VirtualMemoryWithContext(ctx)
	fmt.Printf("[Memory] %.2f%%\n", vmStat.UsedPercent)

	// Disk tổng quát
	diskStat, _ := disk.UsageWithContext(ctx, "/")
	fmt.Printf("[Disk] %.2f%% used\n", diskStat.UsedPercent)

	// Network tổng quát
	netStat, _ := net.IOCountersWithContext(ctx, false)
	if len(netStat) > 0 {
		fmt.Printf("[Network] Send: %v KB, Recv: %v KB\n", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)
	}

	// CPU tổng quát
	cpuP, _ := cpu.PercentWithContext(ctx, time.Second, false)
	if len(cpuP) > 0 {
		fmt.Printf("[CPU] %.2f%%\n", cpuP[0])
	}

	// --- LẤY DANH SÁCH TIẾN TRÌNH (Giới hạn top 5 để tránh lag) ---
	totalMemory := vmStat.Total
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Error] %v \n", err)
	}

	fmt.Println("\n--- Top 5 Processes ---")
	var wg sync.WaitGroup
	limit := 5 // Chỉ lấy 5 thằng đầu tiên cho màn hình sạch sẽ

	for i, p := range processes {
		if i >= limit {
			break
		}
		wg.Add(1)
		go func(proc *process.Process) {
			defer wg.Done()
			name, _ := proc.NameWithContext(ctx)
			cpuPercent, _ := proc.CPUPercentWithContext(ctx)
			memInfo, _ := proc.MemoryInfoWithContext(ctx)
			if memInfo == nil {
				return
			}

			ramPercent := (float64(memInfo.RSS) / float64(totalMemory)) * 100
			createTimeMs, err := proc.CreateTimeWithContext(ctx)
			runningTimeStr := "N/A"
			if err == nil {
				startTime := time.Unix(createTimeMs/1000, (createTimeMs%1000)*1000000)
				runningTimeStr = time.Since(startTime).Truncate(time.Second).String()
			}

			fmt.Printf("PID: %-6d | Name: %-15s | CPU: %-5.2f%% | RAM: %-5.2f%% | Time: %s\n",
				proc.Pid, name, cpuPercent, ramPercent, runningTimeStr)
		}(p)
	}
	wg.Wait()
	return "-----------------------------------"
}
