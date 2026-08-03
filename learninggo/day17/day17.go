package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func task(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Task %d bat dau \n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Task %d ket thuc \n", id)
// }

// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 4; i++ {
// 		wg.Add(1)
// 		go task(i, &wg)
// 	}

// 	fmt.Println("Tong thoi gian hoan thanh: ", time.Since(start))

// }

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30

	}()
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-ch)
	// }

	for value := range ch {
		fmt.Println(value)
	}

	time.Sleep(1 * time.Second)
}
