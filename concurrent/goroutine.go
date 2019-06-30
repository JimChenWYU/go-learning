package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器用
	// 可通过设置该值让应用程序并行运行
	//
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	// 计数加2，表示要等待两个 goroutine
	// WaitGroup 是一个计数信号量，可以用来记录并维护运行的 goroutine，如果 WaitGroup 的值大于0，Wait 方法就会阻塞
	// 为了减少 WaitGroup 的值并最终释放 main 函数，使用 defer 声明在函数退出时调用 Done 方法
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
