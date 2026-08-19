package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

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
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve info: %v \n", err)
	}

	// In ra số lượng tiến trình đang chạy cho gọn
	return fmt.Sprintf("Running Processes: %d", len(processes))
}
