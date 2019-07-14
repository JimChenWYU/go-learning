package main

import (
	"fmt"
	"strconv"
	"testing"
)

// testing.B 包含基准测试使用的函数
// 执行命令 go test -v -run="none" -bench="BenchmarkSprintf"
// -run 选项保证在运行基准测试函数之前没有单元测试会被运行
// -bench 选项表示希望运行哪些基准测试函数
// 这两个选项都可以接收正则表达式，来决定需要运行哪些测试
// b.ResetTimer 是用来重置计时器，防止代码执行前的初始化时间干扰

// go test -v -run="none" -bench=. -benchtime="3s" -benchmem
// -benchmem 展示内存使用情况，单次操作分配的内存，以及从堆上分配内存的次数

// ouput:
//goos: windows
//goarch: amd64
//pkg: go-learning/test-tool
//BenchmarkSprintf-4      30000000               115 ns/op              16 B/op          2 allocs/op
//BenchmarkFormat-4       1000000000               4.33 ns/op            0 B/op          0 allocs/op
//BenchmarkItoa-4         1000000000               4.33 ns/op            0 B/op          0 allocs/op
//PASS
//ok      go-learning/test-tool   13.496s

// 10000000 表示一个执行一千万次，后面数字表示性能，单位为每次操作消耗的纳秒（ns）数

func BenchmarkSprintf(b *testing.B) {
	number := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", number)
	}
}

func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
