package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	/**
	 * 变量 counter 会进行4次读和写操作，每个 goroutine 执行两次。但是，程序终止时，counter变量的值为2。
	 * 之后每个 goroutine 都会覆盖另一个 goroutine 的工作，
	 * 每个 goroutine 都会创建一个 counter 变量的副本，当切换到另一个 goroutine 时，counter 的值已经改变了，但是又没有更新自己的那个副本的值
	 */

	go inCounter(1)
	go inCounter(2)

	wg.Wait()

	fmt.Println("Final Counter: ", counter)
}

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()

		value++

		counter = value
	}
}
