package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup

	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go inCounter(1)
	go inCounter(2)

	wg.Wait()

	fmt.Printf("Final Counter: %d \n", counter)

}

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 临界区，只允许一个 goroutine 进入
		mutex.Lock()
		{
			value := counter

			runtime.Gosched()

			value++

			counter = value
		}
		// 释放锁，允许其他正在等待的 goroutine
		mutex.Unlock()
	}
}
