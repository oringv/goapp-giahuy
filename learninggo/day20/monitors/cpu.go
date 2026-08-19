package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

// ✅ BẮT BUỘC PHẢI CÓ ĐOẠN NÀY
type Monitor interface {
	Check(ctx context.Context) string
	Name() string
}

type CPUMonitor struct{}

func (m *CPUMonitor) Name() string { return "CPU" }

func (m *CPUMonitor) Check(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil || len(percent) == 0 {
		return fmt.Sprintf("[CPU Monitor] Cloud not retrieve CPU info: %v \n", err)
	}
	return fmt.Sprintf("%.2f%%", percent[0])
}
