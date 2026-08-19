package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

type NetMonitor struct{}

func (m *NetMonitor) Name() string { return "Net" } // Để tên là Net giống ảnh

func (m *NetMonitor) Check(ctx context.Context) string {
	netStat, err := net.IOCountersWithContext(ctx, false)
	if err != nil || len(netStat) == 0 {
		return "N/A"
	}

	// Trả về định dạng: Send: X KB, Recv: Y KB
	return fmt.Sprintf("Send: %d KB, Recv: %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)
}
