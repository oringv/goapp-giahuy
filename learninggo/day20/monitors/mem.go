package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryMonitor struct{}

func (m *MemoryMonitor) Name() string { return "MEM" }

func (m *MemoryMonitor) Check(ctx context.Context) string {
	v, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Memory Top Processes] Cloud not retrieve Memory info: %v \n", err)
	}
	// ✅ FIX: Đổi Print thành Printf
	// fmt.Printf("DEBUG: %+v\n", v)
	return fmt.Sprintf("%.2f%%", v.UsedPercent)
}
