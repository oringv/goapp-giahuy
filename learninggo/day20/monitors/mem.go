package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryMonitor struct{}

// ✅ FIX 1: Hàm Name chỉ trả về string (không có bool)
func (m *MemoryMonitor) Name() string {
	return "MEM"
}

func (m *MemoryMonitor) Check(ctx context.Context) (string, bool) {
	// ✅ FIX 2: Đổi tên biến từ 'v' thành 'vmStat' cho đồng bộ bên dưới
	vmStat, err := mem.VirtualMemoryWithContext(ctx)

	if err != nil {
		// ✅ FIX 3: Thêm ', false' để đủ 2 giá trị trả về (string, bool)
		return fmt.Sprintf("[Memory Monitor] Could not retrieve Memory info: %v \n", err), false
	}

	// Bây giờ vmStat đã tồn tại nên dòng này sẽ hết lỗi đỏ
	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)

	return value, vmStat.UsedPercent > 60
}
