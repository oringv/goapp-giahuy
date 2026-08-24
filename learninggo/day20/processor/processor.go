package processor

import (
	"context"
	"fmt"
	"goapp-giahuy/learninggo/day20/models"
	"goapp-giahuy/learninggo/day20/monitors"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/process"
)

// 1. RunMonitor: Bộ máy thu thập dữ liệu chạy ngầm (Goroutine)
func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStats, m monitors.Monitor) {
	defer wg.Done() // Báo cáo đã xong việc khi hàm kết thúc (để WaitGroup trừ đi 1)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C: // Đến nhịp 2 giây
			val, alert := m.Check(ctx) // Đo đạc: lấy giá trị (val) và xem có báo động không (alert)

			stat := models.SystemStats{
				Label:   m.Name(), // Tên bộ đo (CPU, RAM...)
				Value:   val,      // Con số đo được
				IsAlert: alert,    // Trạng thái > 60%
			}

			statCh <- stat // Đẩy dữ liệu vào "đường ống" Channel để gửi về trung tâm

			if alert {
				LogAlert(stat) // Nếu vượt ngưỡng nguy hiểm thì lập tức ghi vào "hộp đen" alert.log
			}
		}
	}
}

// ✅ LogAlert: Ghi lại lịch sử các lần "quá tải" vào file riêng
func LogAlert(stat models.SystemStats) {
	// Mở file: O_APPEND (ghi nối tiếp), O_CREATE (tạo mới nếu chưa có), 0644 (quyền bảo mật chuẩn)
	f, err := os.OpenFile("alert.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	timestamp := time.Now().Format(time.RFC3339) // Lấy giờ chuẩn ISO (có +07:00 giống video)
	line := fmt.Sprintf("[%s] ALERT: %s = %s\n", timestamp, stat.Label, stat.Value)
	f.WriteString(line) // Ghi dòng cảnh báo vào file
}

// 2. GetTopProcesses: Quét toàn bộ máy tính để tìm những ứng dụng ngốn tài nguyên nhất
func GetTopProcesses(ctx context.Context) ([]models.ProStat, []models.ProStat) {
	allProcesses, err := process.ProcessesWithContext(ctx) // Lấy danh sách mọi thứ đang chạy
	if err != nil {
		fmt.Println("❌ Lỗi lấy danh sách tiến trình")
		return nil, nil
	}

	var procList []models.ProStat
	for i, p := range allProcesses {
		if i > 50 {
			break
		}

		name, _ := p.NameWithContext(ctx)
		cpuP, _ := p.CPUPercentWithContext(ctx)
		memP, _ := p.MemoryPercentWithContext(ctx)
		mInfo, err := p.MemoryInfoWithContext(ctx)

		// XỬ LÝ LỖI: Nếu không lấy được thông tin (do quyền Admin) thì dùng 'continue' để bỏ qua, quét thằng tiếp theo
		if err != nil || mInfo == nil {
			continue
		}

		createT, _ := p.CreateTimeWithContext(ctx)
		runTime := "N/A"
		if createT > 0 {
			// Tính toán ứng dụng đã mở được bao lâu
			runTime = time.Since(time.Unix(createT/1000, 0)).Truncate(time.Second).String()
		}

		procList = append(procList, models.ProStat{
			PID: p.Pid, Name: name, CPU: cpuP,
			Memory: mInfo.RSS, RamPercent: float64(memP), RunningTime: runTime,
		})
	}

	// Sắp xếp: Đưa những thằng ngốn CPU nhiều nhất lên đầu bảng
	cpuSorted := make([]models.ProStat, len(procList))
	copy(cpuSorted, procList)
	sort.Slice(cpuSorted, func(i, j int) bool { return cpuSorted[i].CPU > cpuSorted[j].CPU })

	// Sắp xếp: Đưa những thằng ngốn RAM nhiều nhất lên đầu bảng
	memSorted := make([]models.ProStat, len(procList))
	copy(memSorted, procList)
	sort.Slice(memSorted, func(i, j int) bool { return memSorted[i].Memory > memSorted[j].Memory })

	// In kết quả ra màn hình đen (Terminal) để người dùng xem trực tiếp
	fmt.Println()
	PrintTop5(cpuSorted, "Top 5 CPU consuming processes")
	fmt.Println()
	PrintTop5(memSorted, "Top 5 RAM consuming processes")

	return cpuSorted, memSorted // Trả về 2 danh sách để hàm Main mang đi lưu file CSV
}

// 3. PrintTop5: Hàm hỗ trợ vẽ bảng dữ liệu ra Terminal cho đẹp
func PrintTop5(list []models.ProStat, label string) {
	fmt.Printf("== %s ==\n", label)
	for i := 0; i < 5 && i < len(list); i++ { // Chỉ lấy 5 thằng đầu tiên của danh sách đã sắp xếp
		p := list[i]
		// Định dạng: PID - Tên - %CPU - RAM (MB) - Thời gian chạy
		fmt.Printf("[%d] %-15s - CPU: %6.2f%% - RAM: %7.2f MB (%5.2f%%) - Running: %s\n",
			p.PID, p.Name, p.CPU, float64(p.Memory)/1024/1024, p.RamPercent, p.RunningTime)
	}
}

// 4. ExportToCSV: Lưu trữ toàn bộ dữ liệu vào file CSV để xem lại bằng Excel/Edit CSV
func ExportToCSV(cpuList, memList []models.ProStat) {
	f, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// Nếu file mới tinh (Dung lượng = 0) thì mới ghi dòng Tiêu đề (Header)
	if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
		f.WriteString("Timestamp,PID,Name,CPU (%),RAM (MB),RAM (%),Running Time\n")
	}

	ts := time.Now().Format(time.RFC3339) // Giờ ghi log chuẩn ISO

	// Hàm phụ để ghi dữ liệu từng dòng
	save := func(list []models.ProStat) {
		for i := 0; i < 5 && i < len(list); i++ {
			p := list[i]

			f.WriteString(fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s\n",
				ts, p.PID, p.Name, p.CPU, float64(p.Memory)/1024/1024, p.RamPercent, p.RunningTime))
		}
	}

	save(cpuList)
	save(memList)
	fmt.Println("✅ Đã ghi log vào file process_stats.csv")
}
