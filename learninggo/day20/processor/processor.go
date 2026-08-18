package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/monitors"
	"sync"
	"time"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Khai báo danh sách các monitor sử dụng Interface
	// Nếu ở monitors bạn đặt là Monitor thì ở đây dùng monitors.Monitor
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
				fmt.Printf("%s : %s\n", m.Name(), m.Check(ctx))
			}
		}
	}
}
