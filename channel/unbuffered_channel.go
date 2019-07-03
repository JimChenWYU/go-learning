package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// 创建无缓冲的通道，记录球击打数
	court := make(chan int)

	wg.Add(2)

	// 两位选手
	go Player("Nadal", court)
	go Player("Djokovic", court)

	// 发球
	court <- 1

	wg.Wait()
}

func Player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待球被击打
		ball, ok := <-court

		// 没有球击打过来，即通道关闭则胜利
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)

			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}
