package main

import (
	"fmt"
)

/**
 * 可定义多个 init 函数，这些函数会在程序执行开始的时候被调用，调用顺序为定义的顺序
 * init 函数主要用于设置包，初始化变量或者程序其他的引导工作
 */

func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init3")
}
