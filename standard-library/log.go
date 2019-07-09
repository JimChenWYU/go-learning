package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

/**
 * Ldate = 1 << iota  - 日期：2019/07/08
 * Ltime - 时间：20:40:04
 * Lmicroseconds - 毫秒级时间：20:40:04.049050
 * Llongfile - 完整路径的文件名和行号：/home/jim/go-learning/standard-library/log.go:18
 * Lshortfile - 最终文件名元素和行号：log.go:18
 * LstdFlags - 标准日志记录器的初始值
 */

/**
 * iota 初始值为0，在常量声明区里有特殊的作用，这个关键字让编译器为每个常量复制相同的表达式，直到声明区结束，或者遇到一个新的赋值语句
 * const (
		Ldate = 1 << iota // 1 << 0 == 000000001 = 1
		Ltime             // 1 << 1 == 000000010 = 2
		Lmicroseconds     // 1 << 2 == 000000100 = 4
		Llongfile         // 1 << 3 == 000001000 = 8
		Lshortfile        // 1 << 4 == 000010000 = 16
		LstdFlags = Ldate | Ltime  // 打破 iota 常数链 = 000000001 | 000000010 = 000000011 = 3
	)
*/
func init() {
	log.SetPrefix("TRACE: ")
	// 按位来设置需要的日志参数
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var (
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要日志
	Warning *log.Logger // 需要注意的日志
	Error   *log.Logger // 非常严重的日志
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 日志记录器都是多 goroutine 安全的，意味在多个 goroutine 可以同时调用来自同一个日志记录器的这些函数，不会有彼此的写冲突
func main() {

	Trace.Println("I have something standard to say.")
	Info.Println("Special Information.")
	Warning.Println("There is something you need to know about.")
	Error.Println("Something has failed.")

	// 写到标准日志记录器
	log.Println("message")

	// 在调用 Println() 之后会调用 os.Exit(1)
	log.Fatalln("Fatal message")

	// 在调用 Println() 之后会调用 panic()
	log.Panicln("Panic message")
}
