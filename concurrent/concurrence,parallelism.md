### 并行

让不同的代码片段同时在不同的物理处理器上执行，也就是同时做很多事情。
go 语言如果要使 goroutine 实现并行，必须使用多于一个逻辑处理器，当有多个逻辑处理器时，调度器会将 goroutine 平均分配到每个逻辑处理器上。

### 并发

并发是可以同时管理很多事情，这些事情可能只做到一半就被暂停去做别的事情了。
go 语言运行时会在逻辑处理器上调度 goroutine 来运行，每个逻辑处理器都分别绑定到单个操作系统线程。
