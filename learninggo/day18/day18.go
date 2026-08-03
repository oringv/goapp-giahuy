package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d bat dau \n", id)
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d hoan thanh", id)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan string)

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, ch, &wg)
	}

	wg.Wait()

	fmt.Println("Tong thoi gian hoan thanh: ", time.Since(start))
}
