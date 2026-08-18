package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type NetMonitor struct {
}

func (m *NetMonitor) Name() string {
	return "CPU"
}

func (m *NetMonitor) Check(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil || len(percent) == 0 {
		return "N/A"
	}
	value := fmt.Sprintf("%.2f%%", percent[0])

	return value
}
