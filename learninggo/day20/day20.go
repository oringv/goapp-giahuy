package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func heavyTask(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 0; i < 100e8; i++ {
// 		sum += i
// 	}

// }

// func main() {
// 	numCPU := runtime.NumCPU()
// 	fmt.Println("So luong CPU:", numCPU)

// 	runtime.GOMAXPROCS(numCPU)

// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go heavyTask(&wg)
// 	wg.Wait()
// 	fmt.Println("Tong thoi gian:", time.Since(start))
// }

//**BÀI 87**//

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil {
		return "N/A"
	}
	fmt.Println("%+V", percent)
	value := fmt.Sprintf("%.2f", percent[0])
	return value
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	monitor := CPUMonitor{}
	cpuPercent := monitor.Check(ctx)
	fmt.Println(cpuPercent)
}
