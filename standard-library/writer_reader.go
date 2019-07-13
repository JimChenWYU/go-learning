package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer的值，并将一个字符串写入Buffer
	// 使用实现 io.Writer 的 Write 的方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	_, _ = fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.Stdout 值的地址作为 io.Writer 类型值传入
	_, _ = b.WriteTo(os.Stdout)
}
