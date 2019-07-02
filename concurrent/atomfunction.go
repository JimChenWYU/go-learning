package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter  int64
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(4)

	go inCounter(1)
	go inCounter(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
	fmt.Println("Finish Counter: ", counter)
}

/**
 * 使用 atomic 包的 AddInt64 函数，该函数会同步整型值的加法，
 * 方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
 *
 */

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)

		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()
	}
}

/**
 * LoadInt64 和 StoreInt64，这两个函数提供了一种安全地读和写一个整型值的方式。
 */
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
