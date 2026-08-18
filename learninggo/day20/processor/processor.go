package processor

import (
	"context"
	"goapp-giahuy/learninggo/day20/models" // Import model mới
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStats, m models.Montitor) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	ms := []monitors.Monitor{
		&monitors.CPUMonitor{},
		&monitors.MemoryMonitor{},
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, m := range ms {
				val := m.Check(ctx)
				// Gửi dữ liệu vào channel dưới dạng Struct
				statCh <- models.SystemStats{Label: m.Name(), Value: val}
			}
		}
	}
}
