package main

import (
	"sync"
	"fmt"
)

func main() {
	// 使用
	var a = 1
	var b bool = true
	// 定义加初始化
	c := "string"
	var wg sync.WaitGroup


	wg.Wait()
	
	fmt.Printf("%d %s\n", a, c)
	fmt.Println(b)
}
