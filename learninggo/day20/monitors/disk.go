package monitors

import (
	"context"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskMonitor struct{}

func (m *DiskMonitor) Name() string {
	return "Disk"
}

func (m *DiskMonitor) Check(ctx context.Context) (string, bool) {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}

	diskStat, err := disk.UsageWithContext(ctx, path)

	if err != nil {
		// ✅ FIX: Thêm ", false" vào đây để đủ 2 giá trị trả về
		return fmt.Sprintf("[Disk Monitor] Could not retrieve Disk info: %v \n", err), false
	}

	value := fmt.Sprintf("%.2f%% used", diskStat.UsedPercent)

	// Dòng này đã đúng (trả về 2 món: value và kết quả so sánh > 60)
	return value, diskStat.UsedPercent > 60
}
