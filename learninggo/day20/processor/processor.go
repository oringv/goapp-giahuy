package processor

import (
	"context"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"
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
			// Lấy dữ liệu và gửi vào channel dưới dạng Struct
			val := m.Check(ctx)
			statCh <- models.SystemStats{
				Label: m.Name(),
				Value: val,
			}
		}
	}
}
